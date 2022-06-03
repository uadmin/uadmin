package uadmin

import (
	"fmt"
	"reflect"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func buildMysqlDriver() *dbDriver {
	return &dbDriver{
		createM2MTable:             "CREATE TABLE `{TABLE1}_{TABLE2}` (`table1_id` int(10) unsigned NOT NULL, `table2_id` int(10) unsigned NOT NULL, PRIMARY KEY (`table1_id`,`table2_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;",
		selectM2M:                  "SELECT `table2_id` FROM `{TABLE1}_{TABLE2}` WHERE table1_id={TABLE1_ID};",
		deleteM2M:                  "DELETE FROM `{TABLE1}_{TABLE2}` WHERE `table1_id`={TABLE1_ID};",
		insertM2M:                  "INSERT INTO `{TABLE1}_{TABLE2}` VALUES ({TABLE1_ID}, {TABLE2_ID});",
		open:                       mysqlOpen,
		createDB:                   mysqlCreateDB,
		lastInsertID:               mysqlLastInsertID,
		delete:                     mysqlDelete,
		getQueryOperatorContains:   mysqlGetQueryOperatorContains,
		getQueryOperatorStartsWith: mysqlGetQueryOperatorStartsWith,
		getQueryOperatorEndsWith:   mysqlGetQueryOperatorEndsWith,
		apiRead:                    mysqlApiRead,
	}
}

func mysqlOpen() (db *gorm.DB, err error) {
	if Database.Host == "" || Database.Host == "localhost" {
		Database.Host = "127.0.0.1"
	}
	if Database.Port == 0 {
		Database.Port = 3306
	}

	if Database.User == "" {
		Database.User = "root"
	}

	credential := Database.User

	if Database.Password != "" {
		credential = fmt.Sprintf("%s:%s", Database.User, Database.Password)
	}
	dsn := fmt.Sprintf("%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		credential,
		Database.Host,
		Database.Port,
		Database.Name,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: func() logger.Interface {
			if DebugDB {
				return logger.Default.LogMode(logger.Info)
			}
			return logger.Default.LogMode(logger.Silent)
		}(),
	})

	// Check if the error is DB doesn't exist and create it
	if err != nil && err.Error() == "Error 1049: Unknown database '"+Database.Name+"'" {
		err = mysqlCreateDB()

		if err == nil {
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
				Logger: func() logger.Interface {
					if DebugDB {
						return logger.Default.LogMode(logger.Info)
					}
					return logger.Default.LogMode(logger.Silent)
				}(),
			})
		}
	}

	// Temp solution for 0 foreign key
	db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	return
}

func mysqlCreateDB() error {
	credential := Database.User

	if Database.Password != "" {
		credential = fmt.Sprintf("%s:%s", Database.User, Database.Password)
	}

	dsn := fmt.Sprintf("%s@(%s:%d)/?charset=utf8&parseTime=True&loc=Local",
		credential,
		Database.Host,
		Database.Port,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: func() logger.Interface {
			if DebugDB {
				return logger.Default.LogMode(logger.Info)
			}
			return logger.Default.LogMode(logger.Silent)
		}(),
	})
	if err != nil {
		return err
	}

	Trail(INFO, "Database doens't exist, creating a new database")
	db = db.Exec("CREATE SCHEMA `" + Database.Name + "` DEFAULT CHARACTER SET utf8 COLLATE utf8_bin")

	if db.Error != nil {
		return fmt.Errorf(db.Error.Error())
	}

	return nil
}

func mysqlLastInsertID(db *gorm.DB, tableName string) (*gorm.DB, []int) {
	id := []int{}
	db = db.Raw("SELECT LAST_INSERT_ID() AS lastid")
	db.Table(tableName).Pluck("lastid", &id)
	return db, id
}

func mysqlDelete(model reflect.Value, q string, args []interface{}) (int64, error) {
	db = db.Where(q, args...).Delete(model.Addr().Interface())
	return db.RowsAffected, db.Error
}

func mysqlGetQueryOperatorContains(v, nTerm string) string {
	return strings.TrimSuffix(v, "__contains") + "`" + nTerm + " LIKE BINARY ?"
}

func mysqlGetQueryOperatorStartsWith(v, nTerm string) string {
	return strings.TrimSuffix(v, "__startswith") + "`" + nTerm + " LIKE BINARY ?"
}

func mysqlGetQueryOperatorEndsWith(v, nTerm string) string {
	return strings.TrimSuffix(v, "__endswith") + "`" + nTerm + " LIKE BINARY ?"
}

func mysqlApiRead(SQL string, args []interface{}, m interface{}, customSchema bool) (int64, interface{}, error) {
	if !customSchema {
		if res := db.Raw(SQL, args...).Scan(m); res.Error != nil {
			return 0, m, res.Error
		}
	} else {
		rows, err := db.Raw(SQL, args...).Rows()
		if err != nil {
			return 0, m, err
		}
		m = parseCustomDBSchema(rows)
	}
	var rowsCount int64
	if a, ok := m.([]map[string]interface{}); ok {
		rowsCount = int64(len(a))
	} else {
		rowsCount = int64(reflect.ValueOf(m).Elem().Len())
	}
	return rowsCount, m, nil
}

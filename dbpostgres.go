package uadmin

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func buildPostgresDriver() *dbDriver {
	return &dbDriver{
		createM2MTable:             "CREATE TABLE `{TABLE1}_{TABLE2}` (`table1_id` int(10) unsigned NOT NULL, `table2_id` int(10) unsigned NOT NULL, PRIMARY KEY (`table1_id`,`table2_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;",
		selectM2M:                  "SELECT `table2_id` FROM `{TABLE1}_{TABLE2}` WHERE table1_id={TABLE1_ID};",
		deleteM2M:                  "DELETE FROM `{TABLE1}_{TABLE2}` WHERE `table1_id`={TABLE1_ID};",
		insertM2M:                  "INSERT INTO `{TABLE1}_{TABLE2}` VALUES ({TABLE1_ID}, {TABLE2_ID});",
		open:                       postgresOpen,
		createDB:                   postgresCreateDB,
		lastInsertID:               postgresLastInsertID,
		delete:                     postgresDelete,
		getQueryOperatorContains:   postgresGetQueryOperatorContains,
		getQueryOperatorStartsWith: postgresGetQueryOperatorStartsWith,
		getQueryOperatorEndsWith:   postgresGetQueryOperatorEndsWith,
		apiRead:                    postgresApiRead,
	}
}

func postgresOpen() (*gorm.DB, error) {
	dsnMap := map[string]string{
		"sslmode": "disable",
	}
	if Database.Host == "" || Database.Host == "localhost" {
		Database.Host = "127.0.0.1"
	}
	dsnMap["host"] = Database.Host
	if Database.Port != 0 {
		dsnMap["port"] = strconv.Itoa(Database.Port)
	}

	if Database.User == "" {
		Database.User = "postgres"
	}
	dsnMap["user"] = Database.User
	if Database.Password != "" {
		dsnMap["password"] = Database.Password
	}
	dsnArr := make([]string, len(dsnMap), 0)
	for k, v := range dsnMap {
		dsnArr = append(dsnArr, fmt.Sprintf("%s=%v", k, v))
	}
	dsn := strings.Join(dsnArr, " ")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: func() logger.Interface {
			if DebugDB {
				return logger.Default.LogMode(logger.Info)
			}
			return logger.Default.LogMode(logger.Silent)
		}(),
	})
	return db, err
}

func postgresCreateDB() error {
	panic("postgresql not implemented")
}

func postgresLastInsertID(db *gorm.DB, tableName string) (*gorm.DB, []int) {
	panic("postgresql not implemented")
}

func postgresDelete(model reflect.Value, q string, args []interface{}) (int64, error) {
	panic("postgresql not implemented")
}

func postgresGetQueryOperatorContains(v, nTerm string) string {
	panic("postgresql not implemented")
}

func postgresGetQueryOperatorStartsWith(v, nTerm string) string {
	panic("postgresql not implemented")
}

func postgresGetQueryOperatorEndsWith(v, nTerm string) string {
	panic("postgresql not implemented")
}

func postgresApiRead(SQL string, args []interface{}, m interface{}, customSchema bool) (int64, interface{}, error) {
	panic("postgresql not implemented")
}

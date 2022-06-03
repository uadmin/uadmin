package uadmin

import (
	"reflect"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func buildSqliteDriver() *dbDriver {
	return &dbDriver{
		//createM2MTable: "CREATE TABLE `{TABLE1}_{TABLE2}` (`{TABLE1}_id`	INTEGER NOT NULL,`{TABLE2}_id` INTEGER NOT NULL, PRIMARY KEY(`{TABLE1}_id`,`{TABLE2}_id`));",
		createM2MTable: "CREATE TABLE `{TABLE1}_{TABLE2}` (`table1_id`	INTEGER NOT NULL,`table2_id` INTEGER NOT NULL, PRIMARY KEY(`table1_id`,`table2_id`));",
		selectM2M:                  "SELECT `table2_id` FROM `{TABLE1}_{TABLE2}` WHERE table1_id={TABLE1_ID};",
		deleteM2M:                  "DELETE FROM `{TABLE1}_{TABLE2}` WHERE `table1_id`={TABLE1_ID};",
		insertM2M:                  "INSERT INTO `{TABLE1}_{TABLE2}` VALUES ({TABLE1_ID}, {TABLE2_ID});",
		open:                       sqliteOpen,
		createDB:                   sqliteCreateDB,
		lastInsertID:               sqliteLastInsertID,
		delete:                     sqliteDelete,
		getQueryOperatorContains:   sqliteGetQueryOperatorContains,
		getQueryOperatorStartsWith: sqliteGetQueryOperatorStartsWith,
		getQueryOperatorEndsWith:   sqliteGetQueryOperatorEndsWith,
		apiRead:                    sqliteApiRead,
	}
}

func sqliteOpen() (*gorm.DB, error) {
	dbName := Database.Name
	if dbName == "" {
		dbName = "uadmin.db"
	}
	return gorm.Open(sqlite.Open(dbName), &gorm.Config{
		Logger: func() logger.Interface {
			if DebugDB {
				return logger.Default.LogMode(logger.Info)
			}
			return logger.Default.LogMode(logger.Silent)
		}(),
	})
}

func sqliteCreateDB() error {
	return nil
}

func sqliteLastInsertID(db *gorm.DB, tableName string) (*gorm.DB, []int) {
	id := []int{}
	db = db.Raw("SELECT last_insert_rowid() AS lastid")
	db.Table(tableName).Pluck("lastid", &id)
	return db, id
}

func sqliteDelete(model reflect.Value, q string, args []interface{}) (int64, error) {
	tx := db.Begin()

	if res := tx.Exec("PRAGMA case_sensitive_like=ON;"); res.Error != nil {
		tx.Rollback()
		return res.RowsAffected, res.Error
	}

	var rowsAffected int64
	if res := db.Where(q, args...).Delete(model.Addr().Interface()); res.Error != nil {
		tx.Rollback()
		return res.RowsAffected, res.Error
	} else {
		rowsAffected = res.RowsAffected
	}
	if res := db.Exec("PRAGMA case_sensitive_like=OFF;"); res.Error != nil {
		tx.Rollback()
		return res.RowsAffected, res.Error
	}
	return rowsAffected, tx.Commit().Error
}

func sqliteGetQueryOperatorContains(v, nTerm string) string {
	return strings.TrimSuffix(v, "__contains") + "`" + nTerm + " LIKE ?"
}

func sqliteGetQueryOperatorStartsWith(v, nTerm string) string {
	return strings.TrimSuffix(v, "__startswith") + "`" + nTerm + " LIKE ?"
}

func sqliteGetQueryOperatorEndsWith(v, nTerm string) string {
	return strings.TrimSuffix(v, "__endswith") + "`" + nTerm + " LIKE ?"
}

func sqliteApiRead(SQL string, args []interface{}, m interface{}, customSchema bool) (int64, interface{}, error) {
	tx := db.Begin()
	if res := tx.Exec("PRAGMA case_sensitive_like=ON;"); res.Error != nil {
		return 0, m, res.Error
	}
	if !customSchema {
		if res := tx.Raw(SQL, args...).Scan(m); res.Error != nil {
			tx.Rollback()
			return 0, m, res.Error
		}
	} else {
		rows, err := tx.Raw(SQL, args...).Rows()
		if err != nil {
			return 0, m, err
		}
		m = parseCustomDBSchema(rows)
	}
	if res := tx.Exec("PRAGMA case_sensitive_like=OFF;"); res.Error != nil {
		return 0, m, res.Error
	}
	if res := tx.Commit(); res.Error != nil {
		return 0, m, res.Error
	}
	var rowsCount int64
	if a, ok := m.([]map[string]interface{}); ok {
		rowsCount = int64(len(a))
	} else {
		rowsCount = int64(reflect.ValueOf(m).Elem().Len())
	}
	return rowsCount, m, nil
}

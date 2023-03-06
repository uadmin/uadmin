package uadmin

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"gorm.io/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mssql"

	// Enable MYSQL
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"

	// Enable PostgreSQL
	"gorm.io/driver/postgres"

	"encoding/json"

	"io/ioutil"
	"net/http"
	"time"

	// Enable SQLLite
	"github.com/uadmin/uadmin/colors"
	"gorm.io/driver/sqlite"

	"github.com/thlib/go-timezone-local/tzlocal"
)

var db *gorm.DB

var sqlDialect = map[string]map[string]string{
	"mysql": {
		"createM2MTable": "CREATE TABLE `{TABLE1}_{TABLE2}` (`table1_id` int(10) unsigned NOT NULL, `table2_id` int(10) unsigned NOT NULL, PRIMARY KEY (`table1_id`,`table2_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;",
		"selectM2M":      "SELECT `table2_id` FROM `{TABLE1}_{TABLE2}` WHERE table1_id=?;",
		"deleteM2M":      "DELETE FROM `{TABLE1}_{TABLE2}` WHERE `table1_id`=?;",
		"insertM2M":      "INSERT INTO `{TABLE1}_{TABLE2}` VALUES (?, ?);",
		"selectM2MT2":    "SELECT DISTINCT `table1_id` FROM `{TABLE1}_{TABLE2}` WHERE table2_id IN (?);",
	},
	"postgres": {
		"createM2MTable": `CREATE TABLE "{TABLE1}_{TABLE2}" ("table1_id" BIGINT NOT NULL, "table2_id" BIGINT NOT NULL, PRIMARY KEY ("table1_id","table2_id"))`,
		"selectM2M":      `SELECT "table2_id" FROM "{TABLE1}_{TABLE2}" WHERE table1_id=?;`,
		"deleteM2M":      `DELETE FROM "{TABLE1}_{TABLE2}" WHERE "table1_id"=?;`,
		"insertM2M":      `INSERT INTO "{TABLE1}_{TABLE2}" VALUES (?, ?);`,
		"selectM2MT2":    "SELECT DISTINCT `table1_id` FROM `{TABLE1}_{TABLE2}` WHERE table2_id IN (?);",
	},
	"sqlite": {
		//"createM2MTable": "CREATE TABLE `{TABLE1}_{TABLE2}` (`{TABLE1}_id`	INTEGER NOT NULL,`{TABLE2}_id` INTEGER NOT NULL, PRIMARY KEY(`{TABLE1}_id`,`{TABLE2}_id`));",
		"createM2MTable": "CREATE TABLE `{TABLE1}_{TABLE2}` (`table1_id`	INTEGER NOT NULL,`table2_id` INTEGER NOT NULL, PRIMARY KEY(`table1_id`,`table2_id`));",
		"selectM2M":      "SELECT `table2_id` FROM `{TABLE1}_{TABLE2}` WHERE table1_id=?;",
		"deleteM2M":      "DELETE FROM `{TABLE1}_{TABLE2}` WHERE `table1_id`=?;",
		"insertM2M":      "INSERT INTO `{TABLE1}_{TABLE2}` VALUES (?, ?);",
		"selectM2MT2":    "SELECT DISTINCT `table1_id` FROM `{TABLE1}_{TABLE2}` WHERE table2_id IN (?);",
	},
}

// Database is the active Database settings
var Database *DBSettings
var dbOK = false

// DBSettings !
type DBSettings struct {
	Type     string `json:"type"` // sqlite, mysql, postgres
	Name     string `json:"name"` // File/DB name
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Timezone string `json:"timezone"`
}

// initializeDB opens the connection the DB
func initializeDB(a ...interface{}) {
	// Open the connection the the DB
	db = GetDB()

	// Migrate schema
	for i, model := range a {
		Trail(INFO, "Initializing DB: [%s%d/%d%s]", colors.FGGreenB, i+1, len(a), colors.FGNormal)
		err := db.AutoMigrate(model)
		if err != nil {
			Trail(ERROR, "Unable to migrate schema of %s. %s", reflect.TypeOf(model).Name(), err)
		}
		err = customMigration(model)
		if err != nil {
			Trail(ERROR, "Unable to custom migrate schema of %s. %s", reflect.TypeOf(model).Name(), err)
		}
	}
	Trail(OK, "Initializing DB: [%s%d/%d%s]", colors.FGGreenB, len(a), len(a), colors.FGNormal)
	db.AllowGlobalUpdate = true
}

func customMigration(a interface{}) (err error) {
	t := reflect.TypeOf(a)
	for i := 0; i < t.NumField(); i++ {
		// Check if there is any m2m fields
		if t.Field(i).Type.Kind() == reflect.Slice && t.Field(i).Type.Elem().Kind() == reflect.Struct {
			table1 := strings.ToLower(t.Name())
			table2 := strings.ToLower(t.Field(i).Type.Elem().Name())

			//Check if the table is created for the m2m field
			if !db.Migrator().HasTable(table1 + "_" + table2) {
				sql := sqlDialect[Database.Type]["createM2MTable"]
				sql = strings.Replace(sql, "{TABLE1}", table1, -1)
				sql = strings.Replace(sql, "{TABLE2}", table2, -1)
				err = db.Exec(sql).Error
				if err != nil {
					Trail(ERROR, "Unable to create M2M table. %s", err)
					Trail(ERROR, sql)
					return err
				}
			}
		}
	}
	return err
}

// GetDB returns a pointer to the DB
func GetDB() *gorm.DB {
	if db != nil {
		return db
	}
	var err error

	// Check if there is a database config file
	if Database == nil {
		buf, err := ioutil.ReadFile(".database")
		if err == nil {
			err = json.Unmarshal(buf, &Database)
			if err != nil {
				Trail(WARNING, ".database file is not a valid json file. %s", err)
			}
		}

		// Get environment variables for db settings
		if v := os.Getenv("UADMIN_DB_TYPE"); v != "" {
			if Database == nil {
				Database = &DBSettings{}
			}
			Database.Type = v
		}
		if v := os.Getenv("UADMIN_DB_HOST"); v != "" {
			Database.Host = v
			if Database == nil {
				Database = &DBSettings{}
			}
		}
		if v := os.Getenv("UADMIN_DB_PORT"); v != "" {
			if Database == nil {
				Database = &DBSettings{}
			}
			port, err := strconv.ParseInt(v, 10, 64)
			if err == nil {
				Database.Port = int(port)
			} else {
				Trail(WARNING, "Environment vaiable UADMIN_DB_PORT should be a number but got (%s)", v)
			}
		}
		if v := os.Getenv("UADMIN_DB_NAME"); v != "" {
			if Database == nil {
				Database = &DBSettings{}
			}
			Database.Name = v
		}
		if v := os.Getenv("UADMIN_DB_USER"); v != "" {
			if Database == nil {
				Database = &DBSettings{}
			}
			Database.User = v
		}
		if v := os.Getenv("UADMIN_DB_PASSWORD"); v != "" {
			if Database == nil {
				Database = &DBSettings{}
			}
			Database.Password = v
		}
	}

	if Database == nil {
		Database = &DBSettings{
			Type: "sqlite",
		}
	}

	if strings.ToLower(Database.Type) == "sqlite" {
		dbName := Database.Name
		if dbName == "" {
			dbName = "uadmin.db"
		}
		db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{
			Logger: func() logger.Interface {
				if DebugDB {
					return logger.Default.LogMode(logger.Info)
				}
				return logger.Default.LogMode(logger.Silent)
			}(),
		})
	} else if strings.ToLower(Database.Type) == "mysql" {
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
		tz := "Local"
		if Database.Timezone != "" {
			tz = Database.Timezone
		}
		dsn := fmt.Sprintf("%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
			credential,
			Database.Host,
			Database.Port,
			Database.Name,
			tz,
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: func() logger.Interface {
				if DebugDB {
					return logger.Default.LogMode(logger.Info)
				}
				return logger.Default.LogMode(logger.Silent)
			}(),
			DisableForeignKeyConstraintWhenMigrating: true,
		})

		// Check if the error is DB doesn't exist and create it
		if err != nil && strings.Contains(err.Error(), "Unknown database '"+Database.Name+"'") {
			err = createDB()

			if err == nil {
				db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
					Logger: func() logger.Interface {
						if DebugDB {
							return logger.Default.LogMode(logger.Info)
						}
						return logger.Default.LogMode(logger.Silent)
					}(),
					DisableForeignKeyConstraintWhenMigrating: true,
				})
			}
		}

		// fail if we could not connect to DB
		if err != nil {
			Trail(ERROR, "Unable to connect to db. %s", err)
			os.Exit(2)
		}

		// Temp solution for 0 foreign key
		err = db.Exec("SET PERSIST FOREIGN_KEY_CHECKS=0;").Error
		if err != nil {
			err = db.Exec("SET GLOBAL FOREIGN_KEY_CHECKS=0;").Error
			if err != nil {
				err = db.Exec("SET FOREIGN_KEY_CHECKS=0;").Error
				if err != nil {
					Trail(ERROR, "Unable to run FOREIGN_KEY_CHECKS=0. %s", err)
				}
			}
		}

	} else if strings.ToLower(Database.Type) == "postgres" {
		if Database.Host == "" || Database.Host == "localhost" {
			Database.Host = "127.0.0.1"
		}
		if Database.Port == 0 {
			Database.Port = 5432
		}

		if Database.User == "" {
			Database.User = "postgres"
		}
		tz, _ := tzlocal.RuntimeTZ()
		if Database.Timezone != "" {
			tz = Database.Timezone
		}
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
			Database.Host,
			Database.User,
			Database.Password,
			Database.Name,
			Database.Port,
			tz,
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: func() logger.Interface {
				if DebugDB {
					return logger.Default.LogMode(logger.Info)
				}
				return logger.Default.LogMode(logger.Silent)
			}(),
			DisableForeignKeyConstraintWhenMigrating: true,
		})

		if err != nil && strings.HasSuffix(err.Error(), "server error (FATAL: database \""+Database.Name+"\" does not exist (SQLSTATE 3D000))") {
			err = createDB()

			if err == nil {
				db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
					Logger: func() logger.Interface {
						if DebugDB {
							return logger.Default.LogMode(logger.Info)
						}
						return logger.Default.LogMode(logger.Silent)
					}(),
					DisableForeignKeyConstraintWhenMigrating: true,
				})
			}
		}

		// fail if we could not connect to DB
		if err != nil {
			Trail(ERROR, "Unable to connect to db. %s", err)
			os.Exit(2)
		}
	}

	if err != nil {
		Trail(ERROR, "unable to connect to DB. %s", err)
		db.Error = fmt.Errorf("unable to connect to DB. %s", err)
	}
	return db
}

func createDB() error {
	if Database.Type == "mysql" {
		credential := Database.User

		if Database.Password != "" {
			credential = fmt.Sprintf("%s:%s", Database.User, Database.Password)
		}

		dsn := fmt.Sprintf("%s@(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local",
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
		db = db.Exec("CREATE SCHEMA `" + Database.Name + "` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci")

		if db.Error != nil {
			return fmt.Errorf(db.Error.Error())
		}

		return nil
	} else if Database.Type == "postgres" {
		if Database.Host == "" || Database.Host == "localhost" {
			Database.Host = "127.0.0.1"
		}
		if Database.Port == 0 {
			Database.Port = 5432
		}

		if Database.User == "" {
			Database.User = "postgres"
		}
		tz, _ := tzlocal.RuntimeTZ()
		if Database.Timezone != "" {
			tz = Database.Timezone
		}
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%d sslmode=disable TimeZone=%s",
			Database.Host,
			Database.User,
			Database.Password,
			Database.Port,
			tz,
		)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: func() logger.Interface {
				if DebugDB {
					return logger.Default.LogMode(logger.Info)
				}
				return logger.Default.LogMode(logger.Silent)
			}(),
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			return err
		}

		Trail(INFO, "Database doens't exist, creating a new database")
		db = db.Exec("CREATE DATABASE " + Database.Name + " WITH OWNER = " + Database.User + " ENCODING = 'UTF8' CONNECTION LIMIT = -1 IS_TEMPLATE = False;")

		if db.Error != nil {
			return fmt.Errorf(db.Error.Error())
		}

		return nil
	}
	return fmt.Errorf("CreateDB: Unknown database type " + Database.Type)
}

// ClearDB clears the db object
func ClearDB() {
	db = nil
}

// All fetches all object in the database
func All(a interface{}) (err error) {
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Find(a).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Find(a).Error
		}
	})
	if err != nil {
		Trail(ERROR, "DB error in All(%v). %s", getModelName(a), err.Error())
		return err
	}
	decryptArray(a)
	return nil
}

// Save saves the object in the database
func Save(a interface{}) (err error) {
	encryptRecord(a)
	if Database.Type == "mysql" {
		a = fixDates(a)
	}
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Save(a).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Save(a).Error
		}
	})
	if err != nil {
		Trail(ERROR, "DB error in Save(%v). %s", getModelName(a), err.Error())
		return err
	}
	err = customSave(a)
	if err != nil {
		Trail(ERROR, "DB error in customSave(%v). %s", getModelName(a), err.Error())
		return err
	}
	return nil
}

func fixDates(a interface{}) interface{} {
	value := reflect.ValueOf(a).Elem()
	now := time.Now()
	timeType := reflect.TypeOf(now)
	timePointerType := reflect.TypeOf(&now)
	timeValue := reflect.ValueOf(now)
	timePointerValue := reflect.ValueOf(&now)
	for i := 0; i < value.NumField(); i++ {
		if value.Field(i).Type() == timeType {
			if value.Field(i).Interface().(time.Time).IsZero() {
				value.Field(i).Set(timeValue)
			}
		} else if value.Field(i).Type() == timePointerType {
			if value.Field(i).Interface() != nil && value.Field(i).Interface().(*time.Time).IsZero() {
				value.Field(i).Set(timePointerValue)
			}
		}

	}
	return value.Addr().Interface()
}

func customSave(m interface{}) (err error) {
	a := m
	t := reflect.TypeOf(a)
	if t.Kind() == reflect.Ptr {
		a = reflect.ValueOf(m).Elem().Interface()
		t = reflect.TypeOf(a)
	}
	value := reflect.ValueOf(a)
	for i := 0; i < t.NumField(); i++ {
		// Check if there is any m2m fields
		if t.Field(i).Type.Kind() == reflect.Slice && t.Field(i).Type.Elem().Kind() == reflect.Struct {
			table1 := strings.ToLower(t.Name())
			table2 := strings.ToLower(t.Field(i).Type.Elem().Name())

			// Delete existing records
			sql := sqlDialect[Database.Type]["deleteM2M"]
			sql = strings.Replace(sql, "{TABLE1}", table1, -1)
			sql = strings.Replace(sql, "{TABLE2}", table2, -1)

			TimeMetric("uadmin/db/duration", 1000, func() {
				err = db.Exec(sql, GetID(value)).Error
				for fmt.Sprint(err) == "database is locked" {
					time.Sleep(time.Millisecond * 100)
					err = db.Exec(sql, GetID(value)).Error
				}
			})
			if err != nil {
				Trail(ERROR, "Unable to delete m2m records. %s", err)
				Trail(ERROR, sql)
				return err
			}
			// Insert records
			for index := 0; index < value.Field(i).Len(); index++ {
				sql := sqlDialect[Database.Type]["insertM2M"]
				sql = strings.Replace(sql, "{TABLE1}", table1, -1)
				sql = strings.Replace(sql, "{TABLE2}", table2, -1)

				TimeMetric("uadmin/db/duration", 1000, func() {
					err = db.Exec(sql, GetID(value), GetID(value.Field(i).Index(index))).Error
					for fmt.Sprint(err) == "database is locked" {
						time.Sleep(time.Millisecond * 100)
						err = db.Exec(sql, GetID(value), GetID(value.Field(i).Index(index))).Error
					}
				})
				if err != nil {
					Trail(ERROR, "Unable to insert m2m records. %s", err)
					Trail(ERROR, sql)
					return err
				}
			}

		}
	}
	return nil
}

// Get fetches the first record from the database matching query and args
func Get(a interface{}, query interface{}, args ...interface{}) (err error) {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Where(query, args...).First(a).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Where(query, args...).First(a).Error
		}
	})

	if err != nil {
		if err.Error() != "record not found" {
			Trail(ERROR, "DB error in Get(%s)-(%v). %s", getModelName(a), a, err.Error())
		}
		return err
	}

	err = customGet(a)
	if err != nil {
		Trail(ERROR, "DB error in customGet(%v). %s", getModelName(a), err.Error())
		return err
	}
	decryptRecord(a)
	return nil
}

// Get fetches the first record from the database matching query and args with sorting
func GetTable(table string, a interface{}, query interface{}, args ...interface{}) (err error) {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Table(table).Where(query, args...).First(a).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Table(table).Where(query, args...).First(a).Error
		}
	})

	if err != nil {
		if err.Error() != "record not found" {
			Trail(ERROR, "DB error in GetSorted(%s)-(%v). %s", getModelName(a), a, err.Error())
		}
		return err
	}

	err = customGet(a)
	if err != nil {
		Trail(ERROR, "DB error in customGet(%v). %s", getModelName(a), err.Error())
		return err
	}
	decryptRecord(a)
	return nil
}

// Get fetches the first record from the database matching query and args with sorting
func GetSorted(order string, asc bool, a interface{}, query interface{}, args ...interface{}) (err error) {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	if order != "" {
		orderby := " desc"
		if asc {
			orderby = " asc"
		}
		order = columnEnclosure() + order + columnEnclosure()
		orderby += " "
		order += orderby
	} else {
		order = "id desc"
	}
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Where(query, args...).Order(order).First(a).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Where(query, args...).Order(order).First(a).Error
		}
	})

	if err != nil {
		if err.Error() != "record not found" {
			Trail(ERROR, "DB error in GetSorted(%s)-(%v). %s", getModelName(a), a, err.Error())
		}
		return err
	}

	err = customGet(a)
	if err != nil {
		Trail(ERROR, "DB error in customGet(%v). %s", getModelName(a), err.Error())
		return err
	}
	decryptRecord(a)
	return nil
}

// Get fetches the first record from the database matching query and args with sorting
func GetSortedTable(table string, order string, asc bool, a interface{}, query interface{}, args ...interface{}) (err error) {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	if order != "" {
		orderby := " desc"
		if asc {
			orderby = " asc"
		}
		order = columnEnclosure() + order + columnEnclosure()
		orderby += " "
		order += orderby
	} else {
		order = "id desc"
	}
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Table(table).Where(query, args...).Order(order).First(a).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Table(table).Where(query, args...).Order(order).First(a).Error
		}
	})

	if err != nil {
		if err.Error() != "record not found" {
			Trail(ERROR, "DB error in GetSorted(%s)-(%v). %s", getModelName(a), a, err.Error())
		}
		return err
	}

	err = customGet(a)
	if err != nil {
		Trail(ERROR, "DB error in customGet(%v). %s", getModelName(a), err.Error())
		return err
	}
	decryptRecord(a)
	return nil
}

// Get fetches the first record from the database matching query and args with sorting
func GetValueSorted(table string, column string, order string, asc bool, a interface{}, query interface{}, args ...interface{}) (err error) {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	if order != "" {
		orderby := " desc"
		if asc {
			orderby = " asc"
		}
		order = columnEnclosure() + order + columnEnclosure()
		orderby += " "
		order += orderby
	} else {
		order = "id desc"
	}
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Table(table).Where(query, args...).Order(order).Limit(1).Pluck(column, a).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Table(table).Where(query, args...).Order(order).Limit(1).Pluck(column, a).Error
		}
	})

	if err != nil {
		if err.Error() != "record not found" {
			Trail(ERROR, "DB error in GetValueSorted(%s)-(%v). %s", getModelName(a), a, err.Error())
		}
		return err
	}
	return nil
}

// GetABTest is like Get function but implements AB testing for the results
func GetABTest(r *http.Request, a interface{}, query interface{}, args ...interface{}) (err error) {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	TimeMetric("uadmin/db/duration", 1000, func() {
		Get(a, query, args...)
	})

	// Check if there are any active A/B tests for any field in this model
	abt := getABT(r)
	modelName := getModelName(a)
	abTestsMutex.Lock()
	for k, v := range modelABTests {
		if strings.HasPrefix(k, modelName+"__") && strings.HasSuffix(k, "__"+fmt.Sprint(GetID(reflect.ValueOf(a)))) {
			if len(v) != 0 {
				index := abt % len(v)
				fName := Schema[modelName].Fields[v[index].fname].Name

				// TODO: Support more data types
				switch Schema[modelName].Fields[v[index].fname].Type {
				case cSTRING:
					reflect.ValueOf(a).Elem().FieldByName(fName).SetString(v[index].v)
				case cIMAGE:
					reflect.ValueOf(a).Elem().FieldByName(fName).SetString(v[index].v)
				}

				// Increment impressions
				v[index].imp++
				modelABTests[k] = v
			}
		}
	}
	abTestsMutex.Unlock()
	return nil
}

// GetStringer fetches the first record from the database matching query and args
// and get only fields tagged with `stringer` tag. If no field has `stringer` tag
// then it gets all the fields
func GetStringer(a interface{}, query interface{}, args ...interface{}) (err error) {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	stringers := []string{}
	modelName := getModelName(a)
	for _, f := range Schema[modelName].Fields {
		if f.Stringer {
			stringers = append(stringers, GetDB().Config.NamingStrategy.ColumnName("", f.Name))
		}
	}
	TimeMetric("uadmin/db/duration", 1000, func() {
		if len(stringers) == 0 {
			err = db.Where(query, args...).First(a).Error
			for fmt.Sprint(err) == "database is locked" {
				time.Sleep(time.Millisecond * 100)
				err = db.Where(query, args...).First(a).Error
			}
		} else {
			err = db.Select(stringers).Where(query, args...).First(a).Error
			for fmt.Sprint(err) == "database is locked" {
				time.Sleep(time.Millisecond * 100)
				err = db.Select(stringers).Where(query, args...).First(a).Error
			}
		}
	})
	if err != nil {
		if err.Error() != "record not found" {
			Trail(ERROR, "DB error in Get(%s)-(%v). %s", getModelName(a), a, err.Error())
		}
		return err
	}

	//err = customGet(a)
	if err != nil {
		Trail(ERROR, "DB error in customGet(%v). %s", getModelName(a), err.Error())
		return err
	}
	decryptRecord(a)
	return nil
}

// GetForm fetches the first record from the database matching query and args
// where it selects only visible fields in the form based on given schema
func GetForm(a interface{}, s *ModelSchema, query interface{}, args ...interface{}) (err error) {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	// get a list of visible fields
	columnList := []string{}
	m2mList := []string{}
	for _, f := range s.Fields {
		if !f.Hidden {
			if f.Type == cM2M {
				m2mList = append(m2mList, f.ColumnName)
			} else if f.Type == cFK {
				columnList = append(columnList, columnEnclosure()+f.ColumnName+"_id"+columnEnclosure())
				// } else if f.IsMethod {
			} else {
				columnList = append(columnList, columnEnclosure()+f.ColumnName+columnEnclosure())
			}
		}
	}

	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Select(columnList).Where(query, args...).First(a).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Select(columnList).Where(query, args...).First(a).Error
		}
	})

	if err != nil {
		if err.Error() != "record not found" {
			Trail(ERROR, "DB error in Get(%s)-(%v). %s", getModelName(a), a, err.Error())
		}
		return err
	}

	err = customGet(a, m2mList...)
	if err != nil {
		Trail(ERROR, "DB error in customGet(%v). %s", getModelName(a), err.Error())
		return err
	}
	decryptRecord(a)
	return nil
}

func customGet(m interface{}, m2m ...string) (err error) {
	a := m
	t := reflect.TypeOf(a)
	var ignore bool
	if t.Kind() == reflect.Ptr {
		a = reflect.ValueOf(m).Elem().Interface()
		t = reflect.TypeOf(a)
	}
	value := reflect.ValueOf(a)
	for i := 0; i < t.NumField(); i++ {
		ignore = false
		if len(m2m) != 0 {
			ignore = true
			for _, fName := range m2m {
				if fName == t.Field(i).Name {
					ignore = false
					break
				}
			}
		}
		if ignore {
			continue
		}

		// Skip private fields
		if t.Field(i).Anonymous || t.Field(i).Name[0:1] != strings.ToUpper(t.Field(i).Name[0:1]) {
			continue
		}

		// Check if there is any m2m fields
		if t.Field(i).Type.Kind() == reflect.Slice && t.Field(i).Type.Elem().Kind() == reflect.Struct {
			table1 := strings.ToLower(t.Name())
			table2 := strings.ToLower(t.Field(i).Type.Elem().Name())

			sqlSelect := sqlDialect[Database.Type]["selectM2M"]
			sqlSelect = strings.Replace(sqlSelect, "{TABLE1}", table1, -1)
			sqlSelect = strings.Replace(sqlSelect, "{TABLE2}", table2, -1)

			var rows *sql.Rows
			rows, err = db.Raw(sqlSelect, GetID(value)).Rows()
			if err != nil {
				Trail(ERROR, "Unable to get m2m records. %s", err)
				Trail(ERROR, sqlSelect)
				return err
			}
			defer rows.Close()
			var fkID uint
			tmpDst := reflect.New(reflect.SliceOf(t.Field(i).Type.Elem())).Elem()
			for rows.Next() {
				rows.Scan(&fkID)
				tempModel := reflect.New(t.Field(i).Type.Elem()).Elem()

				Get(tempModel.Addr().Interface(), "id = ?", fkID)
				tmpDst = reflect.Append(tmpDst, tempModel)
			}
			reflect.ValueOf(m).Elem().Field(i).Set(tmpDst)
		}
	}
	return nil
}

// Filter fetches records from the database
func Filter(a interface{}, query interface{}, args ...interface{}) (err error) {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Where(query, args...).Find(a).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Where(query, args...).Find(a).Error
		}
	})

	if err != nil {
		Trail(ERROR, "DB error in Filter(%v). %s\n", getModelName(a), err.Error())
		return err
	}
	decryptArray(a)
	return nil
}

// Filter fetches records from the database
func FilterSorted(order string, asc bool, a interface{}, query interface{}, args ...interface{}) (err error) {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	if order != "" {
		orderby := " desc"
		if asc {
			orderby = " asc"
		}
		order = columnEnclosure() + order + columnEnclosure()
		orderby += " "
		order += orderby
	} else {
		order = "id desc"
	}
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Where(query, args...).Order(order).Find(a).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Where(query, args...).Order(order).Find(a).Error
		}
	})

	if err != nil {
		Trail(ERROR, "DB error in Filter(%v). %s\n", getModelName(a), err.Error())
		return err
	}
	decryptArray(a)
	return nil
}

// Filter fetches records from the database
func FilterTable(table string, a interface{}, query interface{}, args ...interface{}) (err error) {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Table(table).Where(query, args...).Find(a).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Table(table).Where(query, args...).Find(a).Error
		}
	})

	if err != nil {
		Trail(ERROR, "DB error in Filter(%v). %s\n", getModelName(a), err.Error())
		return err
	}
	decryptArray(a)
	return nil
}

// Filter fetches records from the database
func FilterSortedTable(table string, order string, asc bool, a interface{}, query interface{}, args ...interface{}) (err error) {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	if order != "" {
		orderby := " desc"
		if asc {
			orderby = " asc"
		}
		order = columnEnclosure() + order + columnEnclosure()
		orderby += " "
		order += orderby
	} else {
		order = "id desc"
	}
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Table(table).Where(query, args...).Order(order).Find(a).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Table(table).Where(query, args...).Order(order).Find(a).Error
		}
	})

	if err != nil {
		Trail(ERROR, "DB error in Filter(%v). %s\n", getModelName(a), err.Error())
		return err
	}
	decryptArray(a)
	return nil
}

// Filter fetches records from the database
func FilterSortedValue(table string, column string, order string, asc bool, a interface{}, query interface{}, args ...interface{}) (err error) {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	if order != "" {
		orderby := " desc"
		if asc {
			orderby = " asc"
		}
		order = columnEnclosure() + order + columnEnclosure()
		orderby += " "
		order += orderby
	} else {
		order = "id desc"
	}
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Table(table).Where(query, args...).Order(order).Pluck(column, a).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Table(table).Where(query, args...).Order(order).Pluck(column, a).Error
		}
	})

	if err != nil {
		Trail(ERROR, "DB error in Filter(%v). %s\n", getModelName(a), err.Error())
		return err
	}
	decryptArray(a)
	return nil
}

// Preload fills the data from foreign keys into structs. You can pass in preload alist of fields
// to be preloaded. If nothing is passed, every foreign key is preloaded
func Preload(a interface{}, preload ...string) (err error) {
	modelName := strings.ToLower(reflect.TypeOf(a).Elem().Name())
	if len(preload) == 0 {
		if schema, ok := getSchema(modelName); ok {
			for _, f := range schema.Fields {
				if f.Type == "fk" {
					preload = append(preload, f.Name)
				}
			}
		} else {
			Trail(ERROR, "DB.Preload No model named %s", modelName)
			return fmt.Errorf("DB.Preload No model named %s", modelName)
		}
	}
	value := reflect.ValueOf(a).Elem()
	for _, p := range preload {
		fkType := value.FieldByName(p).Type().Name()
		if value.FieldByName(p).Type().Kind() == reflect.Ptr {
			fkType = value.FieldByName(p).Type().Elem().Name()
		}
		fieldStruct, _ := NewModel(strings.ToLower(fkType), true)
		TimeMetric("uadmin/db/duration", 1000, func() {
			err = db.Where("id = ?", value.FieldByName(p+"ID").Interface()).First(fieldStruct.Interface()).Error
			for fmt.Sprint(err) == "database is locked" {
				time.Sleep(time.Millisecond * 100)
				err = db.Where("id = ?", value.FieldByName(p+"ID").Interface()).First(fieldStruct.Interface()).Error
			}
		})

		//		err = Get(fieldStruct.Interface(), "id = ?", value.FieldByName(p+"ID").Interface())
		if err != nil && err.Error() != "record not found" {
			Trail(ERROR, "DB error in Preload(%s).%s %s\n", modelName, p, err.Error())
			return err
		}
		if GetID(fieldStruct) != 0 {
			if value.FieldByName(p).Type().Kind() == reflect.Ptr {
				value.FieldByName(p).Set(fieldStruct)
			} else {
				value.FieldByName(p).Set(fieldStruct.Elem())
			}
		}
	}
	return customGet(a)
}

// Delete records from database
func Delete(a interface{}) (err error) {
	v := reflect.ValueOf(a)
	t := reflect.TypeOf(a)

	// Sanity Check for ID = 0
	if GetID(v) == 0 {
		return nil
	}
	TimeMetric("uadmin/db/duration", 1000, func() {
		if t.Kind() == reflect.Ptr {
			err = db.Delete(a).Error
		} else {
			vp := reflect.New(t)
			vp.Elem().Set(v)
			err = db.Delete(vp.Interface()).Error
		}

		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Delete(a).Error
		}
	})

	if err != nil {
		Trail(ERROR, "DB error in Delete(%v). %s\n", getModelName(a), err.Error())
		return err
	}
	return nil
}

// DeleteList deletes multiple records from database
func DeleteList(a interface{}, query interface{}, args ...interface{}) (err error) {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	v := reflect.ValueOf(a)
	t := reflect.TypeOf(a)

	TimeMetric("uadmin/db/duration", 1000, func() {
		if t.Kind() == reflect.Ptr {
			err = db.Where(query, args...).Delete(a).Error
		} else {
			vp := reflect.New(t)
			vp.Elem().Set(v)
			err = db.Where(query, args...).Delete(vp.Interface()).Error
		}
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			if t.Kind() == reflect.Ptr {
				err = db.Where(query, args...).Delete(a).Error
			} else {
				vp := reflect.New(t)
				vp.Elem().Set(v)
				err = db.Where(query, args...).Delete(vp.Interface()).Error
			}
		}
	})

	if err != nil {
		Trail(ERROR, "DB error in DeleteList(%v). %s\n", getModelName(a), err.Error())
		return err
	}
	return nil
}

// FilterBuilder changes a map filter into a query
func FilterBuilder(params map[string]interface{}) (query string, args []interface{}) {
	keys := []string{}
	for key, value := range params {
		keys = append(keys, key)
		args = append(args, value)
	}
	query = strings.Join(keys, " AND ")
	return
}

// AdminPage !
func AdminPage(order string, asc bool, offset int, limit int, a interface{}, query interface{}, args ...interface{}) (err error) {
	if order != "" {
		orderby := " desc"
		if asc {
			orderby = " asc"
		}
		order = columnEnclosure() + order + columnEnclosure()
		orderby += " "
		order += orderby
	} else {
		order = "id desc"
	}
	if limit > 0 {
		TimeMetric("uadmin/db/duration", 1000, func() {
			err = db.Where(query, args...).Order(order).Offset(offset).Limit(limit).Find(a).Error
			for fmt.Sprint(err) == "database is locked" {
				time.Sleep(time.Millisecond * 100)
				err = db.Where(query, args...).Order(order).Offset(offset).Limit(limit).Find(a).Error
			}
		})

		if err != nil {
			Trail(ERROR, "DB error in AdminPage(%v). %s\n", getModelName(a), err.Error())
			return err
		}
		decryptArray(a)
		return nil
	}
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Where(query, args...).Order(order).Find(a).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Where(query, args...).Order(order).Find(a).Error
		}
	})

	if err != nil {
		Trail(ERROR, "DB error in AdminPage(%v). %s\n", getModelName(a), err.Error())
		return err
	}
	decryptArray(a)
	return nil
}

// FilterList fetches the all record from the database matching query and args
// where it selects only visible fields in the form based on given schema
func FilterList(s *ModelSchema, order string, asc bool, offset int, limit int, a interface{}, query interface{}, args ...interface{}) (err error) {
	// get a list of visible fields
	columnList := []string{}
	for _, f := range s.Fields {
		if f.ListDisplay {
			if f.Type == cFK {
				columnList = append(columnList, columnEnclosure()+GetDB().Config.NamingStrategy.ColumnName("", f.Name)+"_id"+columnEnclosure())
			} else if f.Type == cM2M {
			} else if f.IsMethod {
			} else {
				columnList = append(columnList, columnEnclosure()+GetDB().Config.NamingStrategy.ColumnName("", f.Name)+columnEnclosure())
			}
		}
	}
	if order != "" {
		orderby := " desc"
		if asc {
			orderby = " asc"
		}
		order = columnEnclosure() + order + columnEnclosure()
		orderby += " "
		order += orderby
	} else {
		order = "id desc"
	}
	if limit > 0 {
		TimeMetric("uadmin/db/duration", 1000, func() {
			err = db.Select(columnList).Where(query, args...).Order(order).Offset(offset).Limit(limit).Find(a).Error
			for fmt.Sprint(err) == "database is locked" {
				time.Sleep(time.Millisecond * 100)
				err = db.Select(columnList).Where(query, args...).Order(order).Offset(offset).Limit(limit).Find(a).Error
			}
		})

		if err != nil {
			Trail(ERROR, "DB error in FilterList(%v) query:%s, args(%#v). %s\n", getModelName(a), query, args, err.Error())
			return err
		}
		decryptArray(a)
		return nil
	}
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Select(columnList).Where(query, args...).Order(order).Find(a).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Select(columnList).Where(query, args...).Order(order).Find(a).Error
		}
	})

	if err != nil {
		Trail(ERROR, "DB error in FilterList(%v) query:%s, args(%#v). %s\n", getModelName(a), query, args, err.Error())
		return err
	}
	decryptArray(a)
	return nil
}

// Count return the count of records in a table based on a filter
func Count(a interface{}, query interface{}, args ...interface{}) int {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	var count int64
	var err error
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Model(a).Where(query, args...).Count(&count).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Model(a).Where(query, args...).Count(&count).Error
		}
	})

	if err != nil && strings.Contains(err.Error(), "NULL") {
		err = nil
	}

	if err != nil {
		Trail(ERROR, "DB error in Count(%v). %s\n", getModelName(a), err.Error())
	}
	return int(count)
}

// Sum return the sum of a column in a table based on a filter
func Sum(a interface{}, column string, query interface{}, args ...interface{}) float64 {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	var vals []float64
	var err error
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Model(a).Where(query, args...).Select("SUM("+column+")").Pluck("SUM("+column+")", &vals).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Model(a).Where(query, args...).Select("SUM("+column+")").Pluck("SUM("+column+")", &vals).Error
		}
	})

	if err != nil && strings.Contains(err.Error(), "NULL") {
		err = nil
	}

	if err != nil {
		Trail(ERROR, "DB error in Sum(%v). %s\n", getModelName(a), err.Error())
	}
	if len(vals) == 0 {
		return 0.0
	}
	return vals[0]
}

// Avg return the average of a column in a table based on a filter
func Avg(a interface{}, column string, query interface{}, args ...interface{}) float64 {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	var vals []float64
	var err error
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Model(a).Where(query, args...).Select("AVG("+column+")").Pluck("AVG("+column+")", &vals).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Model(a).Where(query, args...).Select("AVG("+column+")").Pluck("AVG("+column+")", &vals).Error
		}
	})

	if err != nil && strings.Contains(err.Error(), "NULL") {
		err = nil
	}

	if err != nil {
		Trail(ERROR, "DB error in Avg(%v). %s\n", getModelName(a), err.Error())
	}
	if len(vals) == 0 {
		return 0.0
	}
	return vals[0]
}

// Max return the maximum of a column in a table based on a filter
func Max(a interface{}, column string, query interface{}, args ...interface{}) float64 {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	var vals []float64
	var err error
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Model(a).Where(query, args...).Select("MAX("+column+")").Pluck("MAX("+column+")", &vals).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Model(a).Where(query, args...).Select("MAX("+column+")").Pluck("MAX("+column+")", &vals).Error
		}
	})

	if err != nil && strings.Contains(err.Error(), "NULL") {
		err = nil
	}

	if err != nil {
		Trail(ERROR, "DB error in Max(%v). %s\n", getModelName(a), err.Error())
	}
	if len(vals) == 0 {
		return 0.0
	}
	return vals[0]
}

// Min return the minimum of a column in a table based on a filter
func Min(a interface{}, column string, query interface{}, args ...interface{}) float64 {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	var vals []float64
	var err error
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Model(a).Where(query, args...).Select("MIN("+column+")").Pluck("MIN("+column+")", &vals).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Model(a).Where(query, args...).Select("MIN("+column+")").Pluck("MIN("+column+")", &vals).Error
		}
	})

	if err != nil && strings.Contains(err.Error(), "NULL") {
		err = nil
	}

	if err != nil {
		Trail(ERROR, "DB error in Min(%v). %s\n", getModelName(a), err.Error())
	}
	if len(vals) == 0 {
		return 0.0
	}
	return vals[0]
}

// Std return the standard diviation of a column in a table based on a filter
func Std(a interface{}, column string, query interface{}, args ...interface{}) float64 {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	var vals []float64
	var err error
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Model(a).Where(query, args...).Select("STD("+column+")").Pluck("STD("+column+")", &vals).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Model(a).Where(query, args...).Select("STD("+column+")").Pluck("STD("+column+")", &vals).Error
		}
	})

	if err != nil && strings.Contains(err.Error(), "NULL") {
		err = nil
	}

	if err != nil {
		Trail(ERROR, "DB error in Std(%v). %s\n", getModelName(a), err.Error())
	}
	if len(vals) == 0 {
		return 0.0
	}
	return vals[0]
}

// CountTable return the count of records in a table based on a filter
func CountTable(table string, query interface{}, args ...interface{}) int {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	var count int64
	var err error
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Table(table).Where(query, args...).Count(&count).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Table(table).Where(query, args...).Count(&count).Error
		}
	})

	if err != nil && strings.Contains(err.Error(), "NULL") {
		err = nil
	}

	if err != nil {
		Trail(ERROR, "DB error in CountTable(%v). %s\n", table, err.Error())
	}
	return int(count)
}

// SumTable return the sum of a column in a table based on a filter
func SumTable(table string, column string, query interface{}, args ...interface{}) float64 {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	var vals []float64
	var err error
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Table(table).Where(query, args...).Select("SUM("+column+")").Pluck("SUM("+column+")", &vals).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Table(table).Where(query, args...).Select("SUM("+column+")").Pluck("SUM("+column+")", &vals).Error
		}
	})

	if err != nil && strings.Contains(err.Error(), "NULL") {
		err = nil
	}

	if err != nil {
		Trail(ERROR, "DB error in SumTable(%v). %s\n", table, err.Error())
	}
	if len(vals) == 0 {
		return 0.0
	}
	return vals[0]
}

// AvgTable return the average of a column in a table based on a filter
func AvgTable(table string, column string, query interface{}, args ...interface{}) float64 {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	var vals []float64
	var err error
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Table(table).Where(query, args...).Select("AVG("+column+")").Pluck("AVG("+column+")", &vals).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Table(table).Where(query, args...).Select("AVG("+column+")").Pluck("AVG("+column+")", &vals).Error
		}
	})

	if err != nil && strings.Contains(err.Error(), "NULL") {
		err = nil
	}

	if err != nil {
		Trail(ERROR, "DB error in AvgTable(%v). %s\n", table, err.Error())
	}
	if len(vals) == 0 {
		return 0.0
	}
	return vals[0]
}

// MaxTable return the maximum of a column in a table based on a filter
func MaxTable(table string, column string, query interface{}, args ...interface{}) float64 {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	var vals []float64
	var err error
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Table(table).Where(query, args...).Select("MAX("+column+")").Pluck("MAX("+column+")", &vals).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Table(table).Where(query, args...).Select("MAX("+column+")").Pluck("MAX("+column+")", &vals).Error
		}
	})

	if err != nil && strings.Contains(err.Error(), "NULL") {
		err = nil
	}

	if err != nil {
		Trail(ERROR, "DB error in MaxTable(%v). %s\n", table, err.Error())
	}
	if len(vals) == 0 {
		return 0.0
	}
	return vals[0]
}

// MinTable return the minimum of a column in a table based on a filter
func MinTable(table string, column string, query interface{}, args ...interface{}) float64 {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	var vals []float64
	var err error
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Table(table).Where(query, args...).Select("MIN("+column+")").Pluck("MIN("+column+")", &vals).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Table(table).Where(query, args...).Select("MIN("+column+")").Pluck("MIN("+column+")", &vals).Error
		}
	})

	if err != nil && strings.Contains(err.Error(), "NULL") {
		err = nil
	}

	if err != nil {
		Trail(ERROR, "DB error in MinTable(%v). %s\n", table, err.Error())
	}
	if len(vals) == 0 {
		return 0.0
	}
	return vals[0]
}

// StdTable return the standard diviation of a column in a table based on a filter
func StdTable(table string, column string, query interface{}, args ...interface{}) float64 {
	if val, ok := query.(string); ok {
		query = fixQueryEnclosure(val)
	}
	var vals []float64
	var err error
	TimeMetric("uadmin/db/duration", 1000, func() {
		err = db.Table(table).Where(query, args...).Select("STD("+column+")").Pluck("STD("+column+")", &vals).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			err = db.Table(table).Where(query, args...).Select("STD("+column+")").Pluck("STD("+column+")", &vals).Error
		}
	})

	if err != nil && strings.Contains(err.Error(), "NULL") {
		err = nil
	}

	if err != nil {
		Trail(ERROR, "DB error in StdTable(%v). %s\n", table, err.Error())
	}
	if len(vals) == 0 {
		return 0.0
	}
	return vals[0]
}

// Update !
func Update(a interface{}, fieldName string, value interface{}, query string, args ...interface{}) (err error) {
	query = fixQueryEnclosure(query)
	TimeMetric("uadmin/db/duration", 1000, func() {
		// There seems to be a bug in gorm stopping updates using model but it works when we use table
		//err = db.Model(a).Where(query, args...).Update(fieldName, value).Error
		//tableName := db.Config.NamingStrategy.TableName(reflect.TypeOf(a).Elem().Name())
		tableName := getModelNameNorm(a)
		tableName = db.Config.NamingStrategy.TableName(tableName)

		err = db.Table(tableName).Where(query, args...).Update(fieldName, value).Error
		for fmt.Sprint(err) == "database is locked" {
			time.Sleep(time.Millisecond * 100)
			//err = db.Model(a).Where(query, args...).Update(fieldName, value).Error
			err = db.Table(tableName).Where(query, args...).Update(fieldName, value).Error
		}
	})

	if err != nil {
		Trail(ERROR, "DB error in Update(%v). %s\n", getModelName(a), err.Error())
	}
	return nil
}

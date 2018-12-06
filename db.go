package uadmin

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mssql"

	// Enable MYSQL
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/jinzhu/gorm/dialects/postgres"

	// Enable SQLLite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/uadmin/uadmin/colors"
)

var db *gorm.DB

var sqlDialect = map[string]map[string]string{
	"mysql": map[string]string{
		"createM2MTable": "CREATE TABLE `{TABLE1}_{TABLE2}` (`table1_id` int(10) unsigned NOT NULL, `table2_id` int(10) unsigned NOT NULL, PRIMARY KEY (`table1_id`,`table2_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;",
		"selectM2M":      "SELECT `table2_id` FROM `{TABLE1}_{TABLE2}` WHERE table1_id={TABLE1_ID};",
		"deleteM2M":      "DELETE FROM `{TABLE1}_{TABLE2}` WHERE `table1_id`={TABLE1_ID};",
		"insertM2M":      "INSERT INTO `{TABLE1}_{TABLE2}` VALUES ({TABLE1_ID}, {TABLE2_ID});",
	},
	"sqlite": map[string]string{
		//"createM2MTable": "CREATE TABLE `{TABLE1}_{TABLE2}` (`{TABLE1}_id`	INTEGER NOT NULL,`{TABLE2}_id` INTEGER NOT NULL, PRIMARY KEY(`{TABLE1}_id`,`{TABLE2}_id`));",
		"createM2MTable": "CREATE TABLE `{TABLE1}_{TABLE2}` (`table1_id`	INTEGER NOT NULL,`table2_id` INTEGER NOT NULL, PRIMARY KEY(`table1_id`,`table2_id`));",
		"selectM2M": "SELECT `table2_id` FROM `{TABLE1}_{TABLE2}` WHERE table1_id={TABLE1_ID};",
		"deleteM2M": "DELETE FROM `{TABLE1}_{TABLE2}` WHERE `table1_id`={TABLE1_ID};",
		"insertM2M": "INSERT INTO `{TABLE1}_{TABLE2}` VALUES ({TABLE1_ID}, {TABLE2_ID});",
	},
}

// Database is the active Database settings
var Database *DBSettings

// DBSettings !
type DBSettings struct {
	Type     string // SQLLite, MySQL
	Name     string // File/DB name
	User     string
	Password string
	Host     string
	Port     int
}

// InitializeDB opens the connection the DB
func initializeDB(a ...interface{}) {
	// Open the connection the the DB
	if Database == nil {
		Database = &DBSettings{
			Type: "sqlite",
		}
	}
	db = GetDB()
	// Migrate schema
	for i, model := range a {
		db.AutoMigrate(model)
		Trail(WORKING, "Initializing DB: [%s%d/%d%s]", colors.FG_GREEN_B, i+1, len(a), colors.FG_NORMAL)
		customMigration(model)
	}
	Trail(OK, "Initializing DB: [%s%d/%d%s]", colors.FG_GREEN_B, len(a), len(a), colors.FG_NORMAL)
	if DebugDB {
		db.LogMode(true)
	}
}

func customMigration(a interface{}) (err error) {
	t := reflect.TypeOf(a)
	for i := 0; i < t.NumField(); i++ {
		// Check if there is any m2m fields
		if t.Field(i).Type.Kind() == reflect.Slice && t.Field(i).Type.Elem().Kind() == reflect.Struct {
			table1 := strings.ToLower(t.Name())
			table2 := strings.ToLower(t.Field(i).Type.Elem().Name())

			//Check if the table is created for the m2m field
			if !db.HasTable(table1 + "_" + table2) {
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

// AdminModel !
type adminModel interface {
	String() string
	GetID() uint
}

// GetDB returns a pointer to the DB
func GetDB() *gorm.DB {
	if db != nil {
		return db
	}
	var err error
	if strings.ToLower(Database.Type) == "sqlite" {
		dbName := Database.Name
		if dbName == "" {
			dbName = "uadmin.db"
		}
		db, err = gorm.Open("sqlite3", dbName)
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
		dsn := fmt.Sprintf("%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			credential,
			Database.Host,
			Database.Port,
			Database.Name,
		)
		db, err = gorm.Open("mysql", dsn)
	}

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	return db
}

// ClearDB clears the db object
func ClearDB() {
	db = nil
}

// CloseDB closes the connection to the DB
func closeDB() {
	db.Close()
}

// All fetches all object in the database
func All(a interface{}) (err error) {
	err = db.Find(a).Error
	if err != nil {
		Trail(ERROR, "DB error in All(%v). %s", reflect.TypeOf(a).Name(), err.Error())
		return err
	}
	decryptArray(a)
	return nil
}

// allTable fetches all object in the database for a specific table
func allTable(a interface{}, table string) {
	//db.Find(a)
	db.Table(table).Find(a)
}

// Save saves the object in the database
func Save(a interface{}) (err error) {
	encryptRecord(a)
	err = db.Save(a).Error
	if err != nil {
		Trail(ERROR, "DB error in Save(%v). %s", reflect.TypeOf(a).Name(), err.Error())
		return err
	}
	err = customSave(a)
	if err != nil {
		Trail(ERROR, "DB error in customSave(%v). %s", reflect.TypeOf(a).Name(), err.Error())
		return err
	}
	return nil
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
			sql = strings.Replace(sql, "{TABLE1_ID}", fmt.Sprint(GetID(value)), -1)
			err = db.Exec(sql).Error
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
				sql = strings.Replace(sql, "{TABLE1_ID}", fmt.Sprint(GetID(value)), -1)
				sql = strings.Replace(sql, "{TABLE2_ID}", fmt.Sprint(GetID(value.Field(i).Index(index))), -1)
				err = db.Exec(sql).Error
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

// Get fetches the first record from the database
func Get(a interface{}, query interface{}, args ...interface{}) (err error) {
	err = db.Where(query, args...).First(a).Error
	if err != nil {
		if err.Error() != "record not found" {
			Trail(ERROR, "DB error in Get(%s)-(%v). %s", reflect.TypeOf(a).Name(), a, err.Error())
		}
		return err
	}

	err = customGet(a)
	if err != nil {
		Trail(ERROR, "DB error in customGet(%v). %s", reflect.TypeOf(a).Name(), err.Error())
		return err
	}
	decryptRecord(a)
	return nil
}

func customGet(m interface{}) (err error) {
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

			sqlSelect := sqlDialect[Database.Type]["selectM2M"]
			sqlSelect = strings.Replace(sqlSelect, "{TABLE1}", table1, -1)
			sqlSelect = strings.Replace(sqlSelect, "{TABLE2}", table2, -1)
			sqlSelect = strings.Replace(sqlSelect, "{TABLE1_ID}", fmt.Sprint(GetID(value)), -1)

			var rows *sql.Rows
			rows, err = db.Raw(sqlSelect).Rows()
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
	err = db.Where(query, args...).Find(a).Error
	if err != nil {
		Trail(ERROR, "DB error in Filter(%v). %s\n", reflect.TypeOf(a).Name(), err.Error())
		return err
	}
	decryptArray(a)
	return nil
}

// Preload !
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
		err = db.Where("id = ?", value.FieldByName(p+"ID").Interface()).First(fieldStruct.Interface()).Error
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

// preloadFilter !
func preloadFilter(a interface{}, preload []string, query interface{}, args ...interface{}) {
	db = db.Where(query, args)
	for _, p := range preload {
		db = db.Preload(p)
	}
	db.Find(a)
}

// Delete records from database
func Delete(a interface{}) (err error) {
	err = db.Delete(a).Error
	if err != nil {
		Trail(ERROR, "DB error in Delete(%v). %s\n", reflect.TypeOf(a).Name(), err.Error())
		return err
	}
	return nil
}

// DeleteList !
func DeleteList(a interface{}, query interface{}, args ...interface{}) (err error) {
	err = db.Where(query, args...).Delete(a).Error
	if err != nil {
		Trail(ERROR, "DB error in DeleteList(%v). %s\n", reflect.TypeOf(a).Name(), err.Error())
		return err
	}
	return nil
}

// FilterBuilder chnages a map filter into a query
func FilterBuilder(params map[string]interface{}) (query string, args []interface{}) {
	keys := []string{}
	for key, value := range params {
		keys = append(keys, key)
		args = append(args, value)
	}
	query = strings.Join(keys, " AND ")
	return
}

// OrderedWhere !
func orderedWhere(order string, a interface{}, query interface{}, args ...interface{}) {
	db.Where(query, args...).Find(a)
}

// Where get list of menuitems used by the api
func where(a interface{}, params map[string]interface{}) {
	db.Find(a)
	return
}

// AdminPage !
func AdminPage(order string, asc bool, offset int, limit int, a interface{}, query interface{}, args ...interface{}) (err error) {
	if order != "" {
		orderby := " desc"
		if asc {
			orderby = " asc"
		}
		order = "`" + order + "`"
		orderby += " "
		order += orderby
	} else {
		order = "id desc"
	}
	if limit > 1 {
		err = db.Where(query, args...).Order(order).Offset(offset).Limit(limit).Find(a).Error
		if err != nil {
			Trail(ERROR, "DB error in AdminPage(%v). %s\n", reflect.TypeOf(a).Name(), err.Error())
			return err
		}
		decryptArray(a)
		return nil
	}
	err = db.Where(query, args...).Order(order).Find(a).Error
	if err != nil {
		Trail(ERROR, "DB error in AdminPage(%v). %s\n", reflect.TypeOf(a).Name(), err.Error())
		return err
	}
	decryptArray(a)
	return nil
}

// Count return the count of records in a table based on a filter
func Count(a interface{}, query interface{}, args ...interface{}) int {
	var count int
	err := db.Model(a).Where(query, args...).Count(&count).Error
	if err != nil {
		Trail(ERROR, "DB error in Count(%v). %s\n", reflect.TypeOf(a).Name(), err.Error())
	}
	return count
}

// Update !
func Update(a interface{}, fieldName string, value interface{}, query string, args ...interface{}) (err error) {
	err = db.Model(a).Where(query, args...).Update(fieldName, value).Error
	if err != nil {
		Trail(ERROR, "DB error in Update(%v). %s\n", reflect.TypeOf(a).Name(), err.Error())
	}
	return nil
}

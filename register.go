package uadmin

import (
	"crypto/sha512"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/jinzhu/inflection"
	"github.com/uadmin/uadmin/helper"
)

// HideInDashboarder used to check if a model should be hidden in
// dashboard
type HideInDashboarder interface {
	HideInDashboard() bool
}

// SchemaCategory provides a default category for the model. This can be
// customized later from the UI
type SchemaCategory interface {
	SchemaCategory() string
}

// CustomTranslation is where you can register custom translation files.
// To register a custom translation file, always assign it with it's key
// in the this format "category/name". For example:
//
//	uadmin.CustomTranslation = append(uadmin.CustomTranslation, "ui/billing")
//
// This will register the file and you will be able to use it if `uadmin.Tf`.
// By default there is only one registered custom translation which is "uadmin/system".
var CustomTranslation = []string{
	"uadmin/system",
}

var modelList []interface{}

// Register is used to register models to uadmin
func Register(m ...interface{}) {
	modelList = []interface{}{}

	if len(models) == 0 {
		models = map[string]interface{}{}

		// Initialize system models
		modelList = []interface{}{
			DashboardMenu{},
			User{},
			UserGroup{},
			Session{},
			UserPermission{},
			GroupPermission{},
			Language{},
			Log{},
			Setting{},
			SettingCategory{},
			Approval{},
			ABTest{},
			ABTestValue{},
			//Builder{},
			//BuilderField{},
		}
	}

	// System models count
	SMCount := len(modelList)

	// Now add user defined models
	modelList = append(modelList,
		m...,
	)

	// Initialize the Database
	initializeDB(modelList...)

	// Setup languages
	initializeLanguage()

	// check if trail dashboard menu item is added
	if Count([]DashboardMenu{}, "menu_name = ?", "Trail") == 0 {
		dashboard := DashboardMenu{
			MenuName: "Trail",
			URL:      "trail",
			Hidden:   false,
			Cat:      "System",
		}
		Save(&dashboard)
	}

	// Store models in Model global variable
	// and initialize the dashboard
	dashboardMenus := []DashboardMenu{}
	All(&dashboardMenus)
	var modelExists bool
	Schema = map[string]ModelSchema{}
	for i := range modelList {
		modelExists = false
		t := reflect.TypeOf(modelList[i])
		name := strings.ToLower(t.Name())
		models[name] = modelList[i]

		// Get Hidden model status
		hideItem := false
		if hider, ok := modelList[i].(HideInDashboarder); ok {
			hideItem = hider.HideInDashboard()
		}

		// Get Category Name
		cat := "System"
		// Check if the model is a system model
		if i >= SMCount {
			if category, ok := modelList[i].(SchemaCategory); ok {
				cat = category.SchemaCategory()
			} else {
				cat = ""
			}
		}

		// Register Dashboard menu
		// First check if the model is already in dashboard
		dashboardIndex := 0
		for index, val := range dashboardMenus {
			if name == val.URL {
				modelExists = true
				dashboardIndex = index
				break
			}
		}

		// If not in dashboard, then add it
		if !modelExists {
			dashboard := DashboardMenu{
				MenuName: inflection.Plural(strings.Join(helper.SplitCamelCase(t.Name()), " ")),
				URL:      name,
				Hidden:   hideItem,
				Cat:      cat,
			}
			Save(&dashboard)
		} else {
			// If model exists, synchronize it if changed
			if hideItem != dashboardMenus[dashboardIndex].Hidden {
				dashboardMenus[dashboardIndex].Hidden = hideItem
				Save(&dashboardMenus[dashboardIndex])
			}
			if cat != dashboardMenus[dashboardIndex].Cat {
				dashboardMenus[dashboardIndex].Cat = cat
				Save(&dashboardMenus[dashboardIndex])
			}
		}
	}

	// Check if encrypt key is there or generate it
	if _, err := os.Stat(".key"); os.IsNotExist(err) && os.Getenv("UADMIN_KEY") == "" {
		EncryptKey = generateByteArray(32)
		ioutil.WriteFile(".key", EncryptKey, 0600)
	} else {
		EncryptKey = []byte(os.Getenv("UADMIN_KEY"))
		if len(EncryptKey) == 0 {
			EncryptKey, _ = ioutil.ReadFile(".key")
		}
	}

	// Check if JWT key is there or generate it
	if _, err := os.Stat(".jwt"); os.IsNotExist(err) && os.Getenv("UADMIN_JWT") == "" {
		JWT = GenerateBase64(64)
		ioutil.WriteFile(".jwt", []byte(JWT), 0600)
	} else {
		JWT = os.Getenv("UADMIN_JWT")
		if len(JWT) == 0 {
			buf, _ := ioutil.ReadFile(".jwt")
			JWT = string(buf)
		}
	}
	JWTIssuer = func() string {
		hash := sha512.New()
		hash.Write([]byte(JWT))
		buf := hash.Sum(nil)
		b64 := base64.RawURLEncoding.EncodeToString(buf)
		return b64[:8]
	}()

	// Check if salt is there or generate it
	users := []User{}
	if _, err := os.Stat(".salt"); os.IsNotExist(err) && os.Getenv("UADMIN_SALT") == "" {
		Salt = GenerateBase64(72)
		ioutil.WriteFile(".salt", []byte(Salt), 0600)
		if Count(&users, "") != 0 {
			recoveryPass := GenerateBase64(24)
			recoverUsername := GenerateBase64(8)
			for Count(&users, "username = ?", recoverUsername) != 0 {
				recoverUsername = GenerateBase64(8)
			}
			admin := User{
				FirstName:    "System",
				LastName:     "Recovery Admin",
				Username:     recoverUsername,
				Password:     hashPass(recoveryPass),
				Admin:        true,
				RemoteAccess: false,
				Active:       true,
			}
			admin.Save()
			Trail(WARNING, "Your salt file was missing, and a new one was generated NO USERS CAN LOGIN UNTIL PASSWORDS ARE RESET.")
			Trail(INFO, "uAdmin generated a recovery user for you. Username:%s Password:%s", admin.Username, recoveryPass)
		}
	} else {
		Salt = os.Getenv("UADMIN_SALT")
		if Salt == "" {
			saltBytes, _ := ioutil.ReadFile(".salt")
			Salt = string(saltBytes)
		}
	}

	// Create an admin user if there is no user in the system
	adminUsername := "admin"
	adminPassword := "admin"
	if os.Getenv("UADMIN_USER") != "" {
		adminUsername = os.Getenv("UADMIN_USER")
	}
	if os.Getenv("UADMIN_PASS") != "" {
		adminPassword = os.Getenv("UADMIN_PASS")
	}
	if Count(&users, "") == 0 {
		admin := User{
			FirstName:    "System",
			LastName:     "Admin",
			Username:     adminUsername,
			Password:     hashPass(adminPassword),
			Admin:        true,
			RemoteAccess: true,
			Active:       true,
		}
		admin.Save()
		Trail(INFO, "Auto generated admin user. Username:%s, Password:%s.", adminUsername, adminPassword)
	}

	// Register admin inlines
	RegisterInlines(UserGroup{}, map[string]string{
		"GroupPermission": "UserGroupID",
	})

	RegisterInlines(User{}, map[string]string{
		"UserPermission": "UserID",
	})

	RegisterInlines(ABTest{}, map[string]string{
		"ABTestValue": "ABTestID",
	})

	for k, v := range models {
		Schema[k], _ = getSchema(v)
	}

	// Register JS
	s := Schema["abtest"]
	s.IncludeFormJS = []string{"/static/uadmin/js/abtest_form.js"}
	Schema["abtest"] = s

	// Register Limit Choices To
	s = Schema["abtest"]
	s.FieldByName("ModelName").LimitChoicesTo = loadModels
	s.FieldByName("Field").LimitChoicesTo = loadFields
	Schema["abtest"] = s

	// Load Session data
	if CacheSessions {
		loadSessions()
	}

	// Load Permission data
	if CachePermissions {
		loadPermissions()
	}

	// Check if there are active ABTests
	abTestCount = Count([]ABTest{}, "`active` = ?", true)

	// Load initial data
	err := loadInitialData()
	if err != nil {
		Trail(ERROR, "Unable to load initial data. %s", err)
	}

	// Mark registered as true to prevent auto registeration
	registered = true
}

// RegisterInlines is a function to register a model as an inline for another model
// Parameters:
// ===========
//
//	 model (struct instance): Is the model that you want to add inlines to.
//	 fk (map[interface{}]string): This is a map of the inlines to be added to the model.
//	                              The map's key is the name of the model of the inline
//	                              and the value of the map is the foreign key field's name.
//	Example:
//	========
//	type Person struct {
//	  uadmin.Model
//	  Name string
//	}
//
//	type Card struct {
//	  uadmin.Model
//	  PersonID uint
//	  Person   Person
//	}
//
//	func main() {
//	  ...
//	  uadmin.RegisterInlines(Person{}, map[string]string{
//	    "Card": "PersonID",
//	  })
//	  ...
//	}
func RegisterInlines(model interface{}, fk map[string]string) {
	// TODO: sanity check for the parameters
	// Get the name of the model
	modelName := strings.ToLower(reflect.TypeOf(model).Name())
	if inlines == nil {
		inlines = map[string][]interface{}{}
	}
	if foreignKeys == nil {
		foreignKeys = map[string]map[string]string{}
	}
	inlineList := []interface{}{}
	fkMap := map[string]string{}
	for k, v := range fk {
		kmodel, _ := NewModel(strings.ToLower(k), false)
		t := reflect.TypeOf(kmodel.Interface())
		fkMap[strings.ToLower(t.Name())] = GetDB().Config.NamingStrategy.ColumnName("", v)
		// Check if the field name is in the struct
		if t.Kind() != reflect.Struct {
			Trail(ERROR, "Unable to register inline for (%s) inline %s.%s. Please pass a struct as key.", reflect.TypeOf(model).Name(), t.Name(), v)
			continue
		}
		if _, ok := t.FieldByName(v); !ok {
			Trail(ERROR, "Unable to register inline for (%s) inline %s.%s. Field name is not in struct.", reflect.TypeOf(model).Name(), t.Name(), v)
			continue
		}
		inlineList = append(inlineList, kmodel.Interface())
	}
	inlines[modelName] = inlineList
	inlines[reflect.TypeOf(model).Name()] = inlineList
	foreignKeys[modelName] = fkMap
	delete(Schema, modelName)
	Schema[modelName], _ = getSchema(model)
}

func RegisterHandlersWithMuxer(mux *http.ServeMux, RootURL string) {
	// register static and add parameter

	if !DisableAdminUI {
		// Handler for uAdmin, static and media
		mux.HandleFunc(RootURL, Handler(mainHandler))
		if EnableDAPICORS {
			mux.HandleFunc("/static/", CORSHandler(StaticHandler))
			mux.HandleFunc("/media/", CORSHandler(mediaHandler))
		} else {
			mux.HandleFunc("/static/", Handler(StaticHandler))
			mux.HandleFunc("/media/", Handler(mediaHandler))
		}

		// api handler
		mux.HandleFunc(RootURL+"revertHandler/", Handler(revertLogHandler))
	}

	// dAPI handler
	if EnableDAPICORS {
		mux.HandleFunc(RootURL+"api/", CORSHandler(Handler(apiHandler)))
	} else {
		mux.HandleFunc(RootURL+"api/", Handler(apiHandler))
	}

	handlersRegistered = true
}

func registerHandlers(mux *http.ServeMux) {
	// register static and add parameter
	// if !strings.HasSuffix(RootURL, "/") {
	// 	RootURL = RootURL + "/"
	// }
	// if !strings.HasPrefix(RootURL, "/") {
	// 	RootURL = "/" + RootURL
	// }

	if !DisableAdminUI {
		// Handler for uAdmin, static and media
		mux.HandleFunc(RootURL, Handler(mainHandler))
		if EnableDAPICORS {
			mux.HandleFunc("/static/", CORSHandler(StaticHandler))
			mux.HandleFunc("/media/", CORSHandler(mediaHandler))
		} else {
			mux.HandleFunc("/static/", Handler(StaticHandler))
			mux.HandleFunc("/media/", Handler(mediaHandler))
		}

		// api handler
		mux.HandleFunc(RootURL+"revertHandler/", Handler(revertLogHandler))
	}

	// dAPI handler
	if EnableDAPICORS {
		mux.HandleFunc(RootURL+"api/", CORSHandler(Handler(apiHandler)))
	} else {
		mux.HandleFunc(RootURL+"api/", Handler(apiHandler))
	}

	if !DisableDAPIAuth {
		http.HandleFunc(RootURL+".well-known/openid-configuration/", Handler(JWTConfigHandler))
	}

	handlersRegistered = true
}

package uadmin

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/uadmin/uadmin/colors"
	"github.com/uadmin/uadmin/helper"
)

// HideInDashboarder used to check if a model should be hidden in
// dashboard
type HideInDashboarder interface {
	HideInDashboard() bool
}

// Register is used to register models to admin
func Register(m ...interface{}) {
	modelList := []interface{}{}

	if models == nil {
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
		}
	}

	// System models count
	SMCount := len(modelList)

	// Now add user defined models
	modelList = append(modelList,
		m...,
	)

	// Inialize the Database
	initializeDB(modelList...)

	// Setup languages
	initializeLanguage()

	// Store models in Model global variable
	// and initialize the dashboard
	dashboardMenus := []DashboardMenu{}
	All(&dashboardMenus)
	modelExists := false
	cat := ""
	Schema = map[string]ModelSchema{}
	for i := range modelList {
		t := reflect.TypeOf(modelList[i])
		name := strings.ToLower(t.Name())
		models[name] = modelList[i]

		// Register Dashboard menu
		// First check if the model is already in dashboard
		for _, val := range dashboardMenus {
			if name == val.URL {
				modelExists = true
				break
			}
		}
		// If not in dashboard, then add it
		if !modelExists {
			hideItem := false
			if _, ok := t.MethodByName("HideInDashboard"); ok {
				hider := modelList[i].(HideInDashboarder)
				hideItem = hider.HideInDashboard()
			}

			// Check if the model is a system model
			if i < SMCount {
				cat = "System"
			} else {
				cat = ""
			}
			// TODO: Make the name a plural properly
			dashboard := DashboardMenu{
				MenuName: strings.Join(helper.SplitCamelCase(t.Name()), " ") + "s",
				URL:      name,
				Hidden:   hideItem,
				Cat:      cat,
			}
			Save(&dashboard)
		}
		modelExists = false
	}

	// register static and add paramter
	if !strings.HasSuffix(RootURL, "/") {
		RootURL = RootURL + "/"
	}

	http.HandleFunc(RootURL, mainHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("./media"))))
	http.HandleFunc("/media/", mediaHandler)

	// api handler

	http.HandleFunc(RootURL+"api/", apiHandler)
	http.HandleFunc(RootURL+"revertHandler/", revertLogHandler)
	// http.HandleFunc(RootURL+"/passwordreset/", passwordResetHandler)

	//Schema = map[string]ModelSchema{}

	// Create an admin user if there is no user in the system
	users := []User{}
	if Count(&users, "") == 0 {
		admin := User{
			FirstName:    "System",
			LastName:     "Admin",
			Username:     "admin",
			Password:     hashPass("admin"),
			Admin:        true,
			RemoteAccess: true,
			Active:       true,
		}
		admin.Save()
		Trail(INFO, "Auto generated admin user. Username: admin, Password: admin.")
	}

	// Register admin inlines
	RegisterInlines(UserGroup{}, map[interface{}]string{
		GroupPermission{}: "UserGroupID",
	})

	RegisterInlines(User{}, map[interface{}]string{
		UserPermission{}: "UserID",
	})

	// Mark registered as true to prevent auto registeration
	registered = true
}

// RegisterInlines !
func RegisterInlines(model interface{}, fk map[interface{}]string) {
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
		t := reflect.TypeOf(k)
		fkMap[strings.ToLower(t.Name())] = gorm.ToColumnName(v)
		// Check if the field name is in the struct
		if t.Kind() != reflect.Struct {
			fmt.Printf("%sUnable to register inline for (%s) inline %s.%s. Please pass a struct as key.\n", colors.Error, reflect.TypeOf(model).Name(), t.Name(), v)
			continue
		}
		if _, ok := t.FieldByName(v); !ok {
			fmt.Printf("%sUnable to register inline for (%s) inline %s.%s. Field name is not in struct.\n", colors.Error, reflect.TypeOf(model).Name(), t.Name(), v)
			continue
		}
		inlineList = append(inlineList, k)
	}
	inlines[modelName] = inlineList
	foreignKeys[modelName] = fkMap

	for k, v := range models {
		t := reflect.TypeOf(v)
		Schema[t.Name()], _ = getSchema(v)
		Schema[k] = Schema[t.Name()]
	}
}

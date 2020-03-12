package uadmin

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
)

func dAPISchemaHandler(w http.ResponseWriter, r *http.Request, s *Session) {
	urlParts := strings.Split(r.URL.Path, "/")
	model, _ := NewModel(urlParts[0], false)
	modelName := gorm.ToColumnName(model.Type().Name())
	params := getURLArgs(r)

	// Check permission
	allow := false
	if disableSchemer, ok := model.Interface().(APIDisabledSchemer); ok {
		allow = disableSchemer.APIDisabledSchema(r)
		// This is a "Disable" method
		allow = !allow
		if !allow {
			ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": "Permission denied",
			})
			return
		}
	}
	if publicSchemer, ok := model.Interface().(APIPublicSchemer); ok {
		allow = publicSchemer.APIPublicSchema(r)
	}
	if !allow && s != nil {
		allow = s.User.GetAccess(modelName).Read
	}
	if !allow {
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Permission denied",
		})
		return
	}

	schema, _ := getSchema(urlParts[0])

	// Get Language
	lang := r.URL.Query().Get("language")
	if langC, err := r.Cookie("language"); err != nil || (langC != nil && langC.Value == "") {
		lang = GetDefaultLanguage().Code
	} else {
		lang = langC.Value
	}

	// Translation
	translateSchema(&schema, lang)

	if r.URL.Query().Get("$choices") == "1" {
		// Load Choices for FK
		for i := range schema.Fields {
			if schema.Fields[i].Type == cFK || schema.Fields[i].Type == cM2M {
				choices := getChoices(schema.Fields[i].TypeName)
				schema.Fields[i].Choices = choices
			}
		}
	}

	returnDAPIJSON(w, r, map[string]interface{}{
		"status": "ok",
		"result": schema,
	}, params, "schema", model)

	go func() {
		// Check if log is required
		log := APILogSchema
		if logSchemer, ok := model.Interface().(APILogSchemer); ok {
			log = logSchemer.APILogSchema(r)
		}

		if log {
			user := ""
			if s != nil {
				user = s.User.Username
			}
			activity, _ := json.Marshal(map[string]interface{}{
				"_IP": r.RemoteAddr,
			})
			log := Log{
				Username:  user,
				Action:    Action(0).GetSchema(),
				TableName: modelName,
				Activity:  string(activity),
			}
			log.Save()
		}
	}()
}

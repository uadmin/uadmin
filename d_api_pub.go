package uadmin

import (
	"context"
	"net/http"
	"strconv"
	"strings"
)

func DAPIHandlerPublic(w http.ResponseWriter, r *http.Request, modelKV DApiModelKeyVal) {

	var s *Session

	if modelKV.session != nil {
		s = modelKV.session
	} else {
		s = IsAuthenticated(r)
		if s != nil {
			s.ThroughAPI = true
		}

		modelKV.session = s
	}

	// Parse the Form
	err := r.ParseMultipartForm(2 << 10)
	if err != nil {
		r.ParseForm()
	}

	// Add Custom headers
	for k, v := range CustomDAPIHeaders {
		w.Header().Add(k, v)
	}

	ctx := context.WithValue(r.Context(), CKey("dAPI"), true)
	ctx = context.WithValue(r.Context(), CKey("modelName"), modelKV)

	r = r.WithContext(ctx)

	// auth dAPI
	if modelKV.PathCommandName == Auth.String() {
		dAPIAuthHandler(w, r, s)
		return
	}

	if modelKV.PathCommandName == AllModels.String() {
		if s == nil || !s.User.Admin {
			w.WriteHeader(http.StatusForbidden)
			ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": "access denied",
			})
			return
		}
		dAPIAllModelsHandler(w, r, s)
		return
	}

	// Check if there is no command and show help
	if modelKV.PathCommandName == "" || modelKV.PathCommandName == Help.String() {
		if s == nil {
			w.WriteHeader(http.StatusForbidden)
			ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": "access denied",
			})
			return
		}

		w.Write([]byte(dAPIHelp))
		return
	}

	// sanity check
	// check model name
	modelExists := false
	var model interface{}
	for k, v := range models {
		if modelKV.PathCommandName == k {
			modelExists = true
			model = v
			break
		}
	}
	if !modelExists {
		w.WriteHeader(404)
		ReturnJSON(w, r, map[string]string{
			"status":  "error",
			"err_msg": "Model name not found (" + modelKV.PathCommandName + ")",
		})
		return
	}

	//check command
	commandExists := false
	command := ""
	secondPartIsANumber := false
	if _, err := strconv.Atoi(modelKV.CommandName); err == nil {
		secondPartIsANumber = true
	}
	if modelKV.CommandName != "" && !secondPartIsANumber {
		for _, i := range DataCommands {
			if modelKV.CommandName == i {
				commandExists = true
				command = i

				// trim command from URL
				r.URL.Path = strings.TrimPrefix(r.URL.Path, modelKV.CommandName)
				r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")

				break
			}
		}
	} else {
		commandExists = true
		switch r.Method {
		case http.MethodGet:
			command = "read"
		case http.MethodPost:
			command = "add"
		case http.MethodPut:
			command = "edit"
		case http.MethodDelete:
			command = "delete"
		}
	}

	if !commandExists {
		w.WriteHeader(404)
		ReturnJSON(w, r, map[string]string{
			"status":  "error",
			"err_msg": "Invalid command (" + command + ")",
		})
		return
	}

	// Route the request to the correct handler based on the command
	if command == "read" {
		// check if there is a prequery
		if APIPreQueryReadHandler != nil && !APIPreQueryReadHandler(w, r) {
			return
		}
		if preQuery, ok := model.(APIPreQueryReader); ok && !preQuery.APIPreQueryRead(w, r) {
		} else {
			dAPIReadHandler(w, r, s)
		}
		return
	}
	if command == "add" {
		// check if there is a prequery
		if APIPreQueryAddHandler != nil && !APIPreQueryAddHandler(w, r) {
			return
		}
		if preQuery, ok := model.(APIPreQueryAdder); ok && !preQuery.APIPreQueryAdd(w, r) {
		} else {
			dAPIAddHandler(w, r, s)
		}
		return
	}
	if command == "edit" {
		// check if there is a prequery
		if APIPreQueryEditHandler != nil && !APIPreQueryEditHandler(w, r) {
			return
		}
		if preQuery, ok := model.(APIPreQueryEditor); ok && !preQuery.APIPreQueryEdit(w, r) {
		} else {
			dAPIEditHandler(w, r, s)
		}
		return
	}
	if command == "delete" {
		// check if there is a prequery
		if APIPreQueryDeleteHandler != nil && !APIPreQueryDeleteHandler(w, r) {
			return
		}
		if preQuery, ok := model.(APIPreQueryDeleter); ok && !preQuery.APIPreQueryDelete(w, r) {
		} else {
			dAPIDeleteHandler(w, r, s)
		}
		return
	}
	if command == "schema" {
		// check if there is a prequery
		if preQuery, ok := model.(APIPreQuerySchemer); ok && !preQuery.APIPreQuerySchema(w, r) {
		} else {
			dAPISchemaHandler(w, r, s)
		}
		return
	}
	if command == "method" {
		dAPIMethodHandler(w, r, s)
		if r.URL.Query().Get("$next") != "" {
			if strings.HasPrefix(r.URL.Query().Get("$next"), "$back") && r.Header.Get("Referer") != "" {
				http.Redirect(w, r, r.Header.Get("Referer")+strings.TrimPrefix(r.URL.Query().Get("$next"), "$back"), http.StatusSeeOther)
			} else {
				http.Redirect(w, r, r.URL.Query().Get("$next"), http.StatusSeeOther)
			}
		}
	}
}

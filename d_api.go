package uadmin

import (
	"context"
	"net/http"
	"strconv"
	"strings"
)

const dAPIHelp = `
Data Access API Help (dAPI)
===========================

Command:
========
URL                              Command
======================================================================================================
/modelname/read/                 Read Multiple
/modelname/read/1/               Read One
/modelname/add/?f__0=0&f__1=1    Add Multiple
/modelname/add/                  Add One
/modelname/edit/?f=1&_f=0        Edit Multiple
/modelname/edit/1/               Edit One
/modelname/delete/?f=1           Delete Multiple
/modelname/delete/1/             Delete One
/modelname/method/METHOD_NAME/1/ Run method on model where id=1
/modelname/schema/               Schema
/$allmodels/                     All Models
/auth/                     	     Authorization methods
/help/                     	     Help


Field Filtering:
================
Filter                 Description
======================================================================================================
__gt                   Greater Than
__gte                  Greater Than or Equal To
__lt                   Less Than
__lte                  Less Than or Equal To
__in                   Find a value matching any of these values 
__is                   Stands for IS NULL
__contains             Search for string values that contains
__between              Selects values within a given range
__startswith           Search for string values that starts with a given substring
__endswith             Search for string values that ends with a given substring
__re                   Regex
__icontains            Similar to __contains except it is case insensitive
__istartswith          Similar to __startswith except it is case insensitive
__iendswith            Similar to __endswith except it is case insensitive
!{FIELD}__{OP}         Negates the operator e.g. !id__in means: id NOT IN (?)
$or=f=0|f=1            OR operator (f=0 OR f=1)
$or=f1=0+f2=0|f=1      OR operator with a nested AND ((f1=0 AND f2=0) OR f=1)


URL Symbols:
============
Symbol                 Description                               Example
=======================================================================================================
-                      Descending Order                          $order=-fieldname
_                      Writing Data (Add/Edit)                   $_f=f1


Special Parameters:
===================
Query                  Description
======================================================================================================
$limit=1               Number of records that you want to return
$offset=1              Starting point to read in the list of records
$order=f1,-f           Used to sort the results. Use "-" for descending order and comma for
                       more field
$f=f1,f2               Selecting Fields
$groupby=f             Groups rows that have the same values into summary rows
$deleted=1             Returns results including deleted records
$join=[inner__]m[__f]  Joins results from another model based on a foreign key
$m2m={0,1,fill,id}     Returns results from M2M fields where:
                         0       : Don't return
                         [1,fill]: Return all fields
                         id      : Only return IDs
$m2m=f__{id,fill}      Returns results from a specific M2M field
$q=abc                 Searches all string based fields for read, edit, and delete requests
$preload={0,1}         Fills the data from foreign keys into structs
                         0 : Don't return
                         1 : Return preloaded data
$next=/                Used in read method that redirects the user to the specified path 
                       after processing the request
                         $back: Send the user back
$stat=1                Returns the query execution time in milliseconds


Aggregation Operators:
======================
Operator              Description
======================================================================================================
__sum                 Used in $f that returns the total sum of a numeric field
__avg                 Used in $f that returns the average value of a numeric field 
__min                 Used in $f that returns the smallest value of the selected column
__max                 Used in $f that returns the largest value of the selected column
__count               Used in $f that returns the number of rows that matches a specified criteria
__date                Used in $f that returns DATE() of the field
__year                Used in $f that returns YEAR() of the field
__month               Used in $f that returns MONTH() of the field
__day                 Used in $f that returns DAY() of the field

For full documentation: https://uadmin-docs.readthedocs.io/en
/latest/dapi.html
`

// CKey is the standard key used in uAdmin for context keys
type CKey string

func dAPIHandler(w http.ResponseWriter, r *http.Request, s *Session) {
	// Parse the Form
	err := r.ParseMultipartForm(2 << 10)
	if err != nil {
		r.ParseForm()
	}

	// Add Custom headers
	for k, v := range CustomDAPIHeaders {
		w.Header().Add(k, v)
	}

	// Removes the dApi from the path and just leaves the command path
	//  http://route.com/api/d/modelname/add/?f__0=0&f__1=1 -> modelname/add/?f__0=0&f__1=1
	r.URL.Path = strings.TrimPrefix(r.URL.Path, RootURL+"api/d")
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

	// Splits the command path for handling
	// modelname/add/?f__0=0&f__1=1 -> [ modelname , add, ?f__0=0&f__1=1]
	urlParts := strings.Split(r.URL.Path, "/")

	pathCommand := urlParts[0]
	commandName := ""
	dataCommand := ""
	dataCommandExtra := ""

	if len(urlParts) > 1 {
		commandName = urlParts[1]
		commandName = strings.TrimPrefix(commandName, "/")

		r.URL.Path = strings.TrimPrefix(r.URL.Path, commandName)
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")
	}

	if len(urlParts) > 2 {
		dataCommand = urlParts[2]
		dataCommand = strings.TrimPrefix(dataCommand, "/")
	}

	if len(urlParts) > 3 {
		dataCommandExtra = urlParts[3]
	}

	v, _ := ParseCommandString(pathCommand)

	modelKV := DApiModelKeyVal{
		PathCommand:     v,
		PathCommandName: pathCommand,
		CommandName:     commandName,
		DataCommand:     dataCommand,
		DataForMethod:   dataCommandExtra,
	}
	ctx := context.WithValue(r.Context(), CKey("dAPI"), true)
	ctx = context.WithValue(r.Context(), CKey("modelName"), modelKV)
	r = r.WithContext(ctx)

	// auth dAPI
	if pathCommand == Auth.String() {
		dAPIAuthHandler(w, r, s)
		return
	}

	if pathCommand == AllModels.String() {
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
	if pathCommand == "" || pathCommand == Help.String() {
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
		if pathCommand == k {
			modelExists = true
			model = v

			//TODO: remove path
			// add model to context
			var dApiModel DApiModelKeyVal
			dApiModel.CommandName = pathCommand

			//get id if exists
			//TODO: validate that the id is really an id and not something else
			dApiModel.DataCommand = dataCommand
			dApiModel.DataForMethod = dataCommandExtra

			ctx := context.WithValue(r.Context(), CKey("modelName"), dApiModel)
			r = r.WithContext(ctx)

			//TODO: verify where the command is read
			// trim model name from URL
			r.URL.Path = strings.TrimPrefix(r.URL.Path, pathCommand)
			r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")

			break
		}
	}
	if !modelExists {
		w.WriteHeader(404)
		ReturnJSON(w, r, map[string]string{
			"status":  "error",
			"err_msg": "Model name not found (" + pathCommand + ")",
		})
		return
	}

	//check command
	commandExists := false
	command := ""
	secondPartIsANumber := false
	if len(urlParts) > 1 {
		if _, err := strconv.Atoi(dataCommand); err == nil {
			secondPartIsANumber = true
		}
	}
	if len(urlParts) > 1 && !secondPartIsANumber {
		for _, i := range DataCommands {
			if dataCommand == i {
				commandExists = true
				command = i

				// trim command from URL
				r.URL.Path = strings.TrimPrefix(r.URL.Path, dataCommand)
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

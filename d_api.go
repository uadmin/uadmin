package uadmin

import (
	//"fmt"
	"context"
	"net/http"
	"strings"
	"time"
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
/modelname/schema/               Schema
/$allmodels/                     All Models


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
__contains             Search for string values that contract
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

	// Timestamp
	ctx := context.WithValue(r.Context(), CKey("start"), time.Now())
	r = r.WithContext(ctx)

	r.URL.Path = strings.TrimPrefix(r.URL.Path, RootURL+"api/d")
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")

	urlParts := strings.Split(r.URL.Path, "/")

	// Check if there is no command and show help
	if r.URL.Path == "" || r.URL.Path == "/" || len(urlParts) < 2 {
		if urlParts[0] == "$allmodels" {
			dAPIAllModelsHandler(w, r, s)
			return
		}
		w.Write([]byte(dAPIHelp))
		return
	}

	// sanity check
	// check model name
	modelExists := false
	for i := range models {
		if urlParts[0] == i {
			modelExists = true
			break
		}
	}
	if !modelExists {
		w.WriteHeader(404)
		ReturnJSON(w, r, map[string]string{
			"status":  "error",
			"err_msg": "Model name not found (" + urlParts[0] + ")",
		})
		return
	}
	//check command
	commandExists := false
	for _, i := range []string{"read", "add", "edit", "delete", "schema"} {
		if urlParts[1] == i {
			commandExists = true
			break
		}
	}
	if !commandExists {
		w.WriteHeader(404)
		ReturnJSON(w, r, map[string]string{
			"status":  "error",
			"err_msg": "Invalid command (" + urlParts[1] + ")",
		})
		return
	}

	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

	// Route the request to the correct handler based on the command
	if urlParts[1] == "read" {
		dAPIReadHandler(w, r, s)
		return
	}
	if urlParts[1] == "add" {
		dAPIAddHandler(w, r, s)
		return
	}
	if urlParts[1] == "edit" {
		dAPIEditHandler(w, r, s)
		return
	}
	if urlParts[1] == "delete" {
		dAPIDeleteHandler(w, r, s)
		return
	}
	if urlParts[1] == "schema" {
		dAPISchemaHandler(w, r, s)
		return
	}
}
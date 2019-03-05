package main

import (
	"net/http"

	"github.com/rn1hd/todo/api"
	"github.com/rn1hd/todo/models"
	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.RootURL = "/admin/"
	uadmin.Register(
		models.Todo{},
		models.Category{},
		models.Friend{},
		models.Item{},
	)

	uadmin.RegisterInlines(models.Category{}, map[string]string{
		"Todo": "CategoryID",
	})
	uadmin.RegisterInlines(models.Friend{}, map[string]string{
		"Todo": "FriendID",
	})
	uadmin.RegisterInlines(models.Item{}, map[string]string{
		"Todo": "ItemID",
	})

	// API Handler
	http.HandleFunc("/api/", api.APIHandler)

	uadmin.Port = 8000
	uadmin.StartServer()
}

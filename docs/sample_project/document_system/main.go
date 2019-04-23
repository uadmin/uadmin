package main

import (
	"github.com/uadmin/uadmin"
	"github.com/uadmin/uadmin/docs/sample_project/document_system/models"
)

func main() {
	// Register models to uAdmin
	uadmin.Register(
		models.Folder{},
		models.FolderGroup{},
		models.FolderUser{},
		models.Channel{},
		models.Document{},
		models.DocumentGroup{},
		models.DocumentUser{},
		models.DocumentVersion{},
	)

	// Register FolderGroup and FolderUser to Folder model
	uadmin.RegisterInlines(
		models.Folder{},
		map[string]string{
			"foldergroup": "FolderID",
			"folderuser":  "FolderID",
		},
	)

	// Register DocumentVersion, DocumentGroup, and DocumentUser to Document model
	uadmin.RegisterInlines(
		models.Document{},
		map[string]string{
			"documentversion": "DocumentID",
			"documentgroup":   "DocumentID",
			"documentuser":    "DocumentID",
		},
	)

	// Initialize docS variable that calls the document model in the schema
	docS := uadmin.Schema["document"]

	// Assigns CreatedByFormFilter to the FormModifier
	docS.FormModifier = models.CreatedByFormFilter

	// Assigns DocumentListFilter to the ListModifier
	docS.ListModifier = DocumentListFilter

	// Pass the docS back to the schema of document model
	uadmin.Schema["document"] = docS

	// Sets the name of the website that shows on title and dashboard
	uadmin.SiteName = "Document System"

	// Sets a loopback IP address
	uadmin.BindIP = "127.0.0.1"

	// Activates a uAdmin server
	uadmin.StartServer()
}

// DocumentListFilter !
func DocumentListFilter(s *uadmin.ModelSchema, u *uadmin.User) (string, []interface{}) {
	// Checks whether the user is not an admin
	if !u.Admin {
		// Returns the user ID
		return "user_id = ?", []interface{}{u.ID}
	}
	// Returns nothing
	return "", []interface{}{}
}

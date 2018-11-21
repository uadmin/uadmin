package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

// Todo model ...
type Todo struct {
	uadmin.Model
	Name        string
	Description string `uadmin:"html"`
	Category    Category
	CategoryID  uint
	Friend      Friend `uadmin:"help:Who will be a part of your activity?"`
	FriendID    uint
	Item        Item `uadmin:"help:What are the requirements needed in order to accomplish your activity?"`
	ItemID      uint
	TargetDate  time.Time
	Progress    int `uadmin:"progress_bar"`
}

// Preload ...
func (t *Todo) Preload() {
	if t.Category.ID != t.CategoryID {
		uadmin.Get(&t.Category, "id = ?", t.CategoryID)
	}
	if t.Friend.ID != t.FriendID {
		uadmin.Get(&t.Friend, "id = ?", t.FriendID)
	}
	if t.Item.ID != t.ItemID {
		uadmin.Get(&t.Item, "id = ?", t.ItemID)
	}
}

// Validate function ...
func (t Todo) Validate() (errMsg map[string]string) {
	// Initialize the error messages
	errMsg = map[string]string{}
	// Get any records from the database that maches the name of
	// this record and make sure the record is not the record we are
	// editing right now
	todo := Todo{}
	if uadmin.Count(&todo, "name = ? AND id <> ?", t.Name, t.ID) != 0 {
		errMsg["Name"] = "This todo name is already in the system"
	}
	return
}

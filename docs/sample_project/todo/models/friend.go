package models

import "github.com/uadmin/uadmin"

// Nationality ...
type Nationality int

// Chinese ...
func (n Nationality) Chinese() Nationality {
	return 1
}

// Filipino ...
func (n Nationality) Filipino() Nationality {
	return 2
}

// Others ...
func (n Nationality) Others() Nationality {
	return 3
}

// Friend model ...
type Friend struct {
	uadmin.Model
	Name        string `uadmin:"required"`
	Email       string `uadmin:"email"`
	Password    string `uadmin:"password;list_exclude"`
	Nationality Nationality
	Invite      string `uadmin:"link"`
}

// Save !
func (f *Friend) Save() {
	f.Invite = "https://uadmin.io/"
	uadmin.Save(f)
}

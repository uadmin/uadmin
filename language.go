package uadmin

import "fmt"

// Language !
type Language struct {
	Model
	EnglishName    string `uadmin:"required;read_only;filter;search"`
	Name           string `uadmin:"required;read_only;filter;search"`
	Flag           string `uadmin:"image;list_exclude"`
	Code           string `uadmin:"filter;read_only;list_exclude"`
	RTL            bool   `uadmin:"list_exclude"`
	Default        bool   `uadmin:"help:Set as the default language;list_exclude"`
	Active         bool   `uadmin:"help:To show this in available languages;filter"`
	AvailableInGui bool   `uadmin:"help:The App is available in this language;read_only"`
}

// String !
func (l Language) String() string {
	return l.Code
}

// Save !
func (l *Language) Save() {
	if l.Default {
		Update([]Language{}, "`default`", false, "`default` = ?", true)
		defaultLang = *l
	}
	Save(l)
	tempActiveLangs := []Language{}
	Filter(&tempActiveLangs, "`active` = ?", true)
	activeLangs = tempActiveLangs

	tanslationList := []translation{}
	for i := range activeLangs {
		tanslationList = append(tanslationList, translation{
			Active:  activeLangs[i].Active,
			Default: activeLangs[i].Default,
			Code:    activeLangs[i].Code,
			Name:    fmt.Sprintf("%s (%s)", activeLangs[i].Name, activeLangs[i].EnglishName),
		})
	}

	for modelName := range Schema {
		for i := range Schema[modelName].Fields {
			if Schema[modelName].Fields[i].Type == cMULTILINGUAL {
				Schema[modelName].Fields[i].Translations = tanslationList
			}
		}
	}
}

func GetDefaultLanguage() Language {
	return defaultLang
}

func GetActiveLanguages() []Language {
	return activeLangs
}

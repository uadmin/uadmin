package uadmin

// Language !
type Language struct {
	Model
	EnglishName    string `uadmin:"required;read_only;filter;search"`
	Name           string `uadmin:"required;read_only;filter;search"`
	Flag           string `uadmin:"type:image"`
	Code           string `uadmin:"filter;read_only;list_exclude"`
	Direction      string
	Default        bool `uadmin:"help:Set as the default language"`
	Active         bool `uadmin:"help:To show this in available languages;filter"`
	AvailableInGui bool `uadmin:"help:The App is available in this language;read_only"`
}

// String !
func (l Language) String() string {
	return l.Code
}

// Save !
func (l *Language) Save() {
	if l.Default {
		Update(l, "default", false, "default = ?", true)
		defaultLang = *l
	}
	Save(l)
	tempActiveLangs := []Language{}
	Filter(&activeLangs, "active = ?", true)
	activeLangs = tempActiveLangs
}

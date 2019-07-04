package uadmin

import (
	"github.com/uadmin/uadmin/helper"
	"strings"
)

type ReadOnly int

func (ReadOnly) Always() ReadOnly {
	return 1
}

func (ReadOnly) New() ReadOnly {
	return 2
}

func (ReadOnly) Edit() ReadOnly {
	return 3
}

func (r *ReadOnly) GetTag() string {
	if *r == r.Always() {
		return "read_only"
	}
	if *r == r.Edit() {
		return "read_only:true,edit"
	}
	if *r == r.New() {
		return "read_only:true,new"
	}
	return ""
}

type BuilderField struct {
	Model
	Builder           Builder `uadmin:"required;list_exclude"`
	BuilderID         uint
	Name              string    `uadmin:"required;list_exclude"`
	CodeName          string    `uadmin:"read_only"`
	Type              FieldType `uadmin:"required"`
	Required          bool
	Help              string `uadmin:"list_exclude"`
	Search            bool   `uadmin:"list_exclude"`
	ReadOnly          ReadOnly
	Max               string `uadmin:"list_exclude"`
	Min               string `uadmin:"list_exclude"`
	Hidden            bool
	ListExclude       bool `uadmin:"list_exclude"`
	Filter            bool
	CategoricalFilter bool   `uadmin:"list_exclude"`
	Encrypt           bool   `uadmin:"list_exclude"`
	DefaultValue      string `uadmin:"list_exclude"`
	Pattern           string `uadmin:"list_exclude"`
	PatternMsg        string `uadmin:"list_exclude"`
	DisplayName       string `uadmin:"list_exclude"`
}

func (b *BuilderField) GetCode() (code string, imports string) {
	// Type map is a list field type to basic type
	typeMap := map[FieldType]string{
		b.Type.Boolean():      "bool",
		b.Type.Code():         "string",
		b.Type.DateTime():     "time.Time",
		b.Type.DateTimePtr():  "*time.Time",
		b.Type.Email():        "string",
		b.Type.File():         "string",
		b.Type.Float():        "float64",
		b.Type.ForeignKey():   "{FKEY}",
		b.Type.HTML():         "string",
		b.Type.Image():        "string",
		b.Type.Int():          "int",
		b.Type.Link():         "string",
		b.Type.M2M():          "{M2M}",
		b.Type.Money():        "float64",
		b.Type.Multilingual(): "string",
		b.Type.Password():     "string",
		b.Type.ProgressBar():  "float64",
		b.Type.StaticList():   "{LIST}",
		b.Type.String():       "string",
	}

	// required imports for field types
	importMap := map[FieldType]string{
		b.Type.DateTime():    "time",
		b.Type.DateTimePtr(): "time",
	}

	// required tags for field types
	tagMap := map[FieldType]string{
		b.Type.Code():         "code",
		b.Type.Email():        "email",
		b.Type.File():         "file",
		b.Type.HTML():         "html",
		b.Type.Image():        "image",
		b.Type.Link():         "link",
		b.Type.Money():        "money",
		b.Type.Multilingual(): "multilingual",
		b.Type.Password():     "password",
		b.Type.ProgressBar():  "progress_bar",
	}

	fName := b.CodeName
	fType := typeMap[b.Type]

	// check if there are any required imports
	if val, ok := importMap[b.Type]; ok {
		imports = val
	}

	// check if there are any required type tags
	fTags := []string{}
	if val, ok := tagMap[b.Type]; ok {
		fTags = append(fTags, val)
	}

	// check for meta tags
	if b.Required {
		fTags = append(fTags, "required")
	}

	if b.Help != "" {
		fTags = append(fTags, "help:"+b.Help)
	}

	if b.Search {
		fTags = append(fTags, "search")
	}

	if b.ReadOnly != ReadOnly(0) {
		fTags = append(fTags, b.ReadOnly.GetTag())
	}

	if b.Max != "" {
		fTags = append(fTags, "max:"+b.Max)
	}

	if b.Min != "" {
		fTags = append(fTags, "min:"+b.Min)
	}

	if b.Hidden {
		fTags = append(fTags, "hidden")
	}

	if b.ListExclude {
		fTags = append(fTags, "list_exclude")
	}

	if b.Filter {
		fTags = append(fTags, "filter")
	}

	if b.CategoricalFilter {
		fTags = append(fTags, "categorical_filter")
	}

	if b.Encrypt {
		fTags = append(fTags, "encrypt")
	}

	if b.DefaultValue != "" {
		fTags = append(fTags, "default_value:"+b.DefaultValue)
	}

	if b.Pattern != "" {
		fTags = append(fTags, "pattern:"+b.Pattern)
	}

	if b.PatternMsg != "" {
		fTags = append(fTags, "pattern_msg:"+b.PatternMsg)
	}

	if b.DisplayName != "" {
		fTags = append(fTags, "display_name:"+b.DisplayName)
	}

	fTag := ""
	if len(fTags) > 0 {
		fTag += " `uadmin:\""
		fTag += strings.Join(fTags, ";")
		fTag += "\"`"
	}

	return fName + " " + fType + fTag, imports
}

func (b *BuilderField) Save() {
	b.CodeName = helper.ToCamel(b.Name)
	Save(b)
}

func (BuilderField) HideInDashboard() bool {
	return true
}

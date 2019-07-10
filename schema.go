package uadmin

import (
	"strings"
	"time"
)

// ModelSchema for a form
type ModelSchema struct {
	Name          string // Name of the Model
	DisplayName   string // Display Name of the model
	ModelName     string // URL
	ModelID       uint
	Inlines       []*ModelSchema
	InlinesData   []listData
	Fields        []F
	IncludeFormJS []string
	IncludeListJS []string
	FormModifier  func(*ModelSchema, interface{}, *User)
	ListModifier  func(*ModelSchema, *User) (string, []interface{})
	FormTheme     string
	ListTheme     string
}

// FieldByName returns a field from a ModelSchema by name or nil if
// it doen't exist
func (s ModelSchema) FieldByName(a string) *F {
	for i := range s.Fields {
		if strings.ToLower(s.Fields[i].Name) == strings.ToLower(a) {
			return &s.Fields[i]
		}
	}
	return &F{}
}

// GetFormTheme returns the theme for this model or the
// global theme if there is no assigned theme for the model
func (s *ModelSchema) GetFormTheme() string {
	if s.FormTheme == "" {
		return Theme
	}
	return s.FormTheme
}

// GetListTheme returns the theme for this model or the
// global theme if there is no assigned theme for the model
func (s *ModelSchema) GetListTheme() string {
	if s.ListTheme == "" {
		return Theme
	}
	return s.ListTheme
}

// ApprovalAction is a selection of approval actions
type ApprovalAction int

// Approved is an accepted change
func (ApprovalAction) Approved() ApprovalAction {
	return 1
}

// Rejected is a rejected change
func (ApprovalAction) Rejected() ApprovalAction {
	return 2
}

// F is a field
type F struct {
	Name              string
	DisplayName       string
	Type              string
	TypeName          string
	Value             interface{}
	Help              string
	Max               interface{}
	Min               interface{}
	Format            string
	DefaultValue      string
	Required          bool
	Pattern           string
	PatternMsg        string
	Hidden            bool
	ReadOnly          string
	Searchable        bool
	Filter            bool
	ListDisplay       bool
	FormDisplay       bool
	CategoricalFilter bool
	Translations      []translation
	Choices           []Choice
	IsMethod          bool
	ErrMsg            string
	ProgressBar       map[float64]string
	LimitChoicesTo    func(interface{}, *User) []Choice
	UploadTo          string
	Encrypt           bool
	Approval          bool
	NewValue          interface{}
	OldValue          interface{}
	ChangedBy         string
	ChangeDate        *time.Time
	ApprovalAction    ApprovalAction
	ApprovalDate      *time.Time
	ApprovalBy        string
	ApprovalID        uint
	WebCam            bool
	Stringer          bool
}

// Choice is a struct for list choices
type Choice struct {
	V        string
	K        uint
	Selected bool
}

package uadmin

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

// ModelSchema for a form
type ModelSchema struct {
	Name          string // Name of the Model e.g. OrderItem
	DisplayName   string // Display Name of the model e.g. Order Items
	ModelName     string // URL e.g. orderitem
	TableName     string // DB table name e.g. order_items
	ModelID       uint
	Inlines       []*ModelSchema
	InlinesData   []listData
	Fields        []F
	IncludeFormJS []string
	IncludeListJS []string
	FormModifier  func(*ModelSchema, interface{}, *User)            `json:"-"`
	ListModifier  func(*ModelSchema, *User) (string, []interface{}) `json:"-"`
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

func (m ModelSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name          string // Name of the Model
		DisplayName   string // Display Name of the model
		ModelName     string // URL
		TableName     string
		ModelID       uint
		Inlines       []*ModelSchema
		InlinesData   []listData
		Fields        []F
		IncludeFormJS []string
		IncludeListJS []string
		FormModifier  *string
		ListModifier  *string
		FormTheme     string
		ListTheme     string
	}{
		Name:          m.Name,
		DisplayName:   m.DisplayName,
		ModelName:     m.ModelName,
		TableName:     m.TableName,
		ModelID:       m.ModelID,
		Inlines:       m.Inlines,
		InlinesData:   m.InlinesData,
		Fields:        m.Fields,
		IncludeFormJS: m.IncludeFormJS,
		IncludeListJS: m.IncludeListJS,
		FormModifier: func() *string {
			v := ""
			if m.FormModifier == nil {
				return nil
			} else {
				v = runtime.FuncForPC(reflect.ValueOf(m.FormModifier).Pointer()).Name()
				return &v
			}
		}(),
		ListModifier: func() *string {
			v := ""
			if m.ListModifier == nil {
				return nil
			} else {
				v = runtime.FuncForPC(reflect.ValueOf(m.ListModifier).Pointer()).Name()
				return &v
			}
		}(),
		FormTheme: m.FormTheme,
		ListTheme: m.ListTheme,
	})
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
	ColumnName        string
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
	ProgressBar       map[float64]string                `json:"-"`
	LimitChoicesTo    func(interface{}, *User) []Choice `json:"-"`
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

func (f F) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name              string
		DisplayName       string
		ColumnName        string
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
		ProgressBar       map[string]string
		LimitChoicesTo    *string
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
	}{
		Name:              f.Name,
		DisplayName:       f.DisplayName,
		ColumnName:        f.ColumnName,
		Type:              f.Type,
		TypeName:          f.TypeName,
		Value:             f.Value,
		Help:              f.Help,
		Max:               f.Max,
		Min:               f.Min,
		Format:            f.Format,
		DefaultValue:      f.DefaultValue,
		Required:          f.Required,
		Pattern:           f.Pattern,
		PatternMsg:        f.PatternMsg,
		Hidden:            f.Hidden,
		ReadOnly:          f.ReadOnly,
		Searchable:        f.Searchable,
		Filter:            f.Filter,
		ListDisplay:       f.ListDisplay,
		FormDisplay:       f.FormDisplay,
		CategoricalFilter: f.CategoricalFilter,
		Translations:      f.Translations,
		Choices:           f.Choices,
		IsMethod:          f.IsMethod,
		ErrMsg:            f.ErrMsg,
		ProgressBar: func() map[string]string {
			tempMap := map[string]string{}
			for k, v := range f.ProgressBar {
				tempMap[fmt.Sprint(k)] = v
			}
			return tempMap
		}(),
		LimitChoicesTo: func() *string {
			v := ""

			if f.LimitChoicesTo == nil {
				return nil
			} else {
				v = runtime.FuncForPC(reflect.ValueOf(f.LimitChoicesTo).Pointer()).Name()
				return &v
			}
		}(),
		UploadTo:       f.UploadTo,
		Encrypt:        f.Encrypt,
		Approval:       f.Approval,
		NewValue:       f.NewValue,
		OldValue:       f.OldValue,
		ChangedBy:      f.ChangedBy,
		ChangeDate:     f.ChangeDate,
		ApprovalAction: f.ApprovalAction,
		ApprovalDate:   f.ApprovalDate,
		ApprovalBy:     f.ApprovalBy,
		ApprovalID:     f.ApprovalID,
		WebCam:         f.WebCam,
		Stringer:       f.Stringer,
	})
}

// Choice is a struct for list choices
type Choice struct {
	V        string
	K        uint
	Selected bool
}

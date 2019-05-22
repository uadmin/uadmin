package uadmin

import (
	"strconv"
	"strings"
	"time"
)

type DataType int

func (DataType) String() DataType {
	return 1
}

func (DataType) Integer() DataType {
	return 2
}

func (DataType) Float() DataType {
	return 3
}

func (DataType) Boolean() DataType {
	return 4
}

func (DataType) File() DataType {
	return 5
}

func (DataType) Image() DataType {
	return 6
}

func (DataType) DateTime() DataType {
	return 7
}

type Setting struct {
	Model
	Name         string
	DefaultValue string
	DataType     DataType
	Value        string
	Help         string
	Category     SettingCategory `uadmin:"required"`
	CategoryID   uint
	Code         string `uadmin:"read_only"`
}

func (s *Setting) Save() {
	Preload(s)
	s.Code = strings.ReplaceAll(s.Category.Name, " ", "") + "." + strings.ReplaceAll(s.Name, " ", "")
	Save(s)
}

func (Setting) HideInDashboarder() bool {
	return true
}

func (s *Setting) ParseFormValue(v []string) {
	switch s.DataType {
	case s.DataType.Boolean():
		tempV := len(v) == 1 && v[0] == "on"
		if tempV {
			s.Value = "1"
		} else {
			s.Value = "0"
		}
	case s.DataType.DateTime():
		if len(v) == 1 && v[0] != "" {
			s.Value = v[0] + ":00"
		} else {
			s.Value = ""
		}
	default:
		if len(v) == 1 && v[0] != "" {
			s.Value = v[0]
		} else {
			s.Value = ""
		}
	}
}

func GetSetting(code string) interface{} {
	var err error
	s := Setting{}
	Get(&s, "code = ?", code)

	if s.ID == 0 {
		return nil
	}

	var v interface{}

	switch s.DataType {
	case s.DataType.String():
		if s.Value == "" {
			v = s.DefaultValue
		} else {
			v = s.Value
		}
	case s.DataType.Integer():
		if s.Value != "" {
			v, err = strconv.ParseInt(s.Value, 10, 64)
		}
		if err != nil {
			v, err = strconv.ParseInt(s.DefaultValue, 10, 64)
		}
		if err != nil {
			v = 0
		}
	case s.DataType.Float():
		if s.Value != "" {
			v, err = strconv.ParseFloat(s.Value, 64)
		}
		if err != nil {
			v, err = strconv.ParseFloat(s.DefaultValue, 64)
		}
		if err != nil {
			v = 0.0
		}
	case s.DataType.Boolean():
		if s.Value != "" {
			v = s.Value == "1"
		}
		if v == nil {
			v = s.DefaultValue == "1"
		}
	case s.DataType.File():
		if s.Value == "" {
			v = s.DefaultValue
		} else {
			v = s.Value
		}
	case s.DataType.Image():
		if s.Value == "" {
			v = s.DefaultValue
		} else {
			v = s.Value
		}
	case s.DataType.DateTime():
		if s.Value != "" {
			v, err = time.Parse("2006-01-02 15:04:05", s.Value)
		}
		if err != nil {
			v, err = time.Parse("2006-01-02 15:04:05", s.DefaultValue)
		}
		if err != nil {
			v = time.Now()
		}
	}
	return v
}

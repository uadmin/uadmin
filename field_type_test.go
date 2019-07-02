package uadmin

import (
	"testing"
)

func TestFieldType(t *testing.T) {
	f := FieldType(0)

	examples := []struct {
		f FieldType
		v int
	}{
		{f.String(), 1},
		{f.Int(), 2},
		{f.Float(), 3},
		{f.Multilingual(), 4},
		{f.Email(), 5},
		{f.Boolean(), 6},
		{f.DateTime(), 7},
		{f.DateTimePtr(), 8},
		{f.ForeignKey(), 9},
		{f.M2M(), 10},
		{f.ProgressBar(), 11},
		{f.HTML(), 12},
		{f.StaticList(), 13},
		{f.File(), 14},
		{f.Image(), 15},
		{f.Money(), 16},
		{f.Code(), 17},
		{f.Link(), 18},
		{f.Password(), 19},
	}

	for _, e := range examples {
		if int(e.f) != e.v {
			t.Errorf("FieldType of value %s returned %d, expected %d", GetString(e.f), int(e.f), e.v)
		}
	}
}

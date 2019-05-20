package uadmin

import (
	"io/ioutil"
	"strings"
)

type Builder struct {
	Model
	Name string
}

func (b *Builder) Save() {
	Save(b)

	fields := []BuilderField{}
	Filter(&fields, "builder_id = ?", b.ID)

	code := `package models
	
import (
	"github.com/uadmin/uadmin"
)
`

	code += "type " + b.Name + " struct{\n"
	code += "\tuadmin.Model\n"

	for _, f := range fields {
		code += "\t" + f.Name + " string\n"
	}

	code += "}"

	ioutil.WriteFile(strings.ToLower("models/"+b.Name+".go"), []byte(code), 0644)
}

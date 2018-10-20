package uadmin

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Experimenting with AST for model recognition
func getStruct() (s []interface{}) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset

	//fName := exPath + "/main.go"

	content, _ := ioutil.ReadFile(exPath + "/main.go")
	f, err := parser.ParseFile(fset, "main.go", content, parser.ParseComments)
	if err != nil {
		Trail(ERROR, "%s", err)
		return
	}

	x := f.Decls[2]

	y := x.(*ast.GenDecl)

	z := y.Specs[0]

	t := z.(*ast.TypeSpec)
	fields := t.Type.(*ast.StructType).Fields.List

	fmt.Printf("%s{\n", t.Name.Name)
	for i := range fields {
		if len(fields[i].Names) > 0 {
			fmt.Printf("  %v  ", fields[i].Names[0].Name)
			fmt.Printf("%v\n", fields[i].Type.(*ast.Ident).Name)
		} else {
			fmt.Printf("  %v\n", fields[i].Type.(*ast.SelectorExpr).Sel.Name)
		}
	}
	fmt.Println("}")

	/*
		x := f.Decls[2]

		s, ok := x.(ast.GenDecl)
		fmt.Println(ok, s)
		//cmap := ast.NewCommentMap(fset, f, f.Comments)
	*/
	return
}

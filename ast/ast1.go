package go_ast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

func GenerateAst(source string) {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", source, 0)
	if err != nil {
		fmt.Printf("parse souce err := err %v", err)
		return
	}
	ast.Print(fs, f)
}

func Inspect(src string) {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", src, 0)
	if err != nil {
		fmt.Printf("parse source err := err %v", err)
		return
	}
	ast.Inspect(f, func(n ast.Node) bool {
		ast.Print(fs, n)
		return true
	})
}

type Visitor struct {
}

func (v *Visitor) Visit(n ast.Node) ast.Visitor {
	switch n := n.(type) {
	case *ast.GenDecl:
		if n.Tok == token.IMPORT {
			v.addImport(n)
			return nil
		}
	case *ast.InterfaceType:
		v.addContext(n)
		return nil
	}
	return v
}

func (v *Visitor) addImport(gd *ast.GenDecl) {
	hasImported := false
	for _, v := range gd.Specs {
		importSpec, ok := v.(*ast.ImportSpec)
		if ok && importSpec.Path.Value == strconv.Quote("context") {
			hasImported = true
			break
		}
	}
	if !hasImported {
		gd.Specs = append(gd.Specs, &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote("context"),
			},
		})
	}
}

func (v *Visitor) addContext(iface *ast.InterfaceType) {
	if iface.Methods != nil && len(iface.Methods.List) > 0 {
		for _, v := range iface.Methods.List {
			ft := v.Type.(*ast.FuncType)
			hasContext := false
			for _, param := range ft.Params.List {
				if expr, ok := param.Type.(*ast.SelectorExpr); ok {
					if ident, ok := expr.X.(*ast.Ident); ok {
						if ident.Name == "context" {
							hasContext = true
							break
						}
					}
				}
			}
			if !hasContext {
				ctxField := &ast.Field{
					Names: []*ast.Ident{ast.NewIdent("ctx")},
					Type: &ast.SelectorExpr{
						X:   ast.NewIdent("context"),
						Sel: ast.NewIdent("Context"),
					},
				}
				list := []*ast.Field{
					ctxField,
				}
				ft.Params.List = append(ft.Params.List, list...)
			}
		}
	}
}

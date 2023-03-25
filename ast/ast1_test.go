package go_ast

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"testing"
)

var sourceCode = `package hello

import "fmt"

func greet() {
    var msg = "Hello World!"
    fmt.Println(msg)
}
`

func TestGenerateAst(t *testing.T) {
	GenerateAst(sourceCode)
}

func TestInspect(t *testing.T) {
	Inspect(sourceCode)
}

func TestAddContextInMethodParams(t *testing.T) {
	var sc = `package main

import (
	"context"
)

type Foo interface {
	FooA(i int)
	FooB(j int)
	FooC(ctx context.Context)
}

type Bar interface {
	BarA(i int)
	BarB()
	BarC()
}`
	fs := token.NewFileSet()
	visitor := &Visitor{}
	f, err := parser.ParseFile(fs, "", sc, 0)
	if err != nil {
		t.Fatalf("parse err %v ", err)
	}
	ast.Walk(visitor, f)
	var output []byte
	buff := bytes.NewBuffer(output)
	if err := format.Node(buff, fs, f); err != nil {
		t.Fatalf("format err %v ", err)
	}
	fmt.Println(buff.String())
}

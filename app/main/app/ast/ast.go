package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "/home/onest/go/src/github.com/olegvn88/gostudy/app/main/app/ast/e.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	for _, decl := range f.Decls {
		fdecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if isExample(fdecl) {
			output, found := findExampleOutput(fdecl.Body, f.Comments)
			if found {
				fmt.Printf("%s needs output '%s'\n", fdecl.Name, output)
			}
		}
	}
}

func findExampleOutput(block *ast.BlockStmt, comments []*ast.CommentGroup) (string, bool) {
	var last *ast.CommentGroup
	for _, group := range comments {
		if (block.Pos() < group.Pos()) && (block.End() > group.End()) {
			last = group
		}
	}

	if last != nil {
		text := last.Text()
		marker := "Output: "
		if strings.HasPrefix(text, marker) {
			return strings.TrimRight(text[len(marker):], "\n"), true
		}
	}
	return "", false
}

func isExample(fdecl *ast.FuncDecl) bool {
	return strings.HasPrefix(fdecl.Name.Name, "Example")
}

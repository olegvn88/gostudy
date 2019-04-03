package main

import (
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, `/home/onest/go/src/github.com/olegvn88/gostudy/app/main/app/ast/e.go`, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	examples := doc.Examples(f)
	for _, example := range examples {
		fmt.Println(example.Name)
		fmt.Println(example.Output)
	}

}

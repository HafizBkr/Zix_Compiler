package main

import (
	"fmt"
	"os"
	"hafizbkrcompiler/parser"
)

func main() {
	file, err := os.Open("input.test")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lexer := parser.NewLexer(file)
	for {
		pos, tok, lit := lexer.Lex()
		if tok == parser.EOF {
			break
		}

		fmt.Printf("%d:%d\t%s\t%s\n", pos.Line, pos.Column, tok, lit)
	}
}

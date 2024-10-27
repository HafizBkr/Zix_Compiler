package main

import (
	"os"
	"hafizbkrcompiler/lexer"
	"hafizbkrcompiler/parser"
)

func main() {
	// Vérifie si le fichier est fourni en argument
	if len(os.Args) < 2 {
		println("Veuillez fournir un fichier à analyser.")
		return
	}

	// Ouvre le fichier
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		println("Erreur lors de l'ouverture du fichier:", err.Error())
		return
	}
	defer file.Close()

	// Crée un nouveau lexer
	lex := lexer.NewLexer(file)
	// Crée un nouveau parser
	p := parser.NewParser(lex)

	// Évalue le code
	p.Evaluate()
}

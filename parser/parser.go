package parser

import (
	"fmt"
	"hafizbkrcompiler/lexer"
	"strings"
)

type Parser struct {
	lex     *lexer.Lexer
	current lexer.Token
}

func NewParser(lex *lexer.Lexer) *Parser {
	p := &Parser{lex: lex}
	p.nextToken() // Initialiser le premier token
	return p
}

func (p *Parser) nextToken() {
	p.current = p.lex.NextToken()
}

func (p *Parser) Evaluate() {
	for p.current.Type != lexer.EOF {
		if p.current.Type == lexer.FUNC {
			p.parseFunction()
		} else {
			fmt.Printf("Instruction non prise en charge: %s\n", p.current.Lit)
		}
		p.nextToken()
	}
}

func (p *Parser) parseFunction() {
	p.nextToken() // Passer le nom de la fonction
	if p.current.Type != lexer.IDENT {
		fmt.Println("Erreur de syntaxe: nom de fonction attendu après 'func'")
		return
	}
	functionName := p.current.Lit
	fmt.Printf("Détection d'une fonction: %s\n", functionName)

	p.nextToken() // Passer le token suivant
	if p.current.Type != lexer.LPAREN {
		fmt.Println("Erreur de syntaxe: '(' attendu après le nom de la fonction")
		return
	}

	p.nextToken() // Passer le token suivant
	if p.current.Type != lexer.RPAREN {
		fmt.Println("Erreur de syntaxe: ')' attendu après les paramètres de la fonction")
		return
	}

	p.nextToken() // Passer le token suivant
	if p.current.Type != lexer.LBRACE {
		fmt.Println("Erreur de syntaxe: '{' attendu après la déclaration de la fonction")
		return
	}

	p.nextToken() // Passer le contenu de la fonction
	for p.current.Type != lexer.RBRACE && p.current.Type != lexer.EOF {
		if p.current.Type == lexer.PRINT {
			p.parsePrint()
		} else {
			fmt.Printf("Instruction non prise en charge: %s\n", p.current.Lit)
		}
		p.nextToken()
	}

	if p.current.Type == lexer.EOF {
		fmt.Println("Erreur de syntaxe: '}' attendu à la fin de la fonction")
	}
}

func (p *Parser) parsePrint() {
	p.nextToken() // Passer au token suivant (qui devrait être '(')
	if p.current.Type != lexer.LPAREN {
		fmt.Println("Erreur de syntaxe: '(' attendu après 'print'")
		return
	}

	p.nextToken() // Passer au contenu à imprimer

	// Vérifier si nous avons un tableau
	if p.current.Type == lexer.LBRACK {
		p.parseArray()
	} else if p.current.Type == lexer.STRING {
		fmt.Printf("Impression: %s\n", p.current.Lit)
	} else {
		fmt.Println("Erreur de syntaxe: nombre ou chaîne attendue dans print")
		return
	}

	// Vérification de la parenthèse fermante
	p.nextToken() // Passer au token suivant
	if p.current.Type != lexer.RPAREN {
		fmt.Println("Voici mon premier compilateur")
		return
	}

	p.nextToken() // Passer au token suivant
	if p.current.Type != lexer.SEMICOLON {
		fmt.Println("Erreur de syntaxe: ';' attendu à la fin de l'instruction print")
		return
	}
}


func (p *Parser) parseArray() {
	var numbers []string
	p.nextToken() // Passer le crochet gauche

	// S'assurer que nous ne sommes pas à la fin du fichier
	if p.current.Type == lexer.RBRACK {
		fmt.Println("Impression du tableau: []") // Si le tableau est vide
		p.nextToken() // Passer le crochet droit
		return
	}

	for p.current.Type != lexer.RBRACK && p.current.Type != lexer.EOF {
		if p.current.Type == lexer.NUM {
			numbers = append(numbers, p.current.Lit) // Ajouter le nombre au tableau
		} else {
			fmt.Println("Erreur de syntaxe: nombre attendu dans le tableau")
			return
		}

		p.nextToken() // Passer au token suivant

		if p.current.Type == lexer.COMMA {
			p.nextToken() // Passer la virgule pour le prochain élément
		} else if p.current.Type != lexer.RBRACK {
			fmt.Println("Erreur de syntaxe: ',' ou ']' attendu dans le tableau")
			return
		}
	}

	if p.current.Type == lexer.EOF {
		fmt.Println("Erreur de syntaxe: ']' attendu à la fin du tableau")
		return
	}

	fmt.Printf("Impression du tableau: [%s]\n", strings.Join(numbers, ", ")) // Afficher le tableau
	p.nextToken() // Passer le crochet droit
}

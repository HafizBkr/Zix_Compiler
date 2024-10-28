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
		switch p.current.Type {
		case lexer.FUNC:
			p.parseFunction()
		case lexer.IDENT:
			if nextToken := p.lex.PeekToken(); nextToken.Type == lexer.ASSIGN {
				p.parseAssignment()
			} else {
				p.parseVariableDeclaration()
			}
		default:
			fmt.Printf("Instruction non prise en charge: %s\n", p.current.Lit)
		}
		p.nextToken()
	}
}

func (p *Parser) parseFunction() {
	p.nextToken() // Passer le mot-clé 'func'
	if p.current.Type != lexer.IDENT {
		fmt.Println("Erreur de syntaxe: nom de fonction attendu après 'func'")
		return
	}
	functionName := p.current.Lit
	fmt.Printf("Détection d'une fonction: %s\n", functionName)

	p.nextToken()
	if p.current.Type != lexer.LPAREN {
		fmt.Println("Erreur de syntaxe: '(' attendu après le nom de la fonction")
		return
	}

	p.nextToken()
	if p.current.Type != lexer.RPAREN {
		fmt.Println("Erreur de syntaxe: ')' attendu après les paramètres de la fonction")
		return
	}

	p.nextToken()
	if p.current.Type != lexer.LBRACE {
		fmt.Println("Erreur de syntaxe: '{' attendu après la déclaration de la fonction")
		return
	}

	p.nextToken()
	for p.current.Type != lexer.RBRACE && p.current.Type != lexer.EOF {
		switch p.current.Type {
		case lexer.PRINT:
			p.parsePrint()
		case lexer.IDENT:
			if p.lex.PeekToken().Type == lexer.ASSIGN {
				p.parseAssignment()
			} else {
				p.parseVariableDeclaration() // Gérer la déclaration de variables ici
			}
		default:
			fmt.Printf("Instruction ignorée dans la fonction: %s\n", p.current.Lit)
		}
		p.nextToken()
	}

	if p.current.Type == lexer.EOF {
		fmt.Println("Erreur de syntaxe: '}' attendu à la fin de la fonction")
	}
}

func (p *Parser) parseVariableDeclaration() {
	if p.current.Type != lexer.IDENT {
		fmt.Println("Erreur de syntaxe: nom de variable attendu après la déclaration")
		return
	}
	varName := p.current.Lit
	fmt.Printf("Déclaration de la variable: %s\n", varName)

	p.nextToken() // Passer au prochain token
	if p.current.Type != lexer.SEMICOLON {
		fmt.Println("Erreur de syntaxe: ';' attendu à la fin de l'instruction variable")
		return
	}
	fmt.Printf("Variable '%s' déclarée avec succès.\n", varName)
}

func (p *Parser) parseAssignment() {
	varName := p.current.Lit // Assurez-vous que la variable a été déclarée
	p.nextToken() // Passer le nom de la variable
	if p.current.Type != lexer.ASSIGN {
		fmt.Println("Erreur de syntaxe: '=' attendu après le nom de la variable")
		return
	}

	p.nextToken() // Passer au token de valeur
	if p.current.Type != lexer.NUM && p.current.Type != lexer.STRING {
		fmt.Println("Erreur de syntaxe: valeur de variable attendue après '='")
		return
	}

	fmt.Printf("Affectation de la variable '%s' à la valeur: %s\n", varName, p.current.Lit)

	p.nextToken() // Passer à l'instruction suivante
	if p.current.Type != lexer.SEMICOLON {
		fmt.Println("Erreur de syntaxe: ';' attendu à la fin de l'instruction d'affectation")
		return
	}
}


func (p *Parser) parsePrint() {
	p.nextToken()
	if p.current.Type != lexer.LPAREN {
		fmt.Println("Erreur de syntaxe: '(' attendu après 'print'")
		return
	}

	p.nextToken()
	if p.current.Type == lexer.LBRACK {
		p.parseArray()
	} else if p.current.Type == lexer.STRING || p.current.Type == lexer.NUM || p.current.Type == lexer.IDENT {
		if p.current.Type == lexer.STRING {
			fmt.Printf("Impression: %s\n", p.current.Lit)
		} else if p.current.Type == lexer.NUM {
			fmt.Printf("Impression du nombre: %s\n", p.current.Lit)
		} else if p.current.Type == lexer.IDENT {
			fmt.Printf("Impression de la variable: %s\n", p.current.Lit)
		}
	} else {
		fmt.Println("Erreur de syntaxe: nombre ou chaîne attendue dans print")
		return
	}

	p.nextToken()
	if p.current.Type != lexer.RPAREN {
		fmt.Println("Erreur de syntaxe: ')' attendu à la fin de l'instruction print")
		return
	}

	p.nextToken()
	if p.current.Type != lexer.SEMICOLON {
		fmt.Println("Erreur de syntaxe: ';' attendu à la fin de l'instruction print")
		return
	}
}



func (p *Parser) parseArray() {
	var elements []string
	p.nextToken()

	if p.current.Type == lexer.RBRACK {
		fmt.Println("Impression du tableau: []")
		p.nextToken()
		return
	}

	for p.current.Type != lexer.RBRACK && p.current.Type != lexer.EOF {
		if p.current.Type == lexer.NUM {
			elements = append(elements, p.current.Lit)
		} else {
			fmt.Println("Erreur de syntaxe: nombre attendu dans le tableau")
			return
		}

		p.nextToken()
		if p.current.Type == lexer.COMMA {
			p.nextToken()
		} else if p.current.Type != lexer.RBRACK {
			fmt.Println("Erreur de syntaxe: ',' ou ']' attendu dans le tableau")
			return
		}
	}

	if p.current.Type == lexer.EOF {
		fmt.Println("Erreur de syntaxe: ']' attendu à la fin du tableau")
		return
	}

	fmt.Printf("Impression du tableau: [%s]\n", strings.Join(elements, ", "))
	p.nextToken()
}

package lexer

import (
	"bufio"
	"io"
	"unicode"
)

// Types de token
const (
	ILLEGAL = iota
	EOF
	WS
	IDENT
	STRING
	NUM
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	SEMICOLON
	FUNC
	PRINT
	LBRACK // Ajout pour le crochet gauche
	RBRACK // Ajout pour le crochet droit
	COMMA  // Ajout pour la virgule
)

type Token struct {
	Type int
	Lit  string
}

type Lexer struct {
	input       *bufio.Reader // Changer io.Reader à *bufio.Reader
	ch          rune
	currentLine int
}

func NewLexer(input io.Reader) *Lexer {
	l := &Lexer{
		input: bufio.NewReader(input), // Lire en utilisant bufio
	}
	l.next() // Initialiser le premier caractère
	return l
}

func (l *Lexer) next() rune {
	ch, _, err := l.input.ReadRune() // Utiliser l.input qui est un *bufio.Reader
	if err != nil {
		if err == io.EOF {
			return 0
		}
		panic(err)
	}
	if ch == '\n' {
		l.currentLine++
	}
	l.ch = ch
	return l.ch
}

func (l *Lexer) NextToken() Token {
	for unicode.IsSpace(l.ch) {
		l.next()
	}

	switch l.ch {
	case 0:
		return Token{Type: EOF}
	case '(':
		l.next()
		return Token{Type: LPAREN, Lit: string('(')}
	case ')':
		l.next()
		return Token{Type: RPAREN, Lit: string(')')}
	case '{':
		l.next()
		return Token{Type: LBRACE, Lit: string('{')}
	case '}':
		l.next()
		return Token{Type: RBRACE, Lit: string('}')}
	case ';':
		l.next()
		return Token{Type: SEMICOLON, Lit: string(';')}
	case '[':
		l.next()
		return Token{Type: LBRACK, Lit: string('[')}
	case ']':
		l.next()
		return Token{Type: RBRACK, Lit: string(']')}
	case ',':
		l.next()
		return Token{Type: COMMA, Lit: string(',')}
	case 'f':
		l.next()
		if l.ch == 'u' {
			l.next()
			if l.ch == 'n' {
				l.next()
				if l.ch == 'c' {
					l.next()
					return Token{Type: FUNC, Lit: "func"}
				}
			}
		}
	case 'p':
		l.next()
		if l.ch == 'r' {
			l.next()
			if l.ch == 'i' {
				l.next()
				if l.ch == 'n' {
					l.next()
					if l.ch == 't' {
						l.next()
						return Token{Type: PRINT, Lit: "print"}
					}
				}
			}
		}
	case '"':
		return Token{Type: STRING, Lit: l.readString()}
	}

	if unicode.IsDigit(l.ch) {
		return l.readNumber() // Appeler readNumber si le caractère est un chiffre
	}

	if unicode.IsLetter(l.ch) {
		return l.readIdentifier()
	}

	return Token{Type: ILLEGAL, Lit: string(l.ch)}
}

func (l *Lexer) readString() string {
	var lit []rune
	l.ch = l.next() // Passer le premier guillemet

	for l.ch != '"' {
		if l.ch == 0 {
			break // Terminer si on atteint la fin de l'entrée
		}
		lit = append(lit, l.ch)
		l.ch = l.next()
	}

	l.ch = l.next() // Passer le guillemet de fin
	return string(lit)
}

func (l *Lexer) readIdentifier() Token {
	var ident []rune
	for unicode.IsLetter(l.ch) || unicode.IsDigit(l.ch) {
		ident = append(ident, l.ch)
		l.ch = l.next()
	}
	return Token{Type: IDENT, Lit: string(ident)}
}

// Ajout de la méthode readNumber
func (l *Lexer) readNumber() Token {
	var number []rune
	for unicode.IsDigit(l.ch) {
		number = append(number, l.ch)
		l.ch = l.next()
	}
	return Token{Type: NUM, Lit: string(number)}
}

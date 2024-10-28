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
	LBRACK
	RBRACK
	COMMA
	ASSIGN = iota + 13
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	VARIABLE
)

type Token struct {
	Type int
	Lit  string
}

type Lexer struct {
	input       *bufio.Reader
	ch          rune
	currentLine int
}

// NewLexer crée un nouveau lexer à partir d'un io.Reader
func NewLexer(input io.Reader) *Lexer {
	l := &Lexer{
		input: bufio.NewReader(input),
	}
	l.next() // Initialiser le premier caractère
	return l
}

// next lit le prochain caractère de l'entrée
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

// PeekToken retourne le prochain token sans avancer le lexer
func (l *Lexer) PeekToken() Token {
	// Conserver l'état actuel
	currentChar := l.ch
	currentLine := l.currentLine
	tokens := []Token{}

	// Lire le prochain token
	token := l.NextToken()
	tokens = append(tokens, token)

	// Réinitialiser l'état
	l.ch = currentChar
	l.currentLine = currentLine

	return token
}

// NextToken lit le prochain token de l'entrée
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
	case '=':
		l.next()
		return Token{Type: ASSIGN, Lit: "="}
	case '+':
		l.next()
		return Token{Type: PLUS, Lit: "+"}
	case '-':
		l.next()
		return Token{Type: MINUS, Lit: "-"}
	case '*':
		l.next()
		return Token{Type: MULTIPLY, Lit: "*"}
	case '/':
		l.next()
		return Token{Type: DIVIDE, Lit: "/"}
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

// readString lit une chaîne de caractères
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

// readIdentifier lit un identifiant
func (l *Lexer) readIdentifier() Token {
	var ident []rune
	for unicode.IsLetter(l.ch) || unicode.IsDigit(l.ch) {
		ident = append(ident, l.ch)
		l.ch = l.next()
	}
	literal := string(ident)
	tokenType := lookupIdent(literal)
	return Token{Type: tokenType, Lit: literal}
}

// lookupIdent retourne le type de token correspondant à l'identifiant
func lookupIdent(ident string) int {
	switch ident {
	case "func":
		return FUNC
	case "print":
		return PRINT
	case "variable":
		return VARIABLE
	default:
		return IDENT
	}
}

// readNumber lit un nombre
func (l *Lexer) readNumber() Token {
	var number []rune
	for unicode.IsDigit(l.ch) {
		number = append(number, l.ch)
		l.ch = l.next()
	}
	return Token{Type: NUM, Lit: string(number)}
}

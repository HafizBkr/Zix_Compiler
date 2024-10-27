package parser

import (
	"fmt"
	"strconv"
)

// Types de nœuds de l'AST
type Node interface {
	Evaluate() int // Méthode d'évaluation
	Optimize() Node // Méthode d'optimisation
}

type IntegerNode struct {
	Value int
}

type OperatorNode struct {
	Left     Node
	Operator Token
	Right    Node
}

// Nœud pour les instructions if
type IfNode struct {
	Condition Node
	Then      Node
	Else      Node // Optionnel
}

// Parser structure
type Parser struct {
	lexer    *Lexer
	curToken Token
	curLit   string
	curPos   Position
}

// Crée un nouveau parser avec un lexer
func NewParser(lexer *Lexer) *Parser {
	p := &Parser{lexer: lexer}
	p.nextToken()
	return p
}

// Avance au prochain token
func (p *Parser) nextToken() {
	p.curPos, p.curToken, p.curLit = p.lexer.Lex()
}

// Parse une expression complète
func (p *Parser) ParseExpression() (Node, error) {
	return p.parseBinaryExpression(0)
}

// Parse une expression binaire en fonction de la priorité
func (p *Parser) parseBinaryExpression(minPrecedence int) (Node, error) {
	left, err := p.parsePrimary()
	if err != nil {
		return nil, err
	}

	for p.curToken == ADD || p.curToken == SUB || p.curToken == MUL || p.curToken == DIV {
		precedence := p.getPrecedence(p.curToken)
		if precedence < minPrecedence {
			break
		}

		operator := p.curToken
		p.nextToken()

		right, err := p.parseBinaryExpression(precedence + 1)
		if err != nil {
			return nil, err
		}

		left = &OperatorNode{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}
	return left, nil
}

// Parse une valeur primaire (entier ou variable)
func (p *Parser) parsePrimary() (Node, error) {
	switch p.curToken {
	case INT:
		value, err := strconv.Atoi(p.curLit)
		if err != nil {
			return nil, fmt.Errorf("invalid integer: %s", p.curLit)
		}
		p.nextToken()
		return &IntegerNode{Value: value}, nil
	case IDENT:
		// Gérer les variables ici si nécessaire
		p.nextToken()
		return nil, fmt.Errorf("unexpected token: %s", p.curToken)
	default:
		return nil, fmt.Errorf("unexpected token: %s", p.curToken)
	}
}

// Définit la priorité des opérateurs
func (p *Parser) getPrecedence(token Token) int {
	switch token {
	case MUL, DIV:
		return 2
	case ADD, SUB:
		return 1
	default:
		return 0
	}
}

// Évaluation de l'AST
func (n *IntegerNode) Evaluate() int {
	return n.Value
}

func (n *IntegerNode) Optimize() Node {
	return n // Les entiers ne nécessitent pas d'optimisation
}

func (n *OperatorNode) Evaluate() int {
	leftValue := n.Left.Evaluate()  // Évalue le sous-arbre gauche
	rightValue := n.Right.Evaluate() // Évalue le sous-arbre droit

	switch n.Operator {
	case ADD:
		return leftValue + rightValue
	case SUB:
		return leftValue - rightValue
	case MUL:
		return leftValue * rightValue
	case DIV:
		if rightValue == 0 {
			panic("division par zéro")
		}
		return leftValue / rightValue
	default:
		panic("opérateur inconnu")
	}
}

func (n *OperatorNode) Optimize() Node {
	left := n.Left.Optimize()
	right := n.Right.Optimize()

	if n.Operator == ADD {
		if intNode, ok := left.(*IntegerNode); ok && intNode.Value == 0 {
			return right // x + 0 -> x
		}
		if intNode, ok := right.(*IntegerNode); ok && intNode.Value == 0 {
			return left // 0 + x -> x
		}
	}

	// Ajoutez d'autres optimisations ici

	return &OperatorNode{Left: left, Operator: n.Operator, Right: right}
}

// Évaluation pour IfNode
func (n *IfNode) Evaluate() int {
	if n.Condition.Evaluate() != 0 { // Supposons que la condition est une expression qui retourne un int
		return n.Then.Evaluate()
	} else if n.Else != nil {
		return n.Else.Evaluate()
	}
	return 0
}

func (n *IfNode) Optimize() Node {
	// Gérer l'optimisation pour les nœuds IfNode si nécessaire
	return n
}

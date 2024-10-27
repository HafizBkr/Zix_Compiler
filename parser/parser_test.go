package parser

import (
	"strings"
	"testing"
)

func TestArithmeticParser(t *testing.T) {
	input := `3 + 5 * 2 - 8 / 4;`

	// Créer un nouveau lexer avec la chaîne d'entrée
	lexer := NewLexer(strings.NewReader(input))
	parser := NewParser(lexer)

	expectedAST := &OperatorNode{
		Left: &OperatorNode{
			Left:     &IntegerNode{Value: 3},
			Operator: ADD,
			Right: &OperatorNode{
				Left:     &IntegerNode{Value: 5},
				Operator: MUL,
				Right:    &IntegerNode{Value: 2},
			},
		},
		Operator: SUB,
		Right: &OperatorNode{
			Left:     &IntegerNode{Value: 8},
			Operator: DIV,
			Right:    &IntegerNode{Value: 4},
		},
	}

	// Parse l'expression
	ast, err := parser.ParseExpression()
	if err != nil {
		t.Fatalf("ParseExpression() error: %v", err)
	}

	// Vérifiez si l'AST généré correspond à l'AST attendu
	if !compareAST(ast, expectedAST) {
		t.Errorf("expected AST %+v, got %+v", expectedAST, ast)
	}
}

// compareAST est une fonction qui compare deux nœuds AST pour l'égalité
func compareAST(a, b Node) bool {
	switch aNode := a.(type) {
	case *IntegerNode:
		bNode, ok := b.(*IntegerNode)
		return ok && aNode.Value == bNode.Value
	case *OperatorNode:
		bOpNode, ok := b.(*OperatorNode)
		if !ok || aNode.Operator != bOpNode.Operator {
			return false
		}
		return compareAST(aNode.Left, bOpNode.Left) && compareAST(aNode.Right, bOpNode.Right)
	default:
		return false
	}
}

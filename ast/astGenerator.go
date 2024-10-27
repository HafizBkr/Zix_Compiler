package ast

import "fmt"

// Node représente un nœud de l'AST.
type Node interface {
	String() string
}

// Program représente le nœud racine d'un programme.
type Program struct {
	Statements []Node
}

// FunctionParameter représente un paramètre de fonction.
type FunctionParameter struct {
	Name string
	Type string // Type du paramètre (par exemple, "int", "string", etc.)
}

// FunctionDeclaration représente une déclaration de fonction.
type FunctionDeclaration struct {
	Name       string
	Parameters []*FunctionParameter // Liste des paramètres
	ReturnType string               // Type de retour de la fonction (par exemple, "int", "void", etc.)
	Body       []Node
}

// ReturnStatement représente une instruction de retour.
type ReturnStatement struct {
	Expression Node
}

// IfStatement représente une instruction conditionnelle.
type IfStatement struct {
	Condition Node // Condition de l'instruction if
	Then      []Node // Corps de l'instruction if
	Else      []Node // Corps de l'instruction else (optionnel)
}

// Identifier représente une variable ou un nom d'identifiant.
type Identifier struct {
	Name string
}

// BinaryExpression représente une expression binaire.
type BinaryExpression struct {
	Left     Node
	Operator string
	Right    Node
}

// Number représente une valeur numérique.
type Number struct {
	Value int64
}

// String permet d'afficher l'AST sous forme de chaîne.
func (p *Program) String() string {
	var result string
	for _, stmt := range p.Statements {
		result += stmt.String() + "\n"
	}
	return result
}

func (f *FunctionDeclaration) String() string {
	var params []string
	for _, param := range f.Parameters {
		params = append(params, fmt.Sprintf("%s %s", param.Name, param.Type))
	}
	body := ""
	for _, stmt := range f.Body {
		body += stmt.String() + "\n"
	}
	return fmt.Sprintf("func %s(%s) %s {\n%s}", f.Name, fmt.Sprintf("%s", params), f.ReturnType, body)
}

func (r *ReturnStatement) String() string {
	return fmt.Sprintf("return %s;", r.Expression.String())
}

func (i *Identifier) String() string {
	return i.Name
}

func (b *BinaryExpression) String() string {
	return fmt.Sprintf("%s %s %s", b.Left.String(), b.Operator, b.Right.String())
}

func (n *Number) String() string {
	return fmt.Sprintf("%d", n.Value)
}

func (i *IfStatement) String() string {
	thenBody := ""
	for _, stmt := range i.Then {
		thenBody += stmt.String() + "\n"
	}
	elseBody := ""
	for _, stmt := range i.Else {
		elseBody += stmt.String() + "\n"
	}
	if elseBody != "" {
		return fmt.Sprintf("if %s {\n%s} else {\n%s}", i.Condition.String(), thenBody, elseBody)
	}
	return fmt.Sprintf("if %s {\n%s}", i.Condition.String(), thenBody)
}

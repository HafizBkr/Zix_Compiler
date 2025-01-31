package codegen

import (
	"hafizbkrcompiler/ast"
	"strings"
)

type CodeGenerator struct {
	code strings.Builder
}

// NewCodeGenerator crée un nouveau générateur de code.
func NewCodeGenerator() *CodeGenerator {
	return &CodeGenerator{}
}

// Generate génère le code pour un programme.
func (cg *CodeGenerator) Generate(program *ast.Program) string {
	for _, stmt := range program.Statements {
		cg.generateStatement(stmt)
	}
	return cg.code.String()
}

// generateStatement génère du code pour une déclaration.
func (cg *CodeGenerator) generateStatement(stmt ast.Node) {
	switch s := stmt.(type) {
	case *ast.FunctionDeclaration:
		// Extraire les noms des paramètres
		paramNames := make([]string, len(s.Parameters))
		for i, param := range s.Parameters {
			paramNames[i] = param.Name // Assurez-vous que FunctionParameter a un champ Name
		}

		// Générez la déclaration de la fonction
		cg.code.WriteString("func " + s.Name + "(" + strings.Join(paramNames, ", ") + ") {\n")
		for _, bodyStmt := range s.Body {
			cg.generateStatement(bodyStmt)
		}
		cg.code.WriteString("}\n")
	case *ast.ReturnStatement:
		// Générer l'instruction de retour
		cg.code.WriteString("return " + s.Expression.String() + "\n")
	case *ast.Identifier:
		// Identifier
		cg.code.WriteString(s.Name)
	case *ast.BinaryExpression:
		// Générer les expressions binaires
		cg.code.WriteString("(" + s.Left.String() + " " + s.Operator + " " + s.Right.String() + ")")
	default:
		// Gestion des cas supplémentaires (par exemple, des expressions, des instructions de déclaration, etc.)
		cg.code.WriteString("// Autre déclaration\n")
	}
}

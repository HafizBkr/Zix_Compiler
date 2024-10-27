package typechecker

import (
	"errors"
	"hafizbkrcompiler/ast"
	"hafizbkrcompiler/types"
)

// TypeChecker vérifie les types d'un programme.
type TypeChecker struct {
	symbolTable *types.SymbolTable
}

// NewTypeChecker crée un nouveau vérificateur de types.
func NewTypeChecker() *TypeChecker {
	return &TypeChecker{symbolTable: types.NewSymbolTable()}
}

// Check vérifie un programme.
func (tc *TypeChecker) Check(program *ast.Program) error {
	for _, stmt := range program.Statements {
		if err := tc.checkStatement(stmt); err != nil {
			return err
		}
	}
	return nil
}

// checkStatement vérifie une déclaration.
func (tc *TypeChecker) checkStatement(stmt ast.Node) error {
	switch s := stmt.(type) {
	case *ast.FunctionDeclaration:
		// Convertir le ReturnType en type Type
		returnType := types.Type(s.ReturnType)
		tc.symbolTable.Define(s.Name, returnType) // Définit le type de retour dans la table des symboles
		for _, param := range s.Parameters {
			paramType := types.Type(param.Type) // Convertir le type du paramètre
			tc.symbolTable.Define(param.Name, paramType) // Tous les paramètres
		}
		for _, bodyStmt := range s.Body {
			if err := tc.checkStatement(bodyStmt); err != nil {
				return err
			}
		}
	case *ast.ReturnStatement:
		// Vérification simplifiée, ajoutez votre logique ici
		if _, ok := s.Expression.(*ast.BinaryExpression); !ok {
			return errors.New("expression de retour non valide")
		}
	default:
		// Autres types d'instructions
	}
	return nil
}

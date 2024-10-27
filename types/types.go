package types

// Type représente un type de base
type Type string

const (
	IntType      Type = "int"
	StringType   Type = "string"
	BoolType     Type = "bool"
	FloatType    Type = "float"
	ArrayType     Type = "array"
	FuncType     Type = "func"
	ErrorType    Type = "error"
)

// SymbolTable représente une table de symboles pour stocker les types des variables.
type SymbolTable struct {
	symbols map[string]Type
}

// NewSymbolTable crée une nouvelle table de symboles
func NewSymbolTable() *SymbolTable {
	return &SymbolTable{symbols: make(map[string]Type)}
}

// Define ajoute un nouveau symbole et son type dans la table
func (st *SymbolTable) Define(name string, typ Type) {
	st.symbols[name] = typ
}

// Lookup recherche un symbole dans la table et retourne son type
func (st *SymbolTable) Lookup(name string) (Type, bool) {
	typ, ok := st.symbols[name]
	return typ, ok
}

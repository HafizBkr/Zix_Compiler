package parser

type Environment struct {
    variables map[string]int
    functions map[string]func(int) int // Adaptez selon vos besoins
}

// NewEnvironment crée un nouvel environnement
func NewEnvironment() *Environment {
    return &Environment{
        variables: make(map[string]int),
        functions: make(map[string]func(int) int),
    }
}

// DefineFunction pour ajouter des fonctions à l'environnement
func (env *Environment) DefineFunction(name string, fn func(int) int) {
    env.functions[name] = fn
}

 # Zix Compiler

Zix est un compilateur écrit en Go, conçu pour analyser et exécuter un langage de programmation personnalisé. Ce projet comprend les composants essentiels d'un compilateur, notamment le lexer, le parser, le générateur de code, et plus encore.

## Table des matières

- [Fonctionnalités](#fonctionnalités)
- [Architecture](#architecture)
- [Installation](#installation)
- [Utilisation](#utilisation)
- [Contribuer](#contribuer)
- [License](#license)

## Fonctionnalités

- Analyse lexicale avec un lexer .
- Analyse syntaxique avec un parser.
- Génération de code à partir de l'arbre de syntaxe abstraite (AST).
- Vérification de type pour assurer la cohérence du code.
- Exécution de programmes écrits dans le langage personnalisé.

## Architecture

Le projet est structuré comme suit :

.
├── ast
│   └── astGenerator.go
├── codegen
│   └── codegen.go
├── code.hz
├── go.mod
├── input.test
├── lexer
│   └── lexer.go
├── main.go
├── parser
│   ├── environement.go
│   └── parser.go
├── README.md
├── typecheker
│   └── typeCheker.go
└── types
    └── types.go


## Installation

1. Clonez le dépôt:

   ```bash
   git clone https://github.com/HafizBkr/Zix_Compier.git
Accédez au répertoire du projet  :

      
      cd Zix_Compier
      
Installez les dépendances (si nécessaire)  :


      go mod tidy

      
Utilisation
Pour exécuter le compilateur, utilisez la commande suivante en spécifiant un fichier de code source dans le langage personnalisé.


      go run main.go code.hz


Assurez-vous que code.hz contient le code que vous souhaitez compiler et exécuter..:

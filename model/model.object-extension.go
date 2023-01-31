package model

import "github.com/graphql-go/graphql/language/ast"

type ObjectExtension struct {
	Def    *ast.TypeExtensionDefinition
	Model  *Model
	Object *Object
}

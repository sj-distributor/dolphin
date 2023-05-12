package model

import "github.com/graphql-go/graphql/language/ast"

type ObjectExtension struct {
	Def    *ast.TypeExtensionDefinition
	Model  *Model
	Object *Object
}

func (oe *ObjectExtension) IsFederatedType() bool {
	return oe.Object.IsFederatedType()
}

func (oe *ObjectExtension) ExtendsLocalObject() bool {
	return oe.Model.HasObject(oe.Object.Name())
}

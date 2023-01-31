package model

import (
	"github.com/graphql-go/graphql/language/ast"
)

type ObjectRelationship struct {
	Def *ast.FieldDefinition
	Obj *Object
}

func (o *ObjectRelationship) Name() string {
	return o.Def.Name.Value
}

func (o *ObjectRelationship) IsToMany() bool {
	t := getNullableType(o.Def.Type)
	return isListType(t)
}

func (o *ObjectRelationship) IsToOne() bool {
	return !o.IsToMany()
}

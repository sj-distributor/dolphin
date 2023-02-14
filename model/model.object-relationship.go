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

func (o *ObjectRelationship) Target() *Object {
	target := o.Obj.Model.Object(o.TargetType())
	return &target
}
func (o *ObjectRelationship) TargetType() string {
	nt := getNamedType(o.Def.Type).(*ast.Named)
	return nt.Name.Value
}

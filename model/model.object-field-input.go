package model

import (
  "github.com/graphql-go/graphql/language/ast"
)

type ObjectFieldInput struct {
  Def *ast.InputValueDefinition
  Field *ObjectField
}

func (o *ObjectFieldInput) Name() string {
  return o.Def.Name.Value
}
func (o *ObjectFieldInput) IsListType() bool {
  return isListType(getNullableType(o.Def.Type))
}

func (o *ObjectFieldInput) TargetType() string {
  nt := getNamedType(o.Def.Type).(*ast.Named)
  return nt.Name.Value
}

func (o *ObjectFieldInput) NonNullType() string {
  nullType  := ""
  if o.Required() {
    nullType = "!"
  }
  return nullType
}

func (o *ObjectFieldInput) Required() bool {
  isEmpty := false
  bool := isNonNullType(o.Def.Type)
  if bool {
    isEmpty = true
  }
  return isEmpty
}
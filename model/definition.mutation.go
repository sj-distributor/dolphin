package model

import (
	"github.com/graphql-go/graphql/language/kinds"
	"github.com/jinzhu/inflection"

	"github.com/graphql-go/graphql/language/ast"
)

var idInput = ast.InputValueDefinition{
	Kind: kinds.InputValueDefinition,
	Name: nameNode("id"),
	Type: nonNull(namedType("ID")),
}

func mutationDefinition(m *Model) *ast.ObjectDefinition {
	fields := []*ast.FieldDefinition{}

	for _, obj := range m.ObjectEntities() {
		// fields = append(fields, createFieldDefinition(obj), updateFieldDefinition(obj), deleteFieldDefinition(obj), recoveryFieldDefinition(obj))
		fields = append(fields, createFieldDefinition(obj), updateFieldDefinition(obj))
	}
	return &ast.ObjectDefinition{
		Kind:   kinds.ObjectDefinition,
		Name:   nameNode("Mutation"),
		Fields: fields,
	}
}

func createFieldInput(obj Object) *ast.InputValueDefinition {
	d := createObjectDefinition(obj)
	return &ast.InputValueDefinition{
		Kind: kinds.InputValueDefinition,
		Name: nameNode("input"),
		Type: nonNull(namedType(d.Name.Value)),
	}
}

func createFieldDefinition(obj Object) *ast.FieldDefinition {
	return &ast.FieldDefinition{
		Kind: kinds.FieldDefinition,
		Name: nameNode("create" + inflection.Singular(obj.Name())),
		Type: nonNull(namedType(obj.Name())),
		Arguments: []*ast.InputValueDefinition{
			createFieldInput(obj),
		},
	}
}

func updateFieldInput(obj Object) *ast.InputValueDefinition {
	d := updateObjectDefinition(obj)
	return &ast.InputValueDefinition{
		Kind: kinds.InputValueDefinition,
		Name: nameNode("input"),
		Type: nonNull(namedType(d.Name.Value)),
	}
}

func updateFieldDefinition(obj Object) *ast.FieldDefinition {
	return &ast.FieldDefinition{
		Kind: kinds.FieldDefinition,
		Name: nameNode("update" + inflection.Singular(obj.Name())),
		Type: nonNull(namedType(obj.Name())),
		Arguments: []*ast.InputValueDefinition{
			&idInput,
			updateFieldInput(obj),
		},
	}
}

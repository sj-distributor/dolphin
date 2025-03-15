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
		fields = append(fields, createFieldDefinition(obj), updateFieldDefinition(obj), deleteFieldDefinition(obj), recoveryFieldDefinition(obj))
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
		// Description: &ast.StringValue{Kind: kinds.StringValue, Value: "MutationCreate" + strcase.ToCamel(obj.Name()) + "Args"},
	}
}

// MutationUpdateUserArgs
func createFieldDefinition(obj Object) *ast.FieldDefinition {
	return &ast.FieldDefinition{
		Kind:       kinds.FieldDefinition,
		Name:       nameNode("create" + inflection.Singular(obj.Name())),
		Type:       nonNull(namedType(obj.Name())),
		Directives: createObjectHasRoleEnum(obj),
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
		// Description: &ast.StringValue{Kind: kinds.StringValue, Value: "MutationUpdate" + strcase.ToCamel(obj.Name()) + "Args"},
	}
}

func updateFieldDefinition(obj Object) *ast.FieldDefinition {
	return &ast.FieldDefinition{
		Kind:       kinds.FieldDefinition,
		Name:       nameNode("update" + inflection.Singular(obj.Name())),
		Type:       nonNull(namedType(obj.Name())),
		Directives: createObjectHasRoleEnum(obj),
		Arguments: []*ast.InputValueDefinition{
			&idInput,
			updateFieldInput(obj),
		},
	}
}

func deleteFieldDefinition(obj Object) *ast.FieldDefinition {
	return &ast.FieldDefinition{
		Kind:       kinds.FieldDefinition,
		Name:       nameNode("delete" + inflection.Plural(obj.Name())),
		Type:       nonNull(namedType("Boolean")),
		Directives: createObjectHasRoleEnum(obj),
		Arguments: []*ast.InputValueDefinition{
			{
				Kind: kinds.InputValueDefinition,
				Name: nameNode("id"),
				Type: nonNull(listType(nonNull(namedType("ID")))),
			},
			{
				Kind:         kinds.InputValueDefinition,
				Name:         nameNode("unscoped"),
				DefaultValue: &ast.IntValue{Kind: kinds.IntValue, Value: "false"},
				Type:         namedType("Boolean"),
				// Description:  &ast.StringValue{Kind: kinds.StringValue, Value: "MutationDelete" + inflection.Plural(strcase.ToCamel(obj.Name())) + "Args"},
			},
		},
	}
}

func recoveryFieldDefinition(obj Object) *ast.FieldDefinition {
	return &ast.FieldDefinition{
		Kind:       kinds.FieldDefinition,
		Name:       nameNode("recovery" + inflection.Plural(obj.Name())),
		Type:       nonNull(namedType("Boolean")),
		Directives: createObjectHasRoleEnum(obj),
		Arguments: []*ast.InputValueDefinition{
			{
				Kind: kinds.InputValueDefinition,
				Name: nameNode("id"),
				Type: nonNull(listType(nonNull(namedType("ID")))),
				// Description: &ast.StringValue{Kind: kinds.StringValue, Value: "MutationRecovery" + inflection.Plural(strcase.ToCamel(obj.Name())) + "Args"},
			},
		},
	}
}

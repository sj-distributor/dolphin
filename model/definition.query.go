package model

import (
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

func queryDefinition(m *Model) *ast.ObjectDefinition {
	fields := []*ast.FieldDefinition{
		// createFederationServiceQueryField(),
	}

	if m.HasFederatedTypes() {
		fields = append(fields, createFederationEntitiesQueryField())
	}

	for _, obj := range m.ObjectEntities() {
		fields = append(fields, fetchFieldDefinition(obj), listFieldDefinition(obj))
	}
	return &ast.ObjectDefinition{
		Kind: kinds.ObjectDefinition,
		Name: &ast.Name{
			Kind:  kinds.Name,
			Value: "Query",
		},
		Fields: fields,
	}
}

func fetchFieldDefinition(obj Object) *ast.FieldDefinition {
	return &ast.FieldDefinition{
		Kind: kinds.FieldDefinition,
		Name: nameNode(inflection.Singular(strcase.ToLowerCamel(obj.Name()))),
		Type: namedType(obj.Name()),
		Arguments: []*ast.InputValueDefinition{
			{
				Kind: kinds.InputValueDefinition,
				Name: nameNode("id"),
				// Description: &ast.StringValue{Kind: kinds.StringValue, Value: "Input for searching by object ID"}, // 这是描述说明
				Type: nonNull(namedType("ID")),
			},
		},
	}
}

func listFieldDefinition(obj Object) *ast.FieldDefinition {
	// createObjectSortType(obj)
	return &ast.FieldDefinition{
		Kind: kinds.FieldDefinition,
		Name: nameNode(inflection.Plural(strcase.ToLowerCamel(obj.Name()))),
		Type: namedType(obj.Name() + "ResultType"),
		Arguments: []*ast.InputValueDefinition{
			{
				Kind:         kinds.InputValueDefinition,
				Name:         nameNode("current_page"),
				DefaultValue: &ast.IntValue{Kind: kinds.IntValue, Value: "1"},
				Type:         namedType("Int"),
			},
			{
				Kind:         kinds.InputValueDefinition,
				Name:         nameNode("per_page"),
				DefaultValue: &ast.IntValue{Kind: kinds.IntValue, Value: "10"},
				Type:         namedType("Int"),
			},

			{
				Kind: kinds.InputValueDefinition,
				Name: nameNode("q"),
				Type: namedType("String"),
			},
			{
				Kind: kinds.InputValueDefinition,
				Name: nameNode("sort"),
				Type: listType(nonNull(namedType(obj.Name() + "SortType"))),
			},
			{
				Kind: kinds.InputValueDefinition,
				Name: nameNode("filter"),
				Type: namedType(obj.Name() + "FilterType"),
			},
			{
				Kind:         kinds.InputValueDefinition,
				Name:         nameNode("rand"),
				DefaultValue: &ast.IntValue{Kind: kinds.IntValue, Value: "false"},
				Type:         namedType("Boolean"),
			},
		},
	}
}

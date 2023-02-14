package model

import (
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
)

func createFederationEntitiesQueryField() *ast.FieldDefinition {
	return &ast.FieldDefinition{
		Kind: kinds.FieldDefinition,
		Name: nameNode("_entities"),
		Type: nonNull(listType(namedType("_Entity"))),
		Arguments: []*ast.InputValueDefinition{
			{
				Kind: kinds.InputValueDefinition,
				Name: nameNode("representations"),
				Type: nonNull(listType(nonNull(namedType("_Any")))),
			},
		},
	}
}

package model

import (
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
)

// EnrichModelObjects ...
func EnrichModelObjects(m *Model) error {
	id := columnDefinition("id", "ID", "ID", true)
	isDelete := columnDefinition("isDelete", "Int", "Int", false)
	weight := columnDefinition("weight", "Int", "Int", false)
	state := columnDefinition("state", "Int", "Int", false)
	createdAt := columnDefinition("createdAt", "Int", "Int", true)
	updatedAt := columnDefinition("updatedAt", "Int", "Int", false)
	deletedAt := columnDefinition("deletedAt", "Int", "Int", false)
	createdBy := columnDefinition("createdBy", "ID", "ID", false)
	updatedBy := columnDefinition("updatedBy", "ID", "ID", false)
	deletedBy := columnDefinition("deletedBy", "ID", "ID", false)

	for _, o := range m.ObjectEntities() {
		o.Def.Fields = append([]*ast.FieldDefinition{id}, o.Def.Fields...)
		for _, rel := range o.Relationships() {
			if rel.IsToOne() {
				o.Def.Fields = append(o.Def.Fields, columnDefinition(rel.Name()+"Id", "ID", rel.Target().Name(), false))
				// o.Def.Fields = append(o.Def.Fields, columnDefinition(rel.Name()+"Id", rel.Target().Name(), false))
			}
		}
		o.Def.Fields = append(o.Def.Fields, isDelete, weight, state, deletedBy, updatedBy, createdBy, deletedAt, updatedAt, createdAt)
	}
	return nil
}

// EnrichModel ...
func EnrichModel(m *Model) error {
	definitions := []ast.Node{}
	for _, o := range m.ObjectEntities() {
		for _, rel := range o.Relationships() {
			if rel.IsToMany() {
				o.Def.Fields = append(o.Def.Fields, columnDefinitionWithType(rel.Name()+"Ids", kinds.FieldDefinition, listType(nonNull(namedType("ID")))))
			}
		}
		definitions = append(definitions, createObjectDefinition(o), updateObjectDefinition(o), createObjectSortType(o), createObjectFilterType(o))
		definitions = append(definitions, createObjectRelationship(o))
		definitions = append(definitions, objectResultTypeDefinition(&o))
	}

	schemaHeaderNodes := []ast.Node{
		scalarDefinition("Time"),
		scalarDefinition("Upload"),
		scalarDefinition("_Any"),
		scalarDefinition("Any"),
		schemaDefinition(m),
		queryDefinition(m),
		mutationDefinition(m),
		createObjectSortEnum(),
	}
	m.Doc.Definitions = append(schemaHeaderNodes, m.Doc.Definitions...)
	m.Doc.Definitions = append(m.Doc.Definitions, definitions...)
	// m.Doc.Definitions = append(m.Doc.Definitions, createFederationServiceObject())

	return nil
}

func scalarDefinition(name string) *ast.ScalarDefinition {
	return &ast.ScalarDefinition{
		Name: &ast.Name{
			Kind:  kinds.Name,
			Value: name,
		},
		Kind: "ScalarDefinition",
	}
}

func columnDefinition(columnName, columnType, sourceType string, isNonNull bool) *ast.FieldDefinition {
	t := namedType(columnType)
	t1 := namedType(sourceType)

	if isNonNull {
		t = nonNull(t)
		t1 = nonNull(t1)
	}

	tt := getNamedType(t1).(*ast.Named)

	return columnDefinitionWithType(columnName, tt.Name.Value, t)
}
func columnDefinitionWithType(fieldName, kindName string, t ast.Type) *ast.FieldDefinition {
	return &ast.FieldDefinition{
		Name: nameNode(fieldName),
		Kind: kinds.FieldDefinition,
		Type: t,
		Description: &ast.StringValue{
			Kind: kindName,
		},
		Directives: []*ast.Directive{
			{
				Kind: kinds.Directive,
				Name: nameNode("column"),
			},
		},
	}
}

func schemaDefinition(m *Model) *ast.SchemaDefinition {
	return &ast.SchemaDefinition{
		Kind: kinds.SchemaDefinition,
		OperationTypes: []*ast.OperationTypeDefinition{
			{
				Operation: "query",
				Kind:      kinds.OperationTypeDefinition,
				Type: &ast.Named{
					Kind: kinds.Named,
					Name: &ast.Name{
						Kind:  kinds.Name,
						Value: "Query",
					},
				},
			},
			{
				Operation: "mutation",
				Kind:      kinds.OperationTypeDefinition,
				Type: &ast.Named{
					Kind: kinds.Named,
					Name: &ast.Name{
						Kind:  kinds.Name,
						Value: "Mutation",
					},
				},
			},
			{
				Operation: "subscription",
				Kind:      kinds.OperationTypeDefinition,
				Type: &ast.Named{
					Kind: kinds.Named,
					Name: &ast.Name{
						Kind:  kinds.Name,
						Value: "Subscription",
					},
				},
			},
		},
	}
}

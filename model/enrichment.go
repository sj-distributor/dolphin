package model

import (
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
)

// EnrichModelObjects ...
func EnrichModelObjects(m *Model) error {
	id := columnDefinition("id", "ID", true)

	createdAt := columnDefinition("createdAt", "Int", true)
	updatedAt := columnDefinition("updatedAt", "Int", false)
	deletedAt := columnDefinition("deletedAt", "Int", false)
	createdBy := columnDefinition("createdBy", "ID", false)
	updatedBy := columnDefinition("updatedBy", "ID", false)
	deletedBy := columnDefinition("deletedBy", "ID", false)

	for _, o := range m.ObjectEntities() {
		o.Def.Fields = append([]*ast.FieldDefinition{id}, o.Def.Fields...)
		for _, rel := range o.Relationships() {
			if rel.IsToOne() {
				o.Def.Fields = append(o.Def.Fields, columnDefinition(rel.Name()+"Id", "ID", false))
			}
		}
		o.Def.Fields = append(o.Def.Fields, deletedBy, updatedBy, createdBy, deletedAt, updatedAt, createdAt)
	}
	return nil
}

func columnDefinition(columnName, columnType string, isNonNull bool) *ast.FieldDefinition {
	t := namedType(columnType)
	if isNonNull {
		t = nonNull(t)
	}
	return columnDefinitionWithType(columnName, t)
}

func columnDefinitionWithType(fieldName string, t ast.Type) *ast.FieldDefinition {
	return &ast.FieldDefinition{
		Name: nameNode(fieldName),
		Kind: kinds.FieldDefinition,
		Type: t,
	}
}

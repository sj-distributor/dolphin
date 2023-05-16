package model

import (
	"strings"

	"github.com/graphql-go/graphql/language/kinds"

	"github.com/graphql-go/graphql/language/ast"
)

// objectDefinitionFunc ...
func objectDefinitionFunc(obj Object, name string) *ast.InputObjectDefinition {
	fields := []*ast.InputValueDefinition{}
	for _, col := range obj.Fields() {
		t := col.Def.Type
		if strings.Contains(name, CREATE) {
			if !col.IsCreatable() {
				continue
			}

			if col.Name() == "id" {
				t = getNamedType(t)
			}
		}

		if strings.Contains(name, UPDATE) {
			if !col.IsUpdatable() || col.Name() == "id" {
				continue
			}
		}

		if isListType(getNullableType(t)) {
			t = getNullableType(t)
		}

		if col.IsRelationship() {
			if col.IsListType() {
				t = namedType("[" + col.TargetObject().Name() + "Relationship" + "]")
			} else {
				t = namedType(col.TargetObject().Name() + "Relationship")
			}

			fields = append(fields, &ast.InputValueDefinition{
				Kind: kinds.InputValueDefinition,
				Name: nameNode(col.Name()),
				Type: t,
			})
		} else {
			if name == "UpdateInput" {
				t = getNullableType(t)
			}
			fields = append(fields, &ast.InputValueDefinition{
				Kind:        kinds.InputValueDefinition,
				Name:        col.Def.Name,
				Description: col.Def.Description,
				Type:        t,
			})
		}
	}
	return &ast.InputObjectDefinition{
		Kind:   kinds.InputObjectDefinition,
		Name:   nameNode(obj.Name() + name),
		Fields: fields,
	}
}

// objectRelationshipFunc ...
func objectRelationshipFunc(obj Object, name string) *ast.InputObjectDefinition {
	fields := []*ast.InputValueDefinition{}
	for _, col := range obj.Columns() {
		t := col.Def.Type

		if !col.IsUpdatable() || col.IsReadonlyType() || isListType(getNullableType(t)) {
			continue
		}

		if col.Name() == "id" {
			t = getNamedType(t)
		}
		// if strings.Contains(name, CREATE) {
		// 	if !col.IsCreatable() || col.IsReadonlyType() || col.Name() == "id" {
		// 		continue
		// 	}

		// 	if col.Name() == "id" {
		// 		t = getNamedType(t)
		// 	}
		// }

		// if strings.Contains(name, UPDATE) {
		// 	if !col.IsUpdatable() || col.IsReadonlyType() || col.Name() == "id" {
		// 		continue
		// 	}
		// }

		// if isListType(getNullableType(t)) {
		// 	t = getNullableType(t)
		// }

		fields = append(fields, &ast.InputValueDefinition{
			Kind:        kinds.InputValueDefinition,
			Name:        col.Def.Name,
			Description: col.Def.Description,
			Type:        t,
		})
	}
	return &ast.InputObjectDefinition{
		Kind:   kinds.InputObjectDefinition,
		Name:   nameNode(obj.Name() + name),
		Fields: fields,
	}
}

func createObjectDefinition(obj Object) *ast.InputObjectDefinition {
	return objectDefinitionFunc(obj, "CreateInput")
}

func updateObjectDefinition(obj Object) *ast.InputObjectDefinition {
	return objectDefinitionFunc(obj, "UpdateInput")
}

func createObjectRelationship(obj Object) *ast.InputObjectDefinition {
	return objectRelationshipFunc(obj, "Relationship")
}

func updateObjectRelationship(obj Object) *ast.InputObjectDefinition {
	return objectRelationshipFunc(obj, "UpdateRelationship")
}

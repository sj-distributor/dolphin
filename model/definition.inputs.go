package model

import (
	"github.com/graphql-go/graphql/language/kinds"

	"github.com/graphql-go/graphql/language/ast"
)

func createObjectDefinition(obj Object) *ast.InputObjectDefinition {
	fields := []*ast.InputValueDefinition{}
	for _, col := range obj.Columns() {
		if !col.IsCreatable() {
			continue
		}
		t := col.InputType()
		if col.IsIdentifier() {
			t = getNamedType(t)
		}
		fields = append(fields, &ast.InputValueDefinition{
			Kind:        kinds.InputValueDefinition,
			Name:        col.Def.Name,
			Description: col.Def.Description,
			Type:        t,
		})
	}
	return &ast.InputObjectDefinition{
		Kind:   kinds.InputObjectDefinition,
		Name:   nameNode(obj.Name() + "CreateInput"),
		Fields: fields,
	}
}

func updateObjectDefinition(obj Object) *ast.InputObjectDefinition {
	fields := []*ast.InputValueDefinition{}
	for _, col := range obj.Columns() {
		if !col.IsCreatable() || col.Name() == "id" {
			continue
		}
		fields = append(fields, &ast.InputValueDefinition{
			Kind:        kinds.InputValueDefinition,
			Name:        col.Def.Name,
			Description: col.Def.Description,
			Type:        col.InputType(),
		})
	}
	return &ast.InputObjectDefinition{
		Kind:   kinds.InputObjectDefinition,
		Name:   nameNode(obj.Name() + "UpdateInput"),
		Fields: fields,
	}
}

func createObjectRelationship(obj Object) *ast.InputObjectDefinition {
	fields := []*ast.InputValueDefinition{}
	for _, col := range obj.Fields() {
		if !col.IsCreatable() || col.Name() == "id" {
			continue
		}
		t := col.Def.Type
		if col.Name() == "id" {
			t = getNamedType(t)
		}
		if isListType(getNullableType(t)) {
			t = getNullableType(t)
		}

		if col.IsRelationship() {
			required := ""
			// if col.IsRequired() {
			// 	required = "!"
			// }

			// 判断是否一对一或其他
			if col.IsListType() {
				t = namedType("[" + col.TargetObject().Name() + "CreateReverseRelationship" + "]" + required)
			} else {
				t = namedType(col.TargetObject().Name() + "CreateReverseRelationship" + required)
			}

			field := ast.InputValueDefinition{
				Kind: kinds.InputValueDefinition,
				Name: nameNode(col.Name()),
				Type: t,
			}
			fields = append(fields, &field)
		} else {
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
		Name:   nameNode(obj.Name() + "CreateRelationship"),
		Fields: fields,
	}
}

func updateObjectRelationship(obj Object) *ast.InputObjectDefinition {
	fields := []*ast.InputValueDefinition{}
	for _, col := range obj.Fields() {
		if !col.IsUpdatable() {
			continue
		}
		t := col.Def.Type
		if col.Name() == "id" {
			t = getNamedType(t)
		}
		if isListType(getNullableType(t)) {
			t = getNullableType(t)
		}

		if col.IsRelationship() {
			// 判断是否一对一或其他
			t := col.Def.Type
			if col.IsListType() {
				t = namedType("[" + col.TargetObject().Name() + "UpdateReverseRelationship" + "]")
			} else {
				t = namedType(col.TargetObject().Name() + "UpdateReverseRelationship")
			}

			field := ast.InputValueDefinition{
				Kind: kinds.InputValueDefinition,
				Name: nameNode(col.Name()),
				Type: t,
			}
			fields = append(fields, &field)
		} else {
			fields = append(fields, &ast.InputValueDefinition{
				Kind:        kinds.InputValueDefinition,
				Name:        col.Def.Name,
				Description: col.Def.Description,
				Type:        getNullableType(col.Def.Type),
			})
		}
	}
	return &ast.InputObjectDefinition{
		Kind:   kinds.InputObjectDefinition,
		Name:   nameNode(obj.Name() + "UpdateRelationship"),
		Fields: fields,
	}
}

func createReverseRelationship(obj Object) *ast.InputObjectDefinition {
	fields := []*ast.InputValueDefinition{}
	for _, col := range obj.Columns() {
		if !col.IsCreatable() || col.IsReadonlyType() || col.Name() == "id" {
			continue
		}
		t := col.Def.Type
		if col.Name() == "id" {
			t = getNamedType(t)
		}
		if isListType(getNullableType(t)) {
			t = getNullableType(t)
		}

		fields = append(fields, &ast.InputValueDefinition{
			Kind:        kinds.InputValueDefinition,
			Name:        col.Def.Name,
			Description: col.Def.Description,
			Type:        t,
		})
	}
	return &ast.InputObjectDefinition{
		Kind:   kinds.InputObjectDefinition,
		Name:   nameNode(obj.Name() + "CreateReverseRelationship"),
		Fields: fields,
	}
}

func updateReverseRelationship(obj Object) *ast.InputObjectDefinition {
	fields := []*ast.InputValueDefinition{}
	for _, col := range obj.Columns() {
		if !col.IsUpdatable() || col.IsReadonlyType() {
			continue
		}
		t := col.Def.Type
		if col.Name() == "id" {
			t = getNamedType(t)
		}
		if isListType(getNullableType(t)) {
			t = getNullableType(t)
		}

		fields = append(fields, &ast.InputValueDefinition{
			Kind:        kinds.InputValueDefinition,
			Name:        col.Def.Name,
			Description: col.Def.Description,
			Type:        t,
		})
	}
	return &ast.InputObjectDefinition{
		Kind:   kinds.InputObjectDefinition,
		Name:   nameNode(obj.Name() + "UpdateReverseRelationship"),
		Fields: fields,
	}
}

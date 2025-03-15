package model

import (
	"github.com/graphql-go/graphql/language/kinds"

	"github.com/graphql-go/graphql/language/ast"
)

func createObjectSortEnum() *ast.EnumDefinition {
	values := []*ast.EnumValueDefinition{
		{
			Kind: kinds.EnumValueDefinition,
			Name: nameNode("ASC"),
		},
		{
			Kind: kinds.EnumValueDefinition,
			Name: nameNode("DESC"),
		},
	}
	return &ast.EnumDefinition{
		Kind:   kinds.EnumDefinition,
		Name:   nameNode("ObjectSortType"),
		Values: values,
	}
}

func createObjectDirective(obj Object) []*ast.Directive {
	directives := []*ast.Directive{}

	if obj.HasDirective("hasRole") {
		directive := obj.Directive("hasRole")
		arguments := directive.Arguments
		if len(arguments) > 0 {
			directives = append(directives, &ast.Directive{
				Kind:      kinds.Directive,
				Name:      nameNode(directive.Name.Value),
				Arguments: arguments,
			})
		}
	}

	return directives
}

func createObjectColumnDirective(col ObjectField) []*ast.Directive {
	directives := []*ast.Directive{}

	if col.HasDirective("validator") {
		directive := col.Directive("validator")
		arguments := directive.Arguments

		if len(arguments) > 0 {
			directives = append(directives, &ast.Directive{
				Kind:      kinds.Directive,
				Name:      nameNode(directive.Name.Value),
				Arguments: arguments,
			})
		}
	}

	return directives
}

func createObjectSortType(obj Object) *ast.InputObjectDefinition {
	fields := []*ast.InputValueDefinition{}

	for _, col := range obj.Columns() {
		if col.IsReadonlyType() {
			continue
		}

		field := ast.InputValueDefinition{
			Kind: kinds.InputValueDefinition,
			Name: nameNode(col.Name()),
			Type: namedType("ObjectSortType"),
		}
		fields = append(fields, &field)
	}

	for _, rel := range obj.Relationships() {
		field := ast.InputValueDefinition{
			Kind: kinds.InputValueDefinition,
			Name: nameNode(rel.Name()),
			Type: namedType(rel.Target().Name() + "SortType"),
		}
		fields = append(fields, &field)
	}

	return &ast.InputObjectDefinition{
		Kind:   kinds.InputObjectDefinition,
		Name:   nameNode(obj.Name() + "SortType"),
		Fields: fields,
	}
}

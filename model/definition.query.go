package model

import (
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

func queryDefinition(m *Model) *ast.ObjectDefinition {
	fields := []*ast.FieldDefinition{}

	if m.HasFederatedTypes() {
		fields = append(fields, createFederationEntitiesQueryField())
	}

	for _, obj := range m.ObjectEntities() {
		if obj.IsSkip() {
			continue
		}

		fields = append(fields, fetchFieldDefinition(obj), listFieldDefinition(obj))
	}
	fields = append(fields, rootExtensionFields(m, "Query", fields)...)
	return &ast.ObjectDefinition{
		Kind: kinds.ObjectDefinition,
		Name: &ast.Name{
			Kind:  kinds.Name,
			Value: "Query",
		},
		Fields: fields,
	}
}

// rootExtensionFields returns project-owned named operations for a generated root.
func rootExtensionFields(m *Model, rootName string, generatedFields []*ast.FieldDefinition) []*ast.FieldDefinition {
	fields := []*ast.FieldDefinition{}
	seen := make(map[string]struct{}, len(generatedFields))
	for _, field := range generatedFields {
		seen[field.Name.Value] = struct{}{}
	}
	for _, extension := range m.ObjectExtensions() {
		if extension.Object.Name() == rootName {
			for _, field := range extension.Object.Def.Fields {
				if _, duplicate := seen[field.Name.Value]; duplicate {
					continue
				}
				seen[field.Name.Value] = struct{}{}
				fields = append(fields, field)
			}
		}
	}
	return fields
}

func fetchFieldDefinition(obj Object) *ast.FieldDefinition {
	return &ast.FieldDefinition{
		Kind:       kinds.FieldDefinition,
		Name:       nameNode(inflection.Singular(strcase.ToLowerCamel(obj.Name()))),
		Type:       namedType(obj.Name()),
		Directives: createObjectDirective(obj),
		Arguments: []*ast.InputValueDefinition{
			{
				Kind: kinds.InputValueDefinition,
				Name: nameNode("id"),
				Type: namedType("ID"),
			},
			{
				Kind: kinds.InputValueDefinition,
				Name: nameNode("filter"),
				Type: namedType(obj.Name() + "FilterType"),
			},
		},
	}
}

func listFieldDefinition(obj Object) *ast.FieldDefinition {
	// createObjectSortType(obj)
	return &ast.FieldDefinition{
		Kind:       kinds.FieldDefinition,
		Name:       nameNode(inflection.Plural(strcase.ToLowerCamel(obj.Name()))),
		Type:       namedType(obj.Name() + "ResultType"),
		Directives: createObjectDirective(obj),
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

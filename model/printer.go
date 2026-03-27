package model

import (
	"fmt"

	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/printer"
)

func filterDirective(ds []*ast.Directive, name string) []*ast.Directive {
	res := []*ast.Directive{}
	for _, d := range ds {
		if d.Name.Value != name {
			res = append(res, d)
		}
	}
	return res
}

// deduplicateDefinitions removes duplicate DirectiveDefinition and EnumDefinition
// from the AST document to prevent "Cannot redeclare type" errors.
// It also filters out fields in TypeExtensionDefinition that already exist in the base type.
func deduplicateDefinitions(doc *ast.Document) {
	seenDirectives := make(map[string]bool)
	seenEnums := make(map[string]bool)

	// First pass: collect all fields from ObjectDefinitions
	objectFields := make(map[string]map[string]bool) // typeName -> fieldName -> exists
	for _, def := range doc.Definitions {
		if objDef, ok := def.(*ast.ObjectDefinition); ok {
			fieldMap := make(map[string]bool)
			for _, field := range objDef.Fields {
				fieldMap[field.Name.Value] = true
			}
			objectFields[objDef.Name.Value] = fieldMap
		}
	}

	newDefs := []ast.Node{}
	for _, def := range doc.Definitions {
		switch d := def.(type) {
		case *ast.DirectiveDefinition:
			if !seenDirectives[d.Name.Value] {
				seenDirectives[d.Name.Value] = true
				newDefs = append(newDefs, def)
			}
		case *ast.EnumDefinition:
			if !seenEnums[d.Name.Value] {
				seenEnums[d.Name.Value] = true
				newDefs = append(newDefs, def)
			}
		case *ast.TypeExtensionDefinition:
			// Filter out fields that already exist in the base type
			baseTypeName := d.Definition.Name.Value
			if baseFields, exists := objectFields[baseTypeName]; exists {
				filteredFields := []*ast.FieldDefinition{}
				for _, field := range d.Definition.Fields {
					if !baseFields[field.Name.Value] {
						filteredFields = append(filteredFields, field)
					}
				}
				// Only add extension if it has remaining fields
				if len(filteredFields) > 0 {
					d.Definition.Fields = filteredFields
					newDefs = append(newDefs, def)
				}
			} else {
				newDefs = append(newDefs, def)
			}
		default:
			newDefs = append(newDefs, def)
		}
	}
	doc.Definitions = newDefs
}

// PrintSchema
func PrintSchema(model Model) (string, error) {
	for _, o := range model.Objects() {
		fields := []*ast.FieldDefinition{}
		for _, f := range o.Def.Fields {
			f.Directives = filterDirective(f.Directives, "relationship")
			f.Directives = filterDirective(f.Directives, "column")
			f.Directives = filterDirective(f.Directives, "validator")
			f.Directives = filterDirective(f.Directives, "hasPermission")
			fields = append(fields, f)
		}
		o.Def.Fields = fields
		o.Def.Directives = filterDirective(o.Def.Directives, "skip")
		o.Def.Directives = filterDirective(o.Def.Directives, "entity")
		o.Def.Directives = filterDirective(o.Def.Directives, "hasRole")
		o.Def.Directives = filterDirective(o.Def.Directives, "sharding")
	}

	for _, o := range model.ObjectExtensions() {
		fields := []*ast.FieldDefinition{}
		for _, f := range o.Object.Def.Fields {
			f.Directives = filterDirective(f.Directives, "entity")
			f.Directives = filterDirective(f.Directives, "hasRole")
			f.Directives = filterDirective(f.Directives, "sharding")
			fields = append(fields, f)
		}
		o.Object.Def.Fields = fields
	}

	for _, o := range model.InputObjects() {
		fields := []*ast.InputValueDefinition{}
		for _, f := range o.Def.Fields {
			f.Directives = filterDirective(f.Directives, "relationship")
			f.Directives = filterDirective(f.Directives, "column")
			f.Directives = filterDirective(f.Directives, "validator")
			f.Directives = filterDirective(f.Directives, "hasPermission")
			f.Directives = filterDirective(f.Directives, "entity")
			fields = append(fields, f)
		}
		o.Def.Fields = fields
		o.Def.Directives = filterDirective(o.Def.Directives, "skip")
		o.Def.Directives = filterDirective(o.Def.Directives, "entity")
		o.Def.Directives = filterDirective(o.Def.Directives, "hasRole")
		o.Def.Directives = filterDirective(o.Def.Directives, "sharding")
	}

	// Deduplicate directive and enum definitions to prevent redeclaration errors
	deduplicateDefinitions(model.Doc)

	printed := printer.Print(model.Doc)
	printedString, ok := printed.(string)
	if !ok {
		return "", fmt.Errorf("printer.Print returned unexpected type %T", printed)
	}
	return printedString, nil
}

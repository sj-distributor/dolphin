package model

import (
	"fmt"

	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/printer"
)

func cleanDirectives(ds []*ast.Directive) []*ast.Directive {
	res := []*ast.Directive{}
	skipList := []string{"relationship", "column", "validator", "skip", "entity", "hasRole", "sharding"}
	for _, d := range ds {
		isInternal := false
		for _, name := range skipList {
			if d.Name.Value == name {
				isInternal = true
				break
			}
		}
		if !isInternal {
			res = append(res, d)
		}
	}
	return res
}

func cleanNode(node ast.Node) {
	if node == nil {
		return
	}
	switch n := node.(type) {
	case *ast.ObjectDefinition:
		n.Directives = cleanDirectives(n.Directives)
		for _, f := range n.Fields {
			cleanNode(f)
		}
	case *ast.InputObjectDefinition:
		n.Directives = cleanDirectives(n.Directives)
		for _, f := range n.Fields {
			cleanNode(f)
		}
	case *ast.InterfaceDefinition:
		n.Directives = cleanDirectives(n.Directives)
		for _, f := range n.Fields {
			cleanNode(f)
		}
	case *ast.FieldDefinition:
		n.Directives = cleanDirectives(n.Directives)
		for _, arg := range n.Arguments {
			cleanNode(arg)
		}
	case *ast.InputValueDefinition:
		n.Directives = cleanDirectives(n.Directives)
	case *ast.TypeExtensionDefinition:
		cleanNode(n.Definition)
	case *ast.EnumDefinition:
		n.Directives = cleanDirectives(n.Directives)
		for _, v := range n.Values {
			cleanNode(v)
		}
	case *ast.EnumValueDefinition:
		n.Directives = cleanDirectives(n.Directives)
	}
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
	for _, def := range model.Doc.Definitions {
		cleanNode(def)
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

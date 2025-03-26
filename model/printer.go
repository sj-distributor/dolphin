package model

import (
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

// PrintSchema
func PrintSchema(model Model) (string, error) {
	for _, o := range model.Objects() {
		fields := []*ast.FieldDefinition{}
		for _, f := range o.Def.Fields {
			f.Directives = filterDirective(f.Directives, "relationship")
			f.Directives = filterDirective(f.Directives, "column")
			f.Directives = filterDirective(f.Directives, "validator")
			fields = append(fields, f)
		}
		o.Def.Fields = fields
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

	printed := printer.Print(model.Doc)
	printedString, _ := printed.(string)

	return printedString, nil
}

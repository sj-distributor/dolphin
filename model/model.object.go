package model

import "github.com/graphql-go/graphql/language/ast"

type Object struct {
	Def       *ast.ObjectDefinition
	Model     *Model
	Extension *ObjectExtension
}

// Name ...
func (o *Object) Name() string {
	return o.Def.Name.Value
}

// Columns ...
func (o *Object) Columns() []ObjectField {
	columns := []ObjectField{}
	for _, f := range o.Fields() {
		// if f.IsColumn() {
		columns = append(columns, f)
		// }
	}
	return columns
}

// Fields ...
func (o *Object) Fields() []ObjectField {
	fields := []ObjectField{}
	for _, f := range o.Def.Fields {
		fields = append(fields, ObjectField{f, o})
	}
	return fields
}

func (o *Object) Directive(name string) *ast.Directive {
	for _, d := range o.Def.Directives {
		if d.Name.Value == name {
			return d
		}
	}
	return nil
}

func (o *Object) isRelationship(f *ast.FieldDefinition) bool {
	for _, d := range f.Directives {
		if d != nil && d.Name.Value == "relationship" {
			return true
		}
	}
	return false
}

func (o *Object) Relationships() []*ObjectRelationship {
	relationships := []*ObjectRelationship{}
	for _, f := range o.Def.Fields {
		if o.isRelationship(f) {
			relationships = append(relationships, &ObjectRelationship{f, o})
		}
	}
	return relationships
}

func (o *Object) HasDirective(name string) bool {
	return o.Directive(name) != nil
}

func (o *Object) IsFederatedType() bool {
	return o.HasDirective("key")
}

func (o *Object) Interfaces() []string {
	interfaces := []string{}
	for _, item := range o.Def.Interfaces {
		interfaces = append(interfaces, item.Name.Value)
	}
	return interfaces
}

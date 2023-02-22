package model

import (
	"fmt"
	"strings"

	"github.com/graphql-go/graphql/language/ast"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

type Object struct {
	Def       *ast.ObjectDefinition
	Model     *Model
	Extension *ObjectExtension
}

// Name ...
func (o *Object) Name() string {
	return o.Def.Name.Value
}

func (o *Object) PluralName() string {
	return inflection.Plural(o.Name())
}
func (o *Object) ToLowerPluralName() string {
	return strcase.ToLowerCamel(o.PluralName())
}
func (o *Object) LowerName() string {
	return strcase.ToLowerCamel(o.Name())
}
func (o *Object) ToSnakeName() string {
	return strcase.ToSnake(o.Name())
}
func (o *Object) TableName() string {
	return strcase.ToSnake(inflection.Plural(o.LowerName()))
}

// Columns ...
func (o *Object) Columns() []ObjectField {
	columns := []ObjectField{}
	for _, f := range o.Fields() {
		// if f.IsColumn() {
		if !f.IsRelationship() {
			columns = append(columns, f)
		}
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

func (o *Object) HasEmbeddedField() bool {
	for _, f := range o.Fields() {
		if f.IsEmbedded() {
			return true
		}
	}
	return false
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

func (o *Object) IsToManyColumn(c ObjectField) bool {
	if c.Obj.Name() != o.Name() {
		return false
	}
	return o.HasRelationship(strings.TrimSuffix(c.Name(), "Ids"))
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

func (o *Object) Relationship(name string) *ObjectRelationship {
	for _, rel := range o.Relationships() {
		if rel.Name() == name {
			return rel
		}
	}
	panic(fmt.Sprintf("relationship %s->%s not found", o.Name(), name))
}

func (o *Object) HasRelationship(name string) bool {
	for _, rel := range o.Relationships() {
		if rel.Name() == name {
			return true
		}
	}
	return false
}
func (o *Object) HasAnyRelationships() bool {
	return len(o.Relationships()) > 0
}
func (o *Object) NeedsQueryResolver() bool {
	return o.HasAnyRelationships() || o.HasEmbeddedField() || o.Model.HasObjectExtension(o.Name())
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

func (o *Object) IsExtended() bool {
	return o.Extension != nil
}

package model

import (
	"fmt"

	"github.com/graphql-go/graphql/language/ast"
)

type Model struct {
	Doc *ast.Document
	// Objects []Object
}

var defaultScalars map[string]bool = map[string]bool{
	"Int":     true,
	"Float":   true,
	"String":  true,
	"Boolean": true,
	"ID":      true,
	"Any":     true,
	"Time":    true,
}

func (m *Model) Objects() []Object {
	objs := []Object{}
	for _, def := range m.Doc.Definitions {
		def, ok := def.(*ast.ObjectDefinition)
		if ok {
			// for _, child := range def.Fields {
			// 	nt := getNamedType(child.Type).(*ast.Named)
			// 	fmt.Println(nt.Name.Value)
			// }
			// if len(def.Directives[0].Arguments) > 0 {
			// 	fmt.Println(def.Directives[0].Arguments[0].Value.GetValue())
			// }
			objs = append(objs, Object{Def: def, Model: m})
		}
	}
	return objs
}

func (m *Model) HasObject(name string) bool {
	if name == "Query" || name == "Mutation" || name == "Subscription" {
		return true
	}
	for _, o := range m.Objects() {
		if o.Name() == name {
			return true
		}
	}
	return false
}

func (m *Model) ObjectEntities() []Object {
	objs := []Object{}
	for _, def := range m.Doc.Definitions {
		def, ok := def.(*ast.ObjectDefinition)
		if ok {
			obj := Object{Def: def, Model: m}
			if obj.HasDirective("entity") {
				objs = append(objs, obj)
			}
		}
	}
	return objs
}

func (m *Model) HasFederatedTypes() bool {
	for _, o := range m.Objects() {
		if o.IsFederatedType() {
			return true
		}
	}

	return false
}

func (m *Model) ObjectExtensions() []ObjectExtension {
	objs := []ObjectExtension{}
	for _, def := range m.Doc.Definitions {
		def, ok := def.(*ast.TypeExtensionDefinition)
		if ok {
			obj := &Object{Def: def.Definition, Model: m}
			objs = append(objs, ObjectExtension{Def: def, Model: m, Object: obj})
		}
	}
	return objs
}

func (m *Model) Object(name string) Object {
	for _, o := range m.Objects() {
		if o.Name() == name {
			return o
		}
	}
	panic(fmt.Sprintf("Object with name %s not found in model", name))
}

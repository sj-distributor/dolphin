package model

import (
	"github.com/graphql-go/graphql/language/ast"
)

type Model struct {
	Doc *ast.Document
	// Objects []Object
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

func (m *Model) ObjectExtensions() []ObjectExtension {
	objs := []ObjectExtension{}
	for _, def := range m.Doc.Definitions {
		def, ok := def.(*ast.TypeExtensionDefinition)
		if ok {
			obj := &Object{Def: def.Definition, Model: m}
			// obj.Def.Directives = filterDirective(obj.Def.Directives, "entity")
			// fmt.Println(def.Directives[0].Arguments[0].Value.GetValue())
			// fmt.Println(def.Definition.Fields[0].Name)
			// fmt.Println(obj.Def.Fields[0].Name.Value)
			// for _, child := range obj.Def.Fields[0].Arguments {
			// }
			objs = append(objs, ObjectExtension{Def: def, Model: m, Object: obj})
		}
	}
	return objs
}

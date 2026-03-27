package model

import (
	"log"

	"github.com/graphql-go/graphql/language/ast"
)

type Model struct {
	Doc       *ast.Document
	objects   []Object
	objectMap map[string]Object
}

type TypeData struct {
	Name   string          `json:"name"`
	Fields []FieldMetadata `json:"fields"`
}

type FieldMetadata struct {
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	Type      string `json:"type"`
	Required  string `json:"required"`
	Validator string `json:"validator"`
}

func (m *Model) SecretKey() string {
	return GetRandomString(32)
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

func (m *Model) loadObjects() {
	if m.objects != nil {
		return
	}
	objs := []Object{}
	objMap := make(map[string]Object)
	for _, def := range m.Doc.Definitions {
		def, ok := def.(*ast.ObjectDefinition)
		if ok {
			obj := Object{Def: def, Model: m}
			objs = append(objs, obj)
			objMap[obj.Name()] = obj
		}
	}
	m.objects = objs
	m.objectMap = objMap
}

func (m *Model) DocObjects() []Object {
	m.loadObjects()
	return Filter(m.objects, func(o Object) bool {
		return !o.IsSkip()
	})
}

func (m *Model) Objects() []Object {
	m.loadObjects()
	return m.objects
}

func (m *Model) HasObject(name string) bool {
	if name == "Query" || name == "Mutation" || name == "Subscription" {
		return true
	}
	m.loadObjects()
	_, ok := m.objectMap[name]
	return ok
}

func (m *Model) ObjectEntities() []Object {
	return Filter(m.Objects(), func(o Object) bool {
		return o.HasDirective("entity")
	})
}

func (m *Model) ObjectShardings() []Object {
	return Filter(m.Objects(), func(o Object) bool {
		return o.HasDirective("sharding")
	})
}

func (m *Model) HasFederatedTypes() bool {
	return Any(m.Objects(), func(o Object) bool {
		return o.IsFederatedType()
	})
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
	m.loadObjects()
	if o, ok := m.objectMap[name]; ok {
		return o
	}
	log.Fatalf("Object with name %s not found in model", name)
	return Object{}
}

func (m *Model) ObjectExtension(name string) ObjectExtension {
	for _, e := range m.ObjectExtensions() {
		if e.Object.Name() == name {
			return e
		}
	}
	log.Fatalf("Extension for object with name %s not found in model", name)
	return ObjectExtension{}
}

func (m *Model) HasObjectExtension(name string) bool {
	for _, e := range m.ObjectExtensions() {
		if e.Object.Name() == name {
			return true
		}
	}
	return false
}

func (m *Model) GetDefinition(name string) ast.Node {
	for _, def := range m.Doc.Definitions {
		switch d := def.(type) {
		case *ast.ObjectDefinition:
			if d.Name.Value == name {
				return d
			}
		case *ast.InputObjectDefinition:
			if d.Name.Value == name {
				return d
			}
		case *ast.EnumDefinition:
			if d.Name.Value == name {
				return d
			}
		case *ast.ScalarDefinition:
			if d.Name.Value == name {
				return d
			}
		case *ast.InterfaceDefinition:
			if d.Name.Value == name {
				return d
			}
		case *ast.UnionDefinition:
			if d.Name.Value == name {
				return d
			}
		}
	}
	return nil
}

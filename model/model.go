package model

import (
	"fmt"

	"github.com/graphql-go/graphql/language/ast"
)

type Model struct {
	Doc       *ast.Document
	objects   []Object
	objectMap map[string]Object
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
	objs := []Object{}
	for _, obj := range m.Objects() {
		if obj.HasDirective("entity") {
			objs = append(objs, obj)
		}
	}
	return objs
}

func (m *Model) ObjectShardings() []Object {
	objs := []Object{}
	for _, obj := range m.Objects() {
		if obj.HasDirective("sharding") {
			objs = append(objs, obj)
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
	m.loadObjects()
	if o, ok := m.objectMap[name]; ok {
		return o
	}
	panic(fmt.Sprintf("Object with name %s not found in model", name))
}

func (m *Model) ObjectExtension(name string) ObjectExtension {
	for _, e := range m.ObjectExtensions() {
		if e.Object.Name() == name {
			return e
		}
	}
	panic(fmt.Sprintf("Extension for object with name %s not found in model", name))
}

func (m *Model) HasObjectExtension(name string) bool {
	for _, e := range m.ObjectExtensions() {
		if e.Object.Name() == name {
			return true
		}
	}
	return false
}

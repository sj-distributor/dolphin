package model

import (
	"fmt"
	"strings"

	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/iancoleman/strcase"
)

var goTypeMap = map[string]string{
	"String":  "string",
	"Time":    "time.Time",
	"ID":      "string",
	"Float":   "float64",
	"Int":     "int64",
	"Boolean": "bool",
	"Any":     "interface{}",
}

type ObjectField struct {
	Def *ast.FieldDefinition
	Obj *Object
}

func (o *ObjectField) Name() string {
	return o.Def.Name.Value
}

func (o *ObjectField) LowerName() string {
	return strcase.ToLowerCamel(o.Name())
}

func (o *ObjectField) ToSnakeName() string {
	return strcase.ToSnake(o.Name())
}

func (o *ObjectField) MethodName() string {
	name := o.Name()
	return templates.ToGo(name)
}

func (o *ObjectField) RelationshipName() string {
	return strings.Replace(o.MethodName(), "ID", "", -1)
}

func (o *ObjectField) ArgumentsValue() []ObjectFieldInput {
	arguments := []ObjectFieldInput{}
	for _, f := range o.Def.Arguments {
		arguments = append(arguments, ObjectFieldInput{f, o})
	}
	return arguments
}

// TargetType ...
func (o *ObjectField) TargetType() string {
	nt := getNamedType(o.Def.Type).(*ast.Named)
	return nt.Name.Value
}

func (o *ObjectField) IsColumn() bool {
	return o.HasDirective("column")
}

func (o *ObjectField) IsListType() bool {
	return isListType(getNullableType(o.Def.Type))
}

func (o *ObjectField) TargetObject() *Object {
	obj := o.Obj.Model.Object(o.TargetType())
	return &obj
}

func (o *ObjectField) HasTargetObjectExtension() bool {
	return o.Obj.Model.HasObjectExtension(o.TargetType())
}

func (o *ObjectField) TargetObjectExtension() *ObjectExtension {
	e := o.Obj.Model.ObjectExtension(o.TargetType())
	return &e
}

func (o *ObjectField) Directive(name string) *ast.Directive {
	for _, d := range o.Def.Directives {
		if d.Name.Value == name {
			return d
		}
	}
	return nil
}

func (o *ObjectField) IsRelationship() bool {
	return o.HasDirective("relationship")
}

func (o *ObjectField) NeedsQueryResolver() bool {
	return o.IsEmbedded()
}

func (o *ObjectField) HasTargetTypeWithIDField() bool {
	if o.HasTargetObject() && o.TargetObject().HasField("id") {
		return true
	}
	if o.HasTargetObjectExtension() && o.TargetObjectExtension().Object.HasField("id") {
		return true
	}
	return false
}

func (o *ObjectField) IsHasUpperId() bool {
	return strings.Contains(o.Name(), "Id") && o.IsRelationship()
}

// IsIdentifier ...
func (o *ObjectField) IsIdentifier() bool {
	return o.HasName("id")
}

func (o *ObjectField) HasName(name string) bool {
	return o.Name() == name
}

// IsRelationshipIdentifier ...
func (o *ObjectField) IsRelationshipIdentifier() bool {
	return strings.HasSuffix(o.Name(), "Id") || strings.HasSuffix(o.Name(), "Ids")
}

// IsCreatable ...
func (o *ObjectField) IsCreatable() bool {
	return !(o.Name() == "id" || o.Name() == "createdAt" || o.Name() == "updatedAt" || o.Name() == "deletedAt" || o.Name() == "createdBy" || o.Name() == "updatedBy" || o.Name() == "deletedBy")
}

func (o *ObjectField) IsUpdatable() bool {
	return !(o.Name() == "createdAt" || o.Name() == "updatedAt" || o.Name() == "deletedAt" || o.Name() == "createdBy" || o.Name() == "updatedBy" || o.Name() == "deletedBy")
}

// IsReadonlyType ..
func (o *ObjectField) IsReadonlyType() bool {
	if o.IsEmbeddedColumn() {
		return false
	}
	return !(o.IsScalarType() || o.IsEnumType()) || o.Obj.Model.HasObject(o.TargetType())
}

// IsEnumType ...
func (o *ObjectField) IsEnumType() bool {
	return o.Obj.Model.HasEnum(o.TargetType())
}

func (o *ObjectField) IsWritableType() bool {
	return !o.IsReadonlyType()
}

// IsScalarType ...
func (o *ObjectField) IsScalarType() bool {
	return o.Obj.Model.HasScalar(o.TargetType())
}

// IsOptional ...
func (o *ObjectField) IsOptional() bool {
	return !isNonNullType(o.Def.Type)
}

// IsList ...
func (o *ObjectField) IsList() bool {
	return isListType(o.Def.Type)
}

// IsEmbedded ...
func (o *ObjectField) IsEmbedded() bool {
	return !o.IsColumn() && !o.IsRelationship()
}

// HasTargetObject ...
func (o *ObjectField) HasTargetObject() bool {
	return o.Obj.Model.HasObject(o.TargetType())
}

// IsEmbeddedColumn ...
func (o *ObjectField) IsEmbeddedColumn() bool {
	return (o.IsColumn() && o.ColumnType() == "embedded")
}

func (o *ObjectField) IsSearchable() bool {
	t := getNamedType(o.Def.Type).(*ast.Named)
	return t.Name.Value == "String" || t.Name.Value == "Int" || t.Name.Value == "Float"
}

func (o *ObjectField) IsSortable() bool {
	return !o.IsReadonlyType() && o.IsScalarType()
}

func (o *ObjectField) IsID() bool {
	t := getNamedType(o.Def.Type).(*ast.Named)
	return t.Name.Value == "ID"
}

func (o *ObjectField) IsInt() bool {
	t := getNamedType(o.Def.Type).(*ast.Named)
	return t.Name.Value == "Int"
}

func (o *ObjectField) IsString() bool {
	t := getNamedType(o.Def.Type).(*ast.Named)
	return t.Name.Value == "String"
}

func (o *ObjectField) IsRequired() bool {
	return isNonNullType(o.Def.Type)
}

func (m *Model) HasScalar(name string) bool {
	if _, ok := defaultScalars[name]; ok {
		return true
	}
	for _, def := range m.Doc.Definitions {
		scalar, ok := def.(*ast.ScalarDefinition)
		if ok && scalar.Name.Value == name {
			return true
		}
	}
	return false
}

func (m *Model) HasEnum(name string) bool {
	if _, ok := defaultScalars[name]; ok {
		return true
	}
	for _, def := range m.Doc.Definitions {
		e, ok := def.(*ast.EnumDefinition)
		if ok && e.Name.Value == name {
			return true
		}
	}
	return false
}

// InputType ...
func (o *ObjectField) InputType() ast.Type {
	t := o.Def.Type
	if o.IsIdentifier() {
		t = nonNull(getNamedType(t))
	}
	isList := o.IsList()
	isOptional := o.IsOptional()

	if o.IsEmbeddedColumn() {
		_t := getNamedType(t).(*ast.Named)
		t = namedType(_t.Name.Value + "Input")

		if isList {
			t = listType(t)
		}
		if !isOptional {
			t = nonNull(t)
		}
	}
	if o.IsRelationshipIdentifier() {
		t = getNullableType(t)
	}

	return t
}

func (o *ObjectField) GoType() string {
	return o.GoTypeWithPointer()
}

func (o *ObjectField) GoTypeWithPointer() string {
	t := o.Def.Type
	st := ""

	if o.IsOptional() {
		st += "*"
	} else {
		t = getNullableType(t)
	}

	if isListType(t) {
		if o.IsRequired() {
			st = "[]"
		} else {
			st += "[]*"
		}
	}

	v, ok := getNamedType(o.Def.Type).(*ast.Named)
	if ok {
		_t, known := goTypeMap[v.Name.Value]
		if known {
			st += _t
		} else {
			st += v.Name.Value
		}
	}

	return st
}

func (o *ObjectField) ModelTags() string {
	_gorm := "default:null"

	if o.IsString() {
		_gorm = fmt.Sprintf("type:varchar(255) comment '%s';default:null;", o.ToSnakeName())
	} else if o.IsID() {
		_gorm = fmt.Sprintf("type:varchar(36) comment '%s';default:null;", o.ToSnakeName())
	} else if o.IsInt() {
		_gorm = fmt.Sprintf("type:bigint(13) comment '%s';default:null;", o.ToSnakeName())
	}

	_valid := ""

	if o.Name() == "createdBy" || o.Name() == "updatedBy" || o.Name() == "deletedBy" {
		_gorm += fmt.Sprintf("index:%s;", o.ToSnakeName())
	}

	if o.Name() == "createdAt" {
		_gorm += " autoCreateTime:milli;"
	}

	if o.Name() == "updatedAt" {
		_gorm += " autoUpdateTime:milli;"
	}

	if o.Name() == "id" {
		_gorm = "type:varchar(36) comment 'uuid';primary_key;unique_index;NOT NULL;"
	}

	for _, d := range o.Def.Directives {
		if d.Name.Value == "column" {
			for _, arg := range d.Arguments {
				if arg.Name.Value == "gorm" {
					_gorm = fmt.Sprintf("%v", arg.Value.GetValue())
				}
			}
		} else if d.Name.Value == "validator" {
			for _, arg := range d.Arguments {
				if arg.Value.GetValue() != nil {
					_valid += fmt.Sprintf("%v", arg.Name.Value+":"+arg.Value.GetValue().(string)+";")
				}
			}
		}
	}

	str := fmt.Sprintf(`json:"%s" gorm:"%s"`, o.Name(), _gorm)

	if _valid != "" {
		str = fmt.Sprintf(`json:"%s" gorm:"%s" validator:"%s"`, o.Name(), _gorm, _valid)
	}

	return str
}

func (f *FilterMappingItem) IsLike() bool {
	return f.SuffixCamel() == "Like"
}

func (f *FilterMappingItem) SuffixCamel() string {
	return strcase.ToCamel(f.Suffix)
}

func (f *FilterMappingItem) WrapValueVariable(v string) string {
	return fmt.Sprintf(f.ValueFormat, v)
}

type FilterMappingItem struct {
	Suffix      string
	Operator    string
	InputType   ast.Type
	ValueFormat string
}

func (o *ObjectField) FilterMapping() []FilterMappingItem {
	t := getNamedType(o.Def.Type)
	mapping := []FilterMappingItem{
		{"", "= ?", t, "%s"},
		{"_ne", "!= ?", t, "%s"},
		{"_gt", "> ?", t, "%s"},
		{"_lt", "< ?", t, "%s"},
		{"_gte", ">= ?", t, "%s"},
		{"_lte", "<= ?", t, "%s"},
		{"_in", "IN (?)", listType(nonNull(t)), "%s"},
	}
	_t := t.(*ast.Named)
	if _t.Name.Value == "String" {
		mapping = append(mapping,
			FilterMappingItem{"_like", "LIKE ?", t, "strings.Replace(strings.Replace(*%s,\"?\",\"_\",-1),\"*\",\"%%\",-1)"},
			FilterMappingItem{"_prefix", "LIKE ?", t, "fmt.Sprintf(\"%%s%%%%\",*%s)"},
			FilterMappingItem{"_suffix", "LIKE ?", t, "fmt.Sprintf(\"%%%%%%s\",*%s)"},
		)
	}
	return mapping
}

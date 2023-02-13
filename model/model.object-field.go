package model

import (
	"fmt"
	"strings"

	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/graphql-go/graphql/language/ast"
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

func (o *ObjectField) MethodName() string {
	name := o.Name()
	return templates.ToGo(name)
}

// TargetType ...
func (o *ObjectField) TargetType() string {
	nt := getNamedType(o.Def.Type).(*ast.Named)
	return nt.Name.Value
}

func (o *ObjectField) IsColumn() bool {
	return o.HasDirective("column")
}

func (o *ObjectField) Directive(name string) *ast.Directive {
	for _, d := range o.Def.Directives {
		if d.Name.Value == name {
			return d
		}
	}
	return nil
}

// IsIdentifier ...
func (o *ObjectField) IsIdentifier() bool {
	return o.Name() == "id"
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

// IsEmbeddedColumn ...
func (o *ObjectField) IsEmbeddedColumn() bool {
	return (o.IsColumn() && o.ColumnType() == "embedded")
}

func (o *ObjectField) IsSearchable() bool {
	t := getNamedType(o.Def.Type).(*ast.Named)
	return t.Name.Value == "String" || t.Name.Value == "Int" || t.Name.Value == "Float"
}

func (o *ObjectField) IsString() bool {
	t := getNamedType(o.Def.Type).(*ast.Named)
	return t.Name.Value == "String"
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
	return o.GoTypeWithPointer(true)
}
func (o *ObjectField) GoTypeWithPointer(showPointer bool) string {
	t := o.Def.Type
	st := ""

	if o.IsOptional() && showPointer {
		st += "*"
	} else {
		t = getNullableType(t)
	}

	if isListType(t) {
		st += "[]*"
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
	_gorm := fmt.Sprintf("default:null")
	_valid := ""

	dateArr := []interface{}{"createdAt", "updatedAt", "deletedAt", "createdBy", "updatedBy", "deletedBy"}
	// required/是否必填，type/正则校验类型，repeat/是否允许重复数据，relation/是否和公司id一起关联，join/重复数据和字段join，edit/是否允许编辑
	fields := []interface{}{}

	if o.Name() == "id" {
		_gorm = "type:varchar(36) comment 'uuid';primary_key;unique_index;NOT NULL;"
	}

	if IndexOf(dateArr, o.Name()) != -1 {
		tye := "type:varchar(255)"

		comment := "null;default:null"
		switch o.Name() {
		case "createdAt":
			tye = "type:bigint(13)"
			comment = "'createdAt';default:null;index:created_at;"
		case "updatedAt":
			tye = "type:bigint(13)"
			comment = "'updatedAt';default:null;"
		case "deletedAt":
			tye = "type:bigint(13)"
			comment = "'deletedAt';default:null;"
		case "createdBy":
			tye = "type:varchar(36)"
			comment = "'createdBy';default:null;"
		case "updatedBy":
			tye = "type:varchar(36)"
			comment = "'updatedBy';default:null;"
		case "deletedBy":
			tye = "type:varchar(36)"
			comment = "'deletedBy';default:null;"
		}
		_gorm = fmt.Sprintf("%s comment %s", tye, comment)

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
				if arg.Value.GetValue() != nil && IndexOf(fields, arg.Name.Value) != -1 {
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

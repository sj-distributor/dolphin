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

func (o *ObjectField) Type() string {
	if namedType, ok := o.Def.Type.(*ast.Named); ok {
		return namedType.Name.Value
	}
	return "unknown"
}

func (o *ObjectField) RelationshipTypeName() string {
	return o.Def.Description.Kind
}

func (o *ObjectField) RelationshipName() string {
	if !o.IsRelationship() {
		return o.MethodName()
	}
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

func (o *ObjectField) IsRelationshipRequired() bool {
	return o.Obj.Field(strcase.ToLowerCamel(o.RelationshipName())).IsRequired()
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
	return strings.HasSuffix(o.Name(), "Id") && o.Type() == "ID" || strings.HasSuffix(o.Name(), "Id") && o.TargetType() == "ID"
}

// IsCreatable ...
func (o *ObjectField) IsCreatable() bool {
	return !(o.Name() == "id" || o.Name() == "createdAt" || o.Name() == "updatedAt" || o.Name() == "deletedAt" || o.Name() == "createdBy" || o.Name() == "updatedBy" || o.Name() == "deletedBy")
}

func (o *ObjectField) IsUpdatable() bool {
	return !(o.Name() == "createdAt" || o.Name() == "updatedAt" || o.Name() == "deletedAt" || o.Name() == "createdBy" || o.Name() == "updatedBy" || o.Name() == "deletedBy")
}

func (o *ObjectField) IsCreataDocs() bool {
	return o.IsUpdatable() && !(o.Name() == "isDelete" || o.Name() == "weight" || o.Name() == "state")
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

	if o.Name() == "isDelete" {
		_gorm = "type:int(2) comment '是否删除：1/正常、0/删除';default:1;index:is_delete;"
	}

	if o.Name() == "weight" {
		_gorm = "type:int(11) comment '权重：用来排序';default:1;index:weight;"
	}

	if o.Name() == "state" {
		_gorm = "type:int(2) comment '状态：1/正常、0/禁用';default:1;index:state;"
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

func (o *ObjectField) GetArgValue(name string) map[string]map[string]string {
	for _, d := range o.Def.Directives {
		if d.Name.Value == name && len(d.Arguments) > 0 {
			argArr := map[string]map[string]string{
				name: map[string]string{},
			}
			for _, child := range d.Arguments {
				argArr[name][child.Name.Value] = child.Value.GetValue().(string)
			}
			return argArr
		}
	}

	return map[string]map[string]string{}
}

var columnMap = map[string]map[string]string{
	"id": {
		"comment": "uuid",
		"type":    "varchar(36)",
	},
	"createdAt": {
		"comment": "创建时间",
		"type":    "bigint(13)",
	},
	"updatedAt": {
		"comment": "更新时间",
		"type":    "bigint(13)",
	},
	"deletedAt": {
		"comment": "删除时间",
		"type":    "bigint(13)",
	},
	"deletedBy": {
		"comment": "删除人",
		"type":    "varchar(36)",
	},
	"updatedBy": {
		"comment": "更新人",
		"type":    "varchar(36)",
	},
	"createdBy": {
		"comment": "创建人",
		"type":    "varchar(36)",
	},
	"state": {
		"comment": "状态：1/正常、2/禁用、3/下架",
		"type":    "int(2)",
	},
	"weight": {
		"comment": "权重：用来排序",
		"type":    "int(2)",
	},
	"isDelete": {
		"comment": "是否删除：1/正常、2/删除",
		"type":    "int(2)",
	},
}

// 获取字段说明
func (o *ObjectField) GetComment() string {
	column := o.GetArgValue("column")
	value := column["column"]["gorm"]
	str := ""
	if value != "" {
		str = RegexpReplace(value, `comment '`, `';`)
	} else if o.Name() != "id" && o.TargetType() == "ID" {
		str = o.RelationshipName() + "实例Id"
	} else {
		str = columnMap[o.Name()]["comment"]
	}
	return str
}

// 备注说明字段
func (o *ObjectField) GetRemark() string {
	str := ""

	column := o.GetArgValue("column")
	gorm := column["column"]["gorm"]

	if gorm != "" {
		value := RegexpReplace(gorm, `default:`, `;`)
		if value != "" {
			str = "default:" + value
		}
	}
	switch o.Name() {
	case "id":
		str = "create方法不是必填"
	}
	return str
}

// 获取字段说明
func (o *ObjectField) GetType() string {
	column := o.GetArgValue("column")
	value := column["column"]["gorm"]
	str := ""

	if value != "" {
		str = RegexpReplace(value, `type:`, ` `)
	} else if o.Name() != "id" && o.TargetType() == "ID" {
		str = "varchar(36)"
	} else {
		str = columnMap[o.Name()]["type"]
	}
	return str
}

// 获取正则验证
func (o *ObjectField) GetValidator() string {
	column := o.GetArgValue("validator")
	value := column["validator"]["type"]
	str := ""
	if value != "" {
		str = value
	} else {
		switch o.Name() {
		case "state":
			str = "justInt"
		case "weight":
			str = "justInt"
		}
	}
	return str
}

// 获取Arguments
func (o *ObjectField) Arguments() string {
	argString := ""
	for key, child := range o.ArgumentsValue() {

		nullType := child.NonNullType()

		targetType := child.TargetType()

		if child.IsListType() {
			targetType = "[" + targetType + "]"
		}
		if key != len(o.ArgumentsValue())-1 {
			argString = argString + "$" + child.Name() + ": " + targetType + nullType + ", "
		} else {
			argString = argString + "$" + child.Name() + ": " + targetType + nullType
		}
	}

	if argString != "" {
		argString = "(" + argString + ")"
	}

	return argString
}

// 获取Input
func (o *ObjectField) Inputs() string {
	argString := ""

	for key, child := range o.ArgumentsValue() {
		if key != len(o.ArgumentsValue())-1 {
			argString = argString + child.Name() + ": $" + child.Name() + ", "
		} else {
			argString = argString + child.Name() + ": $" + child.Name()
		}
	}

	if argString != "" {
		argString = "(" + argString + ")"
	}

	return argString
}

// 表名
func (o *ObjectField) EntityName() string {
	if len(o.Obj.Def.Directives) > 0 && len(o.Obj.Def.Directives[0].Arguments) > 0 {
		title := o.Obj.Def.Directives[0].Arguments[0].Value.GetValue()
		return title.(string)
	}
	return o.Name()
}

// 获取是否默认显示
func (o *ObjectField) GetDefault() string {
	res := o.GetArgValue("entity")
	entity := res["entity"]

	return entity["default"]
}

// 获取是否默认显示
func (o *ObjectField) GetTableName() string {
	res := o.GetArgValue("entity")
	entity := res["entity"]
	return entity["title"]
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

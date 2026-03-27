package model

import (
	"fmt"
	"strings"

	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/printer"
	"github.com/iancoleman/strcase"
)

// ============================================================
// Go 类型映射
// ============================================================

// goTypeMap 定义 GraphQL 类型到 Go 类型的映射
var goTypeMap = map[string]string{
	"String":  "string",
	"Time":    "time.Time",
	"ID":      "string",
	"Float":   "float64",
	"Int":     "int64",
	"Boolean": "bool",
	"Any":     "interface{}",
}

// ============================================================
// 系统字段常量定义
// ============================================================

// nonCreatableFields 定义创建时应排除的字段
var nonCreatableFields = map[string]bool{
	"id":        true,
	"createdAt": true,
	"updatedAt": true,
	"deletedAt": true,
	"createdBy": true,
	"updatedBy": true,
	"deletedBy": true,
}

// nonUpdatableFields 定义更新时应排除的字段
var nonUpdatableFields = map[string]bool{
	"createdAt": true,
	"updatedAt": true,
	"deletedAt": true,
	"createdBy": true,
	"updatedBy": true,
	"deletedBy": true,
}

// nonDocFields 定义文档中应排除的字段
var nonDocFields = map[string]bool{
	"isDelete": true,
	"weight":   true,
	"state":    true,
}

// indexedFields 定义需要创建索引的字段
var indexedFields = map[string]bool{
	"createdBy": true,
	"updatedBy": true,
	"deletedBy": true,
}

// ============================================================
// 列信息常量
// ============================================================

// columnInfo 存储系统字段的元信息
type columnInfo struct {
	comment string
	dbType  string
}

// columnMetadata 定义系统字段的数据库元信息
var columnMetadata = map[string]columnInfo{
	"id":        {comment: "uuid", dbType: "varchar(36)"},
	"createdAt": {comment: "创建时间", dbType: "bigint(13)"},
	"updatedAt": {comment: "更新时间", dbType: "bigint(13)"},
	"deletedAt": {comment: "删除时间", dbType: "bigint(13)"},
	"deletedBy": {comment: "删除人", dbType: "varchar(36)"},
	"updatedBy": {comment: "更新人", dbType: "varchar(36)"},
	"createdBy": {comment: "创建人", dbType: "varchar(36)"},
	"state":     {comment: "状态：1/正常、2/禁用", dbType: "int(2)"},
	"weight":    {comment: "权重：用来排序", dbType: "int(2)"},
	"isDelete":  {comment: "是否删除：1/正常、2/删除", dbType: "int(2)"},
}

// ============================================================
// ObjectField 结构体定义
// ============================================================

// ObjectField 表示 GraphQL 对象的一个字段
type ObjectField struct {
	Def *ast.FieldDefinition
	Obj *Object
}

// ============================================================
// 名称相关方法
// ============================================================

// Name 返回字段的原始名称
func (o *ObjectField) Name() string {
	return o.Def.Name.Value
}

// LowerName 返回字段名称的小驼峰形式
func (o *ObjectField) LowerName() string {
	return strcase.ToLowerCamel(o.Name())
}

// ToSnakeName 返回字段名称的蛇形命名形式
func (o *ObjectField) ToSnakeName() string {
	return strcase.ToSnake(o.Name())
}

// MethodName 返回适用于 Go 方法的字段名称（首字母大写）
func (o *ObjectField) MethodName() string {
	return templates.ToGo(o.Name())
}

// HasName 检查字段是否具有指定名称
func (o *ObjectField) HasName(name string) bool {
	return o.Name() == name
}

// RelationshipName 返回关系名称（移除 ID 后缀）
func (o *ObjectField) RelationshipName() string {
	if !o.IsRelationship() {
		return o.MethodName()
	}
	return strings.Replace(o.MethodName(), "ID", "", -1)
}

// RelationshipTypeName 返回关系的类型名称
func (o *ObjectField) RelationshipTypeName() string {
	return o.Def.Description.Kind
}

// Signature 返回字段的 GraphQL 签名字符串（不含 Dolphin 指令）
func (o *ObjectField) Signature() string {
	if o.Def == nil {
		return ""
	}

	// 备份原始指令
	origFieldDirectives := o.Def.Directives
	o.Def.Directives = cleanDirectives(o.Def.Directives)

	// 备份参数原始指令
	origArgDirectives := make([][]*ast.Directive, len(o.Def.Arguments))
	for i, arg := range o.Def.Arguments {
		origArgDirectives[i] = arg.Directives
		arg.Directives = cleanDirectives(arg.Directives)
	}

	// 打印签名
	printed := printer.Print(o.Def)

	// 还原原始指令
	o.Def.Directives = origFieldDirectives
	for i, arg := range o.Def.Arguments {
		arg.Directives = origArgDirectives[i]
	}

	if s, ok := printed.(string); ok {
		// 移除可能存在的末尾换行符或空格，并对双引号进行转义以确保 JSON 安全
		res := strings.TrimSpace(s)
		return strings.ReplaceAll(res, "\"", "\\\"")
	}
	return ""
}

// ============================================================
// 类型相关方法
// ============================================================

// Type 返回字段的 GraphQL 类型名称
func (o *ObjectField) Type() string {
	if namedType, ok := o.Def.Type.(*ast.Named); ok {
		return namedType.Name.Value
	}
	return "unknown"
}

// TargetType 返回字段的目标类型（解包 NonNull 和 List）
func (o *ObjectField) TargetType() string {
	nt := getNamedType(o.Def.Type).(*ast.Named)
	return nt.Name.Value
}

// GoType 返回字段对应的 Go 类型
func (o *ObjectField) GoType() string {
	return o.GoTypeWithPointer()
}

// GoTypeWithPointer 返回带指针前缀的 Go 类型
func (o *ObjectField) GoTypeWithPointer() string {
	t := o.Def.Type
	st := ""

	// 处理可选类型的指针
	if o.IsOptional() {
		st += "*"
	} else {
		t = getNullableType(t)
	}

	// 处理列表类型
	if isListType(t) {
		if o.IsRequired() {
			st = "[]"
		} else {
			st += "[]*"
		}
	}

	// 映射 Go 类型
	v, ok := getNamedType(o.Def.Type).(*ast.Named)
	if ok {
		if goType, known := goTypeMap[v.Name.Value]; known {
			st += goType
		} else {
			st += v.Name.Value
		}
	}

	return st
}

// InputType 返回字段的输入类型
func (o *ObjectField) InputType() ast.Type {
	t := o.Def.Type

	if o.IsIdentifier() {
		t = nonNull(getNamedType(t))
	}

	isList := o.IsList()
	isOptional := o.IsOptional()

	// 处理嵌入式列
	if o.IsEmbeddedColumn() {
		nt := getNamedType(t).(*ast.Named)
		t = namedType(nt.Name.Value + "Input")

		if isList {
			t = listType(t)
		}
		if !isOptional {
			t = nonNull(t)
		}
	}

	// 关系标识符始终可选
	if o.IsRelationshipIdentifier() {
		t = getNullableType(t)
	}

	return t
}

// ============================================================
// 类型判断方法
// ============================================================

// IsScalarType 判断是否为标量类型
func (o *ObjectField) IsScalarType() bool {
	return o.Obj.Model.HasScalar(o.TargetType())
}

// IsEnumType 判断是否为枚举类型
func (o *ObjectField) IsEnumType() bool {
	return o.Obj.Model.HasEnum(o.TargetType())
}

// IsID 判断是否为 ID 类型
func (o *ObjectField) IsID() bool {
	return o.TargetType() == "ID"
}

// IsInt 判断是否为 Int 类型
func (o *ObjectField) IsInt() bool {
	return o.TargetType() == "Int"
}

// IsString 判断是否为 String 类型
func (o *ObjectField) IsString() bool {
	return o.TargetType() == "String"
}

// IsList 判断是否为列表类型
func (o *ObjectField) IsList() bool {
	return isListType(o.Def.Type)
}

// IsListType 判断是否为列表类型（解包 NonNull 后）
func (o *ObjectField) IsListType() bool {
	return isListType(getNullableType(o.Def.Type))
}

// IsOptional 判断是否为可选类型（非 NonNull）
func (o *ObjectField) IsOptional() bool {
	return !isNonNullType(o.Def.Type)
}

// IsRequired 判断是否为必填类型
func (o *ObjectField) IsRequired() bool {
	return isNonNullType(o.Def.Type)
}

// IsReadonlyType 判断是否为只读类型
func (o *ObjectField) IsReadonlyType() bool {
	if o.IsEmbeddedColumn() {
		return false
	}
	return !(o.IsScalarType() || o.IsEnumType()) || o.Obj.Model.HasObject(o.TargetType())
}

// IsWritableType 判断是否为可写类型
func (o *ObjectField) IsWritableType() bool {
	return !o.IsReadonlyType()
}

// ============================================================
// 字段角色判断方法
// ============================================================

// IsIdentifier 判断是否为 ID 标识符字段
func (o *ObjectField) IsIdentifier() bool {
	return o.HasName("id")
}

// IsRelationshipIdentifier 判断是否为关系 ID 字段（如 userId）
func (o *ObjectField) IsRelationshipIdentifier() bool {
	name := o.Name()
	targetType := o.TargetType()
	return strings.HasSuffix(name, "Id") && (o.Type() == "ID" || targetType == "ID")
}

// IsShardingID 判断是否为分片 ID 字段
func (o *ObjectField) IsShardingID() bool {
	return o.Name() == "shardingId"
}

// IsHasUpperId 判断是否为关系字段且包含 Id
func (o *ObjectField) IsHasUpperId() bool {
	return strings.Contains(o.Name(), "Id") && o.IsRelationship()
}

// ============================================================
// CRUD 权限判断方法
// ============================================================

// IsCreatable 判断字段是否可用于创建操作
func (o *ObjectField) IsCreatable() bool {
	return !nonCreatableFields[o.Name()]
}

// IsUpdatable 判断字段是否可用于更新操作
func (o *ObjectField) IsUpdatable() bool {
	return !nonUpdatableFields[o.Name()]
}

// IsCreataDocs 判断字段是否应出现在创建文档中
func (o *ObjectField) IsCreataDocs() bool {
	return o.IsUpdatable() && !nonDocFields[o.Name()]
}

// IsSearchable 判断字段是否可搜索
func (o *ObjectField) IsSearchable() bool {
	targetType := o.TargetType()
	return targetType == "String" || targetType == "Int" || targetType == "Float"
}

// IsSortable 判断字段是否可排序
func (o *ObjectField) IsSortable() bool {
	return !o.IsReadonlyType() && o.IsScalarType()
}

// ============================================================
// 指令相关方法
// ============================================================

// Directive 获取字段上的指定指令
func (o *ObjectField) Directive(name string) *ast.Directive {
	for _, d := range o.Def.Directives {
		if d.Name.Value == name {
			return d
		}
	}
	return nil
}

// IsColumn 判断是否为数据库列
func (o *ObjectField) IsColumn() bool {
	return o.HasDirective("column")
}

// IsRelationship 判断是否为关系字段
func (o *ObjectField) IsRelationship() bool {
	return o.HasDirective("relationship")
}

// IsRelationshipRequired 判断关系是否为必填
func (o *ObjectField) IsRelationshipRequired() bool {
	return o.Obj.Field(strcase.ToLowerCamel(o.RelationshipName())).IsRequired()
}

// IsEmbedded 判断是否为嵌入字段
func (o *ObjectField) IsEmbedded() bool {
	return !o.IsColumn() && !o.IsRelationship()
}

// IsEmbeddedColumn 判断是否为嵌入式列
func (o *ObjectField) IsEmbeddedColumn() bool {
	return o.IsColumn() && o.ColumnType() == "embedded"
}

// NeedsQueryResolver 判断是否需要查询解析器
func (o *ObjectField) NeedsQueryResolver() bool {
	return o.IsEmbedded()
}

// ============================================================
// 目标对象相关方法
// ============================================================

// HasTargetObject 判断是否存在目标对象
func (o *ObjectField) HasTargetObject() bool {
	return o.Obj.Model.HasObject(o.TargetType())
}

// TargetObject 返回目标对象
func (o *ObjectField) TargetObject() *Object {
	obj := o.Obj.Model.Object(o.TargetType())
	return &obj
}

// HasTargetObjectExtension 判断目标对象是否有扩展
func (o *ObjectField) HasTargetObjectExtension() bool {
	return o.Obj.Model.HasObjectExtension(o.TargetType())
}

// TargetObjectExtension 返回目标对象的扩展
func (o *ObjectField) TargetObjectExtension() *ObjectExtension {
	e := o.Obj.Model.ObjectExtension(o.TargetType())
	return &e
}

// HasTargetTypeWithIDField 判断目标类型是否有 ID 字段
func (o *ObjectField) HasTargetTypeWithIDField() bool {
	if o.HasTargetObject() && o.TargetObject().HasField("id") {
		return true
	}
	if o.HasTargetObjectExtension() && o.TargetObjectExtension().Object.HasField("id") {
		return true
	}
	return false
}

// ============================================================
// 参数相关方法
// ============================================================

// ArgumentsValue 返回字段的参数列表
func (o *ObjectField) ArgumentsValue() []ObjectFieldInput {
	arguments := []ObjectFieldInput{}
	for _, f := range o.Def.Arguments {
		arguments = append(arguments, ObjectFieldInput{f, o})
	}
	return arguments
}

func (o *ObjectField) GetTypeData() []TypeData {
	result := []TypeData{}
	visited := make(map[string]bool)
	queue := []string{}

	// Return type
	queue = append(queue, o.TargetType())

	// Arguments
	for _, arg := range o.ArgumentsValue() {
		queue = append(queue, arg.TargetType())
	}

	for len(queue) > 0 {
		typeName := queue[0]
		queue = queue[1:]

		if visited[typeName] || defaultScalars[typeName] {
			continue
		}
		visited[typeName] = true

		def := o.Obj.Model.GetDefinition(typeName)
		if def == nil {
			continue
		}

		typeData := TypeData{Name: typeName}

		switch d := def.(type) {
		case *ast.ObjectDefinition:
			for _, field := range d.Fields {
				fieldType := getNamedType(field.Type).(*ast.Named).Name.Value
				queue = append(queue, fieldType)

				typeData.Fields = append(typeData.Fields, FieldMetadata{
					Name:      field.Name.Value,
					Type:      fieldType,
					Desc:      o.getFieldDesc(field),
					Required:  fmt.Sprintf("%t", isNonNullType(field.Type)),
					Validator: o.getValidatorFromDirectives(field.Directives),
				})
			}
		case *ast.InputObjectDefinition:
			for _, field := range d.Fields {
				fieldType := getNamedType(field.Type).(*ast.Named).Name.Value
				queue = append(queue, fieldType)

				typeData.Fields = append(typeData.Fields, FieldMetadata{
					Name:      field.Name.Value,
					Type:      fieldType,
					Desc:      o.getInputValueDesc(field),
					Required:  fmt.Sprintf("%t", isNonNullType(field.Type)),
					Validator: o.getValidatorFromDirectives(field.Directives),
				})
			}
		case *ast.EnumDefinition:
			for _, val := range d.Values {
				typeData.Fields = append(typeData.Fields, FieldMetadata{
					Name:      val.Name.Value,
					Type:      "Enum",
					Desc:      o.getEnumTitle(val),
					Required:  "false",
					Validator: "",
				})
			}
		}
		result = append(result, typeData)
	}

	return result
}

func (o *ObjectField) getFieldDesc(f *ast.FieldDefinition) string {
	for _, d := range f.Directives {
		if d.Name.Value == "entity" {
			for _, arg := range d.Arguments {
				if arg.Name.Value == "title" {
					if val, ok := arg.Value.GetValue().(string); ok {
						return val
					}
				}
			}
		}
	}
	return f.Name.Value
}

func (o *ObjectField) getInputValueDesc(f *ast.InputValueDefinition) string {
	for _, d := range f.Directives {
		if d.Name.Value == "entity" {
			for _, arg := range d.Arguments {
				if arg.Name.Value == "title" {
					if val, ok := arg.Value.GetValue().(string); ok {
						return val
					}
				}
			}
		}
	}
	return f.Name.Value
}

func (o *ObjectField) getValidatorFromDirectives(directives []*ast.Directive) string {
	for _, d := range directives {
		if d.Name.Value == "validator" {
			for _, arg := range d.Arguments {
				if arg.Name.Value == "type" {
					if val, ok := arg.Value.GetValue().(string); ok {
						return val
					}
				}
			}
		}
	}
	return ""
}

func (o *ObjectField) getEnumTitle(v *ast.EnumValueDefinition) string {
	for _, d := range v.Directives {
		if d.Name.Value == "entity" {
			for _, arg := range d.Arguments {
				if arg.Name.Value == "title" {
					if val, ok := arg.Value.GetValue().(string); ok {
						return val
					}
				}
			}
		}
	}
	return v.Name.Value
}

// Arguments 返回格式化的 GraphQL 参数字符串
func (o *ObjectField) Arguments() string {
	args := o.ArgumentsValue()
	if len(args) == 0 {
		return ""
	}

	var parts []string
	for _, child := range args {
		targetType := child.TargetType()
		if child.IsListType() {
			targetType = "[" + targetType + "]"
		}
		parts = append(parts, "$"+child.Name()+": "+targetType+child.NonNullType())
	}

	return "(" + strings.Join(parts, ", ") + ")"
}

// Inputs 返回格式化的 GraphQL 输入字符串
func (o *ObjectField) Inputs() string {
	args := o.ArgumentsValue()
	if len(args) == 0 {
		return ""
	}

	var parts []string
	for _, child := range args {
		parts = append(parts, child.Name()+": $"+child.Name())
	}

	return "(" + strings.Join(parts, ", ") + ")"
}

// ============================================================
// 指令值获取方法
// ============================================================

// GetArgValue 获取指令的参数值
func (o *ObjectField) GetArgValue(name string) map[string]map[string]string {
	for _, d := range o.Def.Directives {
		if d.Name.Value == name && len(d.Arguments) > 0 {
			result := map[string]map[string]string{
				name: {},
			}
			for _, arg := range d.Arguments {
				if val := arg.Value.GetValue(); val != nil {
					result[name][arg.Name.Value] = val.(string)
				}
			}
			return result
		}
	}
	return map[string]map[string]string{}
}

// ============================================================
// 元数据获取方法
// ============================================================

// GetComment 获取字段的注释说明
func (o *ObjectField) GetComment() string {
	column := o.GetArgValue("column")
	if gormValue := column["column"]["gorm"]; gormValue != "" {
		return RegexpReplace(gormValue, `comment '`, `';`)
	}

	// 关系 ID 字段
	if o.Name() != "id" && o.TargetType() == "ID" {
		return o.RelationshipName() + "实例Id"
	}

	// 系统字段
	if meta, ok := columnMetadata[o.Name()]; ok {
		return meta.comment
	}

	return ""
}

// GetType 获取字段的数据库类型
func (o *ObjectField) GetType() string {
	column := o.GetArgValue("column")
	if gormValue := column["column"]["gorm"]; gormValue != "" {
		return RegexpReplace(gormValue, `type:`, ` `)
	}

	// 关系 ID 字段
	if o.Name() != "id" && o.TargetType() == "ID" {
		return "varchar(36)"
	}

	// 系统字段
	if meta, ok := columnMetadata[o.Name()]; ok {
		return meta.dbType
	}

	return ""
}

// GetRemark 获取字段的备注说明
func (o *ObjectField) GetRemark() string {
	column := o.GetArgValue("column")
	if gormValue := column["column"]["gorm"]; gormValue != "" {
		if defaultVal := RegexpReplace(gormValue, `default:`, `;`); defaultVal != "" {
			return "default:" + defaultVal
		}
	}

	if o.Name() == "id" {
		return "create方法不是必填"
	}

	return ""
}

// GetValidator 获取字段的验证器类型
func (o *ObjectField) GetValidator() string {
	column := o.GetArgValue("validator")
	if validatorType := column["validator"]["type"]; validatorType != "" {
		return validatorType
	}

	// 特殊字段的默认验证器
	switch o.Name() {
	case "state", "weight":
		return "justInt"
	}

	return ""
}

// GetDefault 获取字段的默认显示设置
func (o *ObjectField) GetDefault() string {
	res := o.GetArgValue("entity")
	return res["entity"]["default"]
}

// GetTableName 获取实体的表名
func (o *ObjectField) GetTableName() string {
	res := o.GetArgValue("entity")
	return res["entity"]["title"]
}

// EntityName 获取实体名称
func (o *ObjectField) EntityName() string {
	if len(o.Obj.Def.Directives) > 0 && len(o.Obj.Def.Directives[0].Arguments) > 0 {
		if title := o.Obj.Def.Directives[0].Arguments[0].Value.GetValue(); title != nil {
			return title.(string)
		}
	}
	return o.Name()
}

// ============================================================
// GORM 标签生成
// ============================================================

// ModelTags 生成字段的 GORM 模型标签
func (o *ObjectField) ModelTags() string {
	gormTag := o.buildGormTag()
	validTag := o.buildValidatorTag()

	if validTag != "" {
		return fmt.Sprintf(`json:"%s" gorm:"%s" validator:"%s"`, o.Name(), gormTag, validTag)
	}
	return fmt.Sprintf(`json:"%s" gorm:"%s"`, o.Name(), gormTag)
}

// buildGormTag 构建 GORM 标签
func (o *ObjectField) buildGormTag() string {
	// 检查自定义 gorm 指令
	for _, d := range o.Def.Directives {
		if d.Name.Value == "column" {
			for _, arg := range d.Arguments {
				if arg.Name.Value == "gorm" {
					return fmt.Sprintf("%v", arg.Value.GetValue())
				}
			}
		}
	}

	// 系统字段的预定义标签
	name := o.Name()
	switch name {
	case "id":
		return "type:varchar(36);comment:'uuid';primaryKey;uniqueIndex;NOT NULL;"
	case "isDelete":
		return "type:int(2);comment:'是否删除：1/正常、2/删除';default:1;index:is_delete;"
	case "weight":
		return "type:int(11);comment:'权重：用来排序';default:1;index:weight;"
	case "state":
		return "type:int(2);comment:'状态：1/正常、2/禁用';default:1;index:state;"
	}

	// 构建默认标签
	gormTag := ""
	if o.IsString() {
		gormTag = fmt.Sprintf("type:varchar(255);comment:'%s';default:null;", o.ToSnakeName())
	} else if o.IsID() {
		gormTag = fmt.Sprintf("type:varchar(36);comment:'%s';default:null;", o.ToSnakeName())
	} else if o.IsInt() {
		gormTag = fmt.Sprintf("type:bigint(13);comment:'%s';default:null;", o.ToSnakeName())
	} else {
		gormTag = "default:null"
	}

	// 索引字段
	if indexedFields[name] {
		gormTag += fmt.Sprintf("index:%s;", o.ToSnakeName())
	}

	// 时间戳字段
	if name == "createdAt" {
		gormTag += " autoCreateTime:milli;"
	} else if name == "updatedAt" {
		gormTag += " autoUpdateTime:milli;"
	}

	return gormTag
}

// buildValidatorTag 构建验证器标签
func (o *ObjectField) buildValidatorTag() string {
	var parts []string
	for _, d := range o.Def.Directives {
		if d.Name.Value == "validator" {
			for _, arg := range d.Arguments {
				if val := arg.Value.GetValue(); val != nil {
					parts = append(parts, arg.Name.Value+":"+val.(string))
				}
			}
		}
	}
	return strings.Join(parts, ";")
}

// ============================================================
// 过滤器映射
// ============================================================

// FilterMappingItem 表示过滤器映射项
type FilterMappingItem struct {
	Suffix      string
	Operator    string
	InputType   ast.Type
	ValueFormat string
}

// SuffixCamel 返回后缀的驼峰形式
func (f *FilterMappingItem) SuffixCamel() string {
	return strcase.ToCamel(f.Suffix)
}

// IsLike 判断是否为 LIKE 操作符
func (f *FilterMappingItem) IsLike() bool {
	return f.SuffixCamel() == "Like"
}

// WrapValueVariable 用格式包装值变量
func (f *FilterMappingItem) WrapValueVariable(v string) string {
	return fmt.Sprintf(f.ValueFormat, v)
}

// FilterMapping 返回字段的过滤器映射列表
func (o *ObjectField) FilterMapping() []FilterMappingItem {
	t := getNamedType(o.Def.Type)

	// 基础过滤器
	mapping := []FilterMappingItem{
		{Suffix: "", Operator: "= ?", InputType: t, ValueFormat: "%s"},
		{Suffix: "_ne", Operator: "!= ?", InputType: t, ValueFormat: "%s"},
		{Suffix: "_gt", Operator: "> ?", InputType: t, ValueFormat: "%s"},
		{Suffix: "_lt", Operator: "< ?", InputType: t, ValueFormat: "%s"},
		{Suffix: "_gte", Operator: ">= ?", InputType: t, ValueFormat: "%s"},
		{Suffix: "_lte", Operator: "<= ?", InputType: t, ValueFormat: "%s"},
		{Suffix: "_in", Operator: "IN (?)", InputType: listType(nonNull(t)), ValueFormat: "%s"},
	}

	// 字符串类型特有的过滤器
	if namedType := t.(*ast.Named); namedType.Name.Value == "String" {
		mapping = append(mapping,
			FilterMappingItem{
				Suffix:      "_like",
				Operator:    "LIKE ?",
				InputType:   t,
				ValueFormat: `strings.Replace(strings.Replace(*%s,"?","_",-1),"*","%%",-1)`,
			},
			FilterMappingItem{
				Suffix:      "_prefix",
				Operator:    "LIKE ?",
				InputType:   t,
				ValueFormat: `fmt.Sprintf("%%s%%%%",*%s)`,
			},
			FilterMappingItem{
				Suffix:      "_suffix",
				Operator:    "LIKE ?",
				InputType:   t,
				ValueFormat: `fmt.Sprintf("%%%%%%s",*%s)`,
			},
		)
	}

	return mapping
}

// ============================================================
// Model 上的标量和枚举方法
// ============================================================

// HasScalar 判断模型是否有指定的标量类型
func (m *Model) HasScalar(name string) bool {
	if _, ok := defaultScalars[name]; ok {
		return true
	}
	for _, def := range m.Doc.Definitions {
		if scalar, ok := def.(*ast.ScalarDefinition); ok && scalar.Name.Value == name {
			return true
		}
	}
	return false
}

// HasEnum 判断模型是否有指定的枚举类型
func (m *Model) HasEnum(name string) bool {
	if _, ok := defaultScalars[name]; ok {
		return true
	}
	for _, def := range m.Doc.Definitions {
		if enum, ok := def.(*ast.EnumDefinition); ok && enum.Name.Value == name {
			return true
		}
	}
	return false
}

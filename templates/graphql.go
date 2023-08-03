package templates

var GraphqlApi = `[
	{
		"title": "全局字段",
		"name": "",
		"type": 0,
		"fields": [
			{ "name": "deletedBy", "desc": "删除人", "type": "varchar(36)", "required": "false", "validator": "", "remark": "仅供Query使用" },
			{ "name": "updatedBy", "desc": "更新人", "type": "varchar(36)", "required": "false", "validator": "", "remark": "仅供Query使用" },
			{ "name": "createdBy", "desc": "创建人", "type": "varchar(36)", "required": "false", "validator": "", "remark": "仅供Query使用" },
			{ "name": "deletedAt", "desc": "删除时间", "type": "bigint(13)", "required": "false", "validator": "", "remark": "仅供Query使用" },
			{ "name": "updatedAt", "desc": "更新时间", "type": "bigint(13)", "required": "false", "validator": "", "remark": "仅供Query使用" },
			{ "name": "createdAt", "desc": "创建时间", "type": "bigint(13)", "required": "false", "validator": "", "remark": "仅供Query使用" },
			{ "name": "isDelete", "desc": "是否删除：1/正常、2/删除", "type": "int(2)", "required": "false", "validator": "", "remark": "" },
			{ "name": "weight", "desc": "权重：用来排序", "type": "int(2)", "required": "false", "validator": "justInt", "remark": "" },
			{ "name": "state", "desc": "状态：1/正常、2/禁用、3/下架", "type": "int(2)", "required": "false", "validator": "justInt", "remark": "" }
		]
	},{{range $obj := .Model.Objects}}
  {
    "title": "{{$obj.EntityName}}",
    "name": "{{$obj.ToLowerPluralName}}",
    "type": 0,
    "fields": [
      {{range $col := $obj.Columns}}{{if $col.IsCreataDocs}}{ "name": "{{$col.Name}}", "desc": "{{$col.GetComment}}", "type": "{{$col.GetType}}", "required": "{{$col.IsRequired}}", "validator": "{{$col.GetValidator}}", "remark": "{{$col.GetRemark}}" },
			{{end}}{{end}}{{range $rel := $obj.Relationships}}{ "name": "{{$rel.Name}}", "desc": "{{$rel.Target.Name}}连表查询", "type": "relationship", "required": "false", "validator": "", "remark": "{{$rel.LowerName}}实例" }{{end}}
    ],
    "data": [
      { "name": "{{$obj.EntityName}}", "title": "列表", "api": "{{$obj.ToLowerPluralName}}", "type": "list", "method": "Query" },
      { "name": "{{$obj.EntityName}}", "title": "详情", "api": "{{$obj.LowerName}}", "type": "detail", "method": "Query" },
      { "name": "{{$obj.EntityName}}", "title": "新增", "api": "create{{$obj.Name}}", "type": "add", "method": "Mutation" },
      { "name": "{{$obj.EntityName}}", "title": "修改", "api": "update{{$obj.Name}}", "type": "edit", "method": "Mutation" },
      { "name": "{{$obj.EntityName}}", "title": "删除", "api": "delete{{$obj.PluralName}}", "type": "delete", "method": "Mutation" },
      { "name": "{{$obj.EntityName}}", "title": "恢复", "api": "recovery{{$obj.PluralName}}", "type": "recovery", "method": "Mutation" }
    ]
  },
{{end}}
{{- $extComma := "" -}}{{range $ext := .Model.ObjectExtensions}}{{$extComma}}{{$obj := $ext.Object}}{{range $col := $obj.Fields}}
	{
    "title": "{{$col.GetTableName}}",
    "name": "{{$col.Name}}",
    "type": 1,
    "default": {{$col.GetDefault}},
    "fields": [
      {{- $relComma := "" -}}{{range $rel := $col.ArgumentsValue}}{{$relComma}}
			{ "name": "{{$rel.Name}}", "desc": "{{$rel.Name}}", "type": "{{$rel.TargetType}}{{$rel.NonNullType}}", "required": "{{$rel.Required}}", "validator": "", "remark": "" }{{ $relComma = "," }}{{end}}
    ],
    "data": [
      { "name": "{{$col.GetTableName}}", "title": "{{$col.GetTableName}}", "api": "{{$col.Name}}", "type": "detail", "method": "{{$obj.ToCamel}}" }
    ]
  }{{ $extComma = "," }}{{end}}{{end}}
]
`

package templates

var Graphql = `{{range $obj := .Model.Objects}}
  # {{$obj.EntityName}} {{$obj.Name}} 接口字段
  fragment {{$obj.Name}}sFields on {{$obj.Name}} {
    {{range $col := $obj.Columns}}{{$col.Name}}
    {{end}}
  }

  # 列表
  query {{$obj.Name}}s ($currentPage: Int = 1, $perPage: Int = 10, $sort: [{{$obj.Name}}SortType!], $search: String, $filter: {{$obj.Name}}FilterType, $rand: Boolean = false) {
    {{$obj.ToLowerPluralName}}(current_page: $currentPage, per_page: $perPage, sort: $sort, q: $search, filter: $filter, rand: $rand) {
      data {
        {{range $col := $obj.Columns}}{{$col.Name}}
        {{end}}{{range $rel := $obj.Relationships}}{{$rel.Name}} {
          ...{{$rel.Target.Name}}sFields
        }
        {{end}}
      }
      current_page
      per_page
      total
      total_page
    }
  }

  # 详情
  query {{$obj.Name}}Detail ($id: ID, $search: String, $filter: {{$obj.Name}}FilterType) {
    {{$obj.LowerName}}(id: $id, q: $search, filter: $filter) {
      {{range $col := $obj.Columns}}{{$col.Name}}
      {{end}}{{range $rel := $obj.Relationships}}{{$rel.Name}} {
        ...{{$rel.Target.Name}}sFields
      }
      {{end}}
    }
  }

  # 新增
  mutation {{$obj.Name}}Add ($data: {{$obj.Name}}CreateInput!) {
    create{{$obj.Name}}(input: $data) {
      id
    }
  }

  # 修改
  mutation {{$obj.Name}}Edit ($id: ID!, $data: {{$obj.Name}}UpdateInput!) {
    update{{$obj.Name}}(id: $id, input: $data) {
      id
    }
  }

  # 删除
  mutation {{$obj.Name}}sDelete ($id: [ID!]!) {
    delete{{$obj.PluralName}}(id: $id)
  }

  # 恢复删除
  mutation {{$obj.Name}}sRecovery ($id: [ID!]!) {
    recovery{{$obj.PluralName}}(id: $id)
  }
{{end}}
{{range $ext := .Model.ObjectExtensions}}{{$obj := $ext.Object}}
  {{range $col := $obj.Fields}}
  # {{$col.LowerName}} 接口
  {{$obj.LowerName}} {{$col.LowerName}}{{$col.Arguments}} {
    {{$col.Name}}{{$col.Inputs}}{{if $col.IsReadonlyType}} {
      ...{{$col.TargetType}}sFields{{if $col.TargetObject.HasAnyRelationships}}
      {{range $rol := $col.TargetObject.Relationships}}
      {{$rol.Name}} {
        {{range $rel := $rol.Target.Columns}}{{$rel.Name}}
        {{end}}{{range $oRel := $rol.Target.Relationships}}{{$oRel.Name}} {
          ...{{$oRel.Target.Name}}sFields
        }
        {{end}}
      }{{end}}{{end}}
    }{{end}}
  }{{end}}{{end}}
`

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
			{{end}}{{end}}
			{{- $relComma := "" -}}{{range $rel := $obj.Relationships}}{{$relComma}}
			{ "name": "{{$rel.Name}}", "desc": "{{$rel.Target.Name}}连表查询", "type": "relationship", "required": "false", "validator": "", "remark": "{{$rel.LowerName}}实例" }{{ $relComma = "," }}{{end}}
    ],
    "data": [
      { "name": "{{$obj.EntityName}}", "title": "列表", "api": "{{$obj.ToLowerPluralName}}", "type": "list", "method": "Query", "code": "query {{$obj.Name}}s ($currentPage: Int = 1, $perPage: Int = 10, $sort: [{{$obj.Name}}SortType!], $search: String, $filter: {{$obj.Name}}FilterType, $rand: Boolean = false) {\n  {{$obj.ToLowerPluralName}}(current_page: $currentPage, per_page: $perPage, sort: $sort, q: $search, filter: $filter, rand: $rand) {\n    data {\n      {{range $col := $obj.Columns}}{{$col.Name}}\n      {{end}}{{range $rel := $obj.Relationships}}{{$rel.Name}} {\n        ...{{$rel.Target.Name}}sFields\n      }\n      {{end}}\n    }\n    current_page\n    per_page\n    total\n    total_page\n  }\n}" },
      { "name": "{{$obj.EntityName}}", "title": "详情", "api": "{{$obj.LowerName}}", "type": "detail", "method": "Query", "code": "query {{$obj.Name}}Detail ($id: ID, $search: String, $filter: {{$obj.Name}}FilterType) {\n  {{$obj.LowerName}}(id: $id, q: $search, filter: $filter) {\n    {{range $col := $obj.Columns}}{{$col.Name}}\n    {{end}}{{range $rel := $obj.Relationships}}{{$rel.Name}} {\n      ...{{$rel.Target.Name}}sFields\n    }\n    {{end}}\n  }\n}" },
      { "name": "{{$obj.EntityName}}", "title": "新增", "api": "create{{$obj.Name}}", "type": "add", "method": "Mutation", "code": "mutation {{$obj.Name}}Add ($data: {{$obj.Name}}CreateInput!) {\n  create{{$obj.Name}}(input: $data) {\n    id\n  }\n}" },
      { "name": "{{$obj.EntityName}}", "title": "修改", "api": "update{{$obj.Name}}", "type": "edit", "method": "Mutation", "code": "mutation {{$obj.Name}}Edit ($id: ID!, $data: {{$obj.Name}}UpdateInput!) {\n  update{{$obj.Name}}(id: $id, input: $data) {\n    id\n  }\n}" },
      { "name": "{{$obj.EntityName}}", "title": "删除", "api": "delete{{$obj.PluralName}}", "type": "delete", "method": "Mutation", "code": "mutation {{$obj.Name}}sDelete ($id: [ID!]!) {\n  delete{{$obj.PluralName}}(id: $id)\n}" },
      { "name": "{{$obj.EntityName}}", "title": "恢复", "api": "recovery{{$obj.PluralName}}", "type": "recovery", "method": "Mutation", "code": "mutation {{$obj.Name}}sRecovery ($id: [ID!]!) {\n  recovery{{$obj.PluralName}}(id: $id)\n}" }
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

package templates

var GQLGen = `# Generated with dolphin (gqlgen v0.17.85)
{{$config:=.Config}}
schema:
  - gen/*.graphqls

# Where should the generated server code go?
exec:
  filename: gen/generated.go
  package: gen

# Uncomment to enable federation
# federation:
#   filename: gen/federation.go
#   package: gen
#   version: 2
#   options:
#     computed_requires: true

# Where should any generated models go?
model:
  filename: gen/models_gen.go
  package: gen

# Where should the resolver implementations go?
# resolver:
  # layout: follow-schema
  # dir: gen
  # package: gen
  # filename_template: "{name}.resolvers.go"

# Options for code generation
# omit_complexity: true  # skip complexity calculation
# nullable_input_omittable: true  # wrap nullable input fields with Omittable
# call_argument_directives_with_null: true  # call argument directives even with null values
# omit_gqlgen_file_notice: false  # include file generation notice

autobind:
  - "{{.Config.Package}}/gen"

models:
  {{range $obj := .Model.ObjectEntities}}
  {{$obj.Name}}:
    model: {{$config.Package}}/gen.{{$obj.Name}}
    fields:{{range $col := $obj.Columns}}{{if $col.IsReadonlyType}}
      {{$col.Name}}:
        resolver: true{{end}}{{end}}{{range $rel := $obj.Relationships}}
      {{$rel.Name}}:
        resolver: true{{end}}
  {{$obj.Name}}ResultType:
    model: {{$config.Package}}/gen.{{$obj.Name}}ResultType
    fields:
      data:
        resolver: true
      total:
        resolver: true
      current_page:
        resolver: true
      per_page:
        resolver: true
      total_page:
        resolver: true
  Create{{$obj.Name}}Input:
    model: "map[string]interface{}"
  Update{{$obj.Name}}Input:
    model: "map[string]interface{}"
  {{end}}
  _Any:
    model: {{$config.Package}}/gen._Any
`

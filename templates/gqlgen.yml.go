package templates

var GQLGen = `# Generated with dolphin
{{$config:=.Config}}
schema:
  - schema.graphql
exec:
  filename: generated.go
  package: gen
model:
  filename: models_gen.go
  package: gen
resolver:
  filename: resolver.go
  type: Resolver
  package: gen
autobind:
- "{{.Config.Package}}/gen"

models:
  {{range $obj := .Model.ObjectEntities}}{{$obj.Name}}ResultType:
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
  {{$obj.Name}}CreateInput:
    model: "map[string]interface{}"
  {{$obj.Name}}UpdateInput:
    model: "map[string]interface{}"
  {{end}}
  _Any:
    model: {{$config.Package}}/gen._Any
`

/*
 * @Author: Marlon.M
 * @Email: maiguangyang@163.com
 * @Date: 2025-03-15 15:54:57
 */
package templates

var Sorting = `package gen

import (
	"context"

	"gorm.io/gorm"
)

{{range $obj := .Model.ObjectEntities}}
{{if not $obj.IsExtended}}
func (s {{$obj.Name}}SortType) Apply(ctx context.Context, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, TableName("{{$obj.TableName}}", ctx), sorts, joins)
}
func (s {{$obj.Name}}SortType) ApplyWithAlias(ctx context.Context, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := alias + "."

	{{range $col := $obj.Columns}}{{if $col.IsSortable}}
	if s.{{$col.MethodName}} != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("{{$col.Name}}")+" "+s.{{$col.MethodName}}.String())
	}
	{{end}}{{end}}

	{{range $rel := $obj.Relationships}}
	{{if not $rel.Target.IsExtended}}
	{{$varName := (printf "s.%s" $rel.MethodName)}}
	if {{$varName}} != nil {
		_alias := alias + "_{{$rel.Name}}"
		*joins = append(*joins, {{$rel.JoinString}})
		err := {{$varName}}.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}{{end}}{{end}}

	return nil
}
{{end}}
{{end}}
`

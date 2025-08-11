package templates

var Filters = `package gen

import (
	"context"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

{{range $obj := .Model.ObjectEntities}}
{{if not $obj.IsExtended}}
func (f *{{$obj.Name}}FilterType) IsEmpty(ctx context.Context) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *{{$obj.Name}}FilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, TableName("{{$obj.TableName}}", ctx), wheres, values, joins)
}
func (f *{{$obj.Name}}FilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := alias + "."

	_where, _values := f.WhereContent(aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_cs := []string{}
			err := or.ApplyWithAlias(ctx, alias, &_cs, &vs, &js)
			if err != nil {
				return err
			}
			cs = append(cs, strings.Join(_cs, " AND "))
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	{{range $rel := $obj.Relationships}}
	{{if not $rel.Target.IsExtended}}
	{{$varName := (printf "f.%s" $rel.MethodName)}}
	if {{$varName}} != nil {
		_alias := alias + "_{{$rel.Name}}"
		*joins = append(*joins, {{$rel.JoinString}})
		err := {{$varName}}.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}{{end}}{{end}}

	return nil
}

func (f *{{$obj.Name}}FilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}
	whereConditions := []string{}

	{{range $col := $obj.Columns}}{{if $col.IsWritableType}}
		{{range $fm := $col.FilterMapping}} {{$varName := (printf "f.%s%s" $col.MethodName $fm.SuffixCamel)}}
			if {{$varName}} != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("{{$col.Name}}")) == -1 {
				conditions = append(conditions, aliasPrefix + SnakeString("{{$col.Name}}")+" {{$fm.Operator}}")
				whereConditions = append(whereConditions, aliasPrefix+SnakeString("{{$col.Name}}"))
				{{if $fm.IsLike}}values = append(values, "%"+{{$fm.WrapValueVariable $varName}}+"%"){{else}}values = append(values, {{$fm.WrapValueVariable $varName}}){{end}}
			}
		{{end}}
		if f.{{$col.MethodName}}Null != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("{{$col.Name}}")) == -1 {
			whereConditions = append(whereConditions, aliasPrefix+SnakeString("{{$col.Name}}"))
			if *f.{{$col.MethodName}}Null {
				conditions = append(conditions, aliasPrefix+SnakeString("{{$col.Name}}")+" IS NULL" + " OR " + aliasPrefix+SnakeString("{{$col.Name}}")+" =''")
			} else {
				conditions = append(conditions, aliasPrefix+SnakeString("{{$col.Name}}")+" IS NOT NULL" + " OR " + aliasPrefix+SnakeString("{{$col.Name}}")+" <> ''")
			}
		}
	{{end}}
{{end}}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *{{$obj.Name}}FilterType) AndWith(f2 ...*{{$obj.Name}}FilterType) *{{$obj.Name}}FilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &{{$obj.Name}}FilterType{
		And: append(_f2,f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *{{$obj.Name}}FilterType) OrWith(f2 ...*{{$obj.Name}}FilterType) *{{$obj.Name}}FilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &{{$obj.Name}}FilterType{
		Or: append(_f2,f),
	}
}

{{end}}
{{end}}
`

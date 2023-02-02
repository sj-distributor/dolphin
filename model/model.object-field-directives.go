package model

func (o *ObjectField) HasDirective(name string) bool {
	return o.Directive(name) != nil
}

// ColumnType ...
func (o *ObjectField) ColumnType() (value string) {
	directive := o.Directive("column")
	if directive == nil {
		return
	}
	for _, arg := range directive.Arguments {
		if arg.Name.Value == "type" {
			val := arg.Value.GetValue()
			value, _ = val.(string)
			break
		}
	}
	return
}

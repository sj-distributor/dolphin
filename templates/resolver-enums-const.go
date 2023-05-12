package templates

var EnumsConst = `package enums

import (
	"{{.Config.Package}}/auth"
)

const (
	CannotBeEmpty string = "%s cannot be empty"
	DataNotChange string = "the data has not been modified"
)
`

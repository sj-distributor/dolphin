package templates

var Constants = `package gen

type key int

const (
	KeyPrincipalID          key = iota
	KeyLoaders              key = iota
	KeyJWTClaims            key = iota
	KeyMutationTransaction  key = iota
	KeyMutationEvents       key = iota
	KeyExecutableSchema     key = iota
	KeyHeader               key = iota
	KeyAuthorization        key = iota
	KeySecretKey            key = iota
)
`

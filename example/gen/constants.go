package gen

type key int

const (
	KeyMutationTransaction key = iota
	KeyMutationEvents      key = iota
	KeyPrincipalID         key = iota
	KeyJWTClaims           key = iota
	KeyExecutableSchema    key = iota
)

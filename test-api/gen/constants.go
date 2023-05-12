package gen

type key int

const (
	KeyPrincipalID         key = iota
	KeyLoaders             key = iota
	KeyJWTClaims           key = iota
	KeyMutationTransaction key = iota
	KeyMutationEvents      key = iota
	KeyExecutableSchema    key = iota
)

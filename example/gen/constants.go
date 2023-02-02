package gen

type key int

const (
	KeyPrincipalID         key    = iota
	KeyLoaders             key    = iota
	KeyJWTClaims           key    = iota
	KeyMutationTransaction key    = iota
	KeyMutationEvents      key    = iota
	KeyExecutableSchema    key    = iota
	SchemaSDL              string = `scalar Time

scalar Upload

type Query {
  todo(id: ID!): Todo
}

type Mutation {
  createTodo(input: TodoCreateInput!): Todo!
}

type Todo {
  id: ID!
  title: String!
  deletedBy: ID
  updatedBy: ID
  createdBy: ID
  updatedAt: Int
  createdAt: Int!
}

input TodoCreateInput {
  title: String!
}`
)

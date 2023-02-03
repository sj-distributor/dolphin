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
  updateTodo(id: ID!, input: TodoUpdateInput!): Todo!
}

type Todo {
  id: ID!
  title: String!
  remark: String
  deletedBy: ID
  updatedBy: ID
  createdBy: ID
  deletedAt: Int
  updatedAt: Int
  createdAt: Int!
}

input TodoCreateInput {
  title: String!
  remark: String
}

input TodoUpdateInput {
  title: String!
  remark: String
}`
)

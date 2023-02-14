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
  todos(current_page: Int = 1, per_page: Int = 10, q: String, sort: [TodoSortType!], filter: TodoFilterType, rand: Boolean = false): TodoResultType
}

type Mutation {
  createTodo(input: TodoCreateInput!): Todo!
  updateTodo(id: ID!, input: TodoUpdateInput!): Todo!
}

enum ObjectSortType {
  ASC
  DESC
}

type Todo {
  id: ID!
  title: String!
  age: Int
  money: Int!
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
  age: Int
  money: Int!
  remark: String
}

input TodoUpdateInput {
  title: String!
  age: Int
  money: Int!
  remark: String
}

input TodoSortType {
  id: ObjectSortType
  title: ObjectSortType
  age: ObjectSortType
  money: ObjectSortType
  remark: ObjectSortType
  deletedBy: ObjectSortType
  updatedBy: ObjectSortType
  createdBy: ObjectSortType
  deletedAt: ObjectSortType
  updatedAt: ObjectSortType
  createdAt: ObjectSortType
}

input TodoFilterType {
  AND: [TodoFilterType!]
  OR: [TodoFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  id_null: Boolean
  title: String
  title_ne: String
  title_gt: String
  title_lt: String
  title_gte: String
  title_lte: String
  title_in: [String!]
  title_like: String
  title_prefix: String
  title_suffix: String
  title_null: Boolean
  age: Int
  age_ne: Int
  age_gt: Int
  age_lt: Int
  age_gte: Int
  age_lte: Int
  age_in: [Int!]
  age_null: Boolean
  money: Int
  money_ne: Int
  money_gt: Int
  money_lt: Int
  money_gte: Int
  money_lte: Int
  money_in: [Int!]
  money_null: Boolean
  remark: String
  remark_ne: String
  remark_gt: String
  remark_lt: String
  remark_gte: String
  remark_lte: String
  remark_in: [String!]
  remark_like: String
  remark_prefix: String
  remark_suffix: String
  remark_null: Boolean
  deletedBy: ID
  deletedBy_ne: ID
  deletedBy_gt: ID
  deletedBy_lt: ID
  deletedBy_gte: ID
  deletedBy_lte: ID
  deletedBy_in: [ID!]
  deletedBy_null: Boolean
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  createdBy_null: Boolean
  deletedAt: Int
  deletedAt_ne: Int
  deletedAt_gt: Int
  deletedAt_lt: Int
  deletedAt_gte: Int
  deletedAt_lte: Int
  deletedAt_in: [Int!]
  deletedAt_null: Boolean
  updatedAt: Int
  updatedAt_ne: Int
  updatedAt_gt: Int
  updatedAt_lt: Int
  updatedAt_gte: Int
  updatedAt_lte: Int
  updatedAt_in: [Int!]
  updatedAt_null: Boolean
  createdAt: Int
  createdAt_ne: Int
  createdAt_gt: Int
  createdAt_lt: Int
  createdAt_gte: Int
  createdAt_lte: Int
  createdAt_in: [Int!]
  createdAt_null: Boolean
}

type TodoResultType {
  data: [Todo!]!
  total: Int!
  current_page: Int!
  per_page: Int!
  total_page: Int!
}`
)

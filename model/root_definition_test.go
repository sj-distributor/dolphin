package model

import (
	"strings"
	"testing"
)

// TestEnrichModelMergesNamedRootExtensions verifies that an all-skip model keeps
// its project-owned query and mutation fields without generating generic CRUD.
func TestEnrichModelMergesNamedRootExtensions(t *testing.T) {
	parsed, err := Parse(`
type WorkTask @entity @skip {
  operationKey: String
}

extend type Query {
  workTaskByKey(operationKey: String!): WorkTask
}

extend type Mutation {
  claimWork(operationKey: String!): Boolean!
}
`)
	if err != nil {
		t.Fatal(err)
	}
	if err := EnrichModelObjects(&parsed); err != nil {
		t.Fatal(err)
	}
	if err := EnrichModel(&parsed); err != nil {
		t.Fatal(err)
	}
	printed, err := PrintSchema(parsed)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(printed, "type Query {}") || !strings.Contains(printed, "workTaskByKey(operationKey: String!): WorkTask") {
		t.Fatal("named query extension was not merged into Query")
	}
	if strings.Contains(printed, "workTask(id:") || strings.Contains(printed, "workTasks(") {
		t.Fatal("all-skip entity unexpectedly generated query CRUD")
	}
	if strings.Contains(printed, "type Mutation {}") || !strings.Contains(printed, "claimWork(operationKey: String!): Boolean!") {
		t.Fatal("named mutation extension was not merged into Mutation")
	}
	if strings.Contains(printed, "createWorkTask(") || strings.Contains(printed, "updateWorkTask(") {
		t.Fatal("all-skip entity unexpectedly generated mutation CRUD")
	}
}

// TestEnrichModelKeepsGeneratedRootFieldOnExtensionCollision verifies that a
// project extension cannot create an invalid duplicate root field.
func TestEnrichModelKeepsGeneratedRootFieldOnExtensionCollision(t *testing.T) {
	parsed, err := Parse(`
type WorkTask @entity {
  operationKey: String
}

extend type Query {
  workTask(operationKey: String!): WorkTask
}
`)
	if err != nil {
		t.Fatal(err)
	}
	if err := EnrichModelObjects(&parsed); err != nil {
		t.Fatal(err)
	}
	if err := EnrichModel(&parsed); err != nil {
		t.Fatal(err)
	}
	printed, err := PrintSchema(parsed)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Count(printed, "workTask(") != 1 {
		t.Fatalf("root field collision produced duplicate definitions:\n%s", printed)
	}
	if !strings.Contains(printed, "workTask(id: ID") {
		t.Fatal("generated root field must win an extension name collision")
	}
}

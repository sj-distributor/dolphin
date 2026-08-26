package cmd

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"testing"

	"github.com/sj-distributor/dolphin/model"
)

func TestCreateUtilsFileGeneratesUniquenessSupport(t *testing.T) {
	projectDir := t.TempDir()
	if err := os.WriteFile(
		filepath.Join(projectDir, "dolphin.yml"),
		[]byte("package: example.com/generated\n"),
		0o644,
	); err != nil {
		t.Fatal(err)
	}

	installFakeGoimports(t)

	if err := createUtilsFile(projectDir); err != nil {
		t.Fatalf("createUtilsFile() failed: %v", err)
	}

	generatedPath := filepath.Join(projectDir, "utils", "uniqueness.go")
	parsed, err := parser.ParseFile(token.NewFileSet(), generatedPath, nil, 0)
	if err != nil {
		t.Fatalf("generated uniqueness support is not valid Go: %v", err)
	}

	wantFunctions := map[string]bool{
		"UniqueRecordExists":             false,
		"normalizeUniqueColumnName":      false,
		"graphQLFieldNameFromColumnName": false,
	}
	for _, declaration := range parsed.Decls {
		function, ok := declaration.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if _, required := wantFunctions[function.Name.Name]; required {
			wantFunctions[function.Name.Name] = true
		}
	}
	for name, found := range wantFunctions {
		if !found {
			t.Errorf("generated uniqueness support is missing %s", name)
		}
	}
}

func TestRunGQLGenUsesProjectRoot(t *testing.T) {
	projectDir := filepath.Join(t.TempDir(), "project")
	var gotCommand, gotDir string

	err := runGQLGenWith(projectDir, func(command, dir string) error {
		gotCommand = command
		gotDir = dir
		return nil
	})
	if err != nil {
		t.Fatalf("runGQLGenWith() failed: %v", err)
	}
	if gotCommand != "go mod tidy && go run github.com/99designs/gqlgen" {
		t.Errorf("unexpected gqlgen command %q", gotCommand)
	}
	if gotDir != projectDir {
		t.Errorf("gqlgen ran in %q, want project root %q", gotDir, projectDir)
	}
}

func TestCreateDatabaseFilesGeneratesFriendlyErrorSupport(t *testing.T) {
	projectDir := t.TempDir()
	installFakeGoimports(t)
	parsedModel, err := model.Parse("type Query { ping: String }")
	if err != nil {
		t.Fatal(err)
	}

	if err := createDatabaseFiles(
		projectDir,
		&parsedModel,
		&model.Config{Package: "example.com/generated"},
	); err != nil {
		t.Fatalf("createDatabaseFiles() failed: %v", err)
	}

	generatedPath := filepath.Join(projectDir, "gen", "database-errors.go")
	parsed, err := parser.ParseFile(token.NewFileSet(), generatedPath, nil, 0)
	if err != nil {
		t.Fatalf("generated database error support is not valid Go: %v", err)
	}

	foundFormatter := false
	for _, declaration := range parsed.Decls {
		function, ok := declaration.(*ast.FuncDecl)
		if ok && function.Name.Name == "formatDatabaseConnectionError" {
			foundFormatter = true
		}
	}
	if !foundFormatter {
		t.Fatal("generated database error support is missing formatDatabaseConnectionError")
	}
}

func installFakeGoimports(t *testing.T) {
	t.Helper()
	fakeGoPath := t.TempDir()
	binDir := filepath.Join(fakeGoPath, "bin")
	if err := os.MkdirAll(binDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(
		filepath.Join(binDir, "goimports"),
		[]byte("#!/bin/sh\nexit 0\n"),
		0o755,
	); err != nil {
		t.Fatal(err)
	}
	t.Setenv("GOPATH", fakeGoPath)
}

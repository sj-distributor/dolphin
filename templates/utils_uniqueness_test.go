package templates

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestUniquenessTemplateNormalizesColumnNamesSafely(t *testing.T) {
	fset := token.NewFileSet()
	parsed, err := parser.ParseFile(fset, "uniqueness.go", Uniqueness, 0)
	if err != nil {
		t.Fatalf("parse Uniqueness template: %v", err)
	}

	var behaviorSource bytes.Buffer
	foundFunction := false
	for _, declaration := range parsed.Decls {
		switch candidate := declaration.(type) {
		case *ast.FuncDecl:
			if candidate.Name.Name != "normalizeUniqueColumnName" {
				continue
			}
			if err := format.Node(&behaviorSource, fset, candidate); err != nil {
				t.Fatalf("format normalizeUniqueColumnName: %v", err)
			}
			behaviorSource.WriteString("\n\n")
			foundFunction = true
		case *ast.GenDecl:
			for _, spec := range candidate.Specs {
				value, ok := spec.(*ast.ValueSpec)
				if !ok || len(value.Names) != 1 || value.Names[0].Name != "uniqueColumnNamePattern" {
					continue
				}
				if err := format.Node(&behaviorSource, fset, candidate); err != nil {
					t.Fatalf("format uniqueColumnNamePattern: %v", err)
				}
				behaviorSource.WriteString("\n\n")
			}
		}
	}
	if !foundFunction {
		t.Fatal("Uniqueness template is missing normalizeUniqueColumnName")
	}

	program := fmt.Sprintf(`package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/iancoleman/strcase"
)

var (
	_ = regexp.MustCompile
	_ = strings.TrimSpace
	_ = unicode.IsUpper
	_ = strcase.ToSnake
)

%s

func main() {
	cases := map[string]string{
		"requestId": "request_id",
		"apiURL":    "api_url",
		"URL":       "url",
		"request_id": "request_id",
	}
	for input, want := range cases {
		if got := normalizeUniqueColumnName(input); got != want {
			panic(fmt.Sprintf("normalizeUniqueColumnName(%%q) = %%q, want %%q", input, got, want))
		}
	}
	for _, input := range []string{"name desc", "name;drop", "name--"} {
		if got := normalizeUniqueColumnName(input); got != "" {
			panic(fmt.Sprintf("unsafe identifier %%q normalized to %%q", input, got))
		}
	}
}
`, behaviorSource.String())

	programPath := filepath.Join(t.TempDir(), "main.go")
	if err := os.WriteFile(programPath, []byte(program), 0o644); err != nil {
		t.Fatal(err)
	}

	command := exec.Command("go", "run", programPath)
	command.Dir = ".."
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("generated normalization behavior failed: %v\n%s", err, output)
	}
}

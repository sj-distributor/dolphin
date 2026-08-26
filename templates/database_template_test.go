package templates

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

func TestDatabaseTemplateProvidesNonPanickingConnectionAPIs(t *testing.T) {
	source := strings.Split(Database, "var ShardingArray")[0]
	parsed, err := parser.ParseFile(token.NewFileSet(), "database.go", source, 0)
	if err != nil {
		t.Fatalf("parse Database template: %v", err)
	}

	want := map[string]bool{
		"OpenDBFromEnvVars": false,
		"OpenDBWithString":  false,
	}
	foundValidation := false
	foundSilentLogger := false
	foundDefaultLoggerRestore := false
	for _, declaration := range parsed.Decls {
		function, ok := declaration.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if _, required := want[function.Name.Name]; !required {
			continue
		}
		want[function.Name.Name] = true
		ast.Inspect(function.Body, func(node ast.Node) bool {
			if assignment, ok := node.(*ast.AssignStmt); ok {
				for index, left := range assignment.Lhs {
					selector, ok := left.(*ast.SelectorExpr)
					if !ok || selector.Sel.Name != "Logger" || index >= len(assignment.Rhs) {
						continue
					}
					value, ok := assignment.Rhs[index].(*ast.SelectorExpr)
					if ok && value.Sel.Name == "Default" {
						foundDefaultLoggerRestore = true
					}
				}
			}
			call, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}
			if identifier, isIdentifier := call.Fun.(*ast.Ident); isIdentifier && identifier.Name == "panic" {
				t.Errorf("%s must return errors instead of panicking", function.Name.Name)
			}
			if identifier, isIdentifier := call.Fun.(*ast.Ident); isIdentifier && identifier.Name == "validateDatabaseURL" {
				foundValidation = true
			}
			if selector, isSelector := call.Fun.(*ast.SelectorExpr); isSelector && selector.Sel.Name == "LogMode" && len(call.Args) == 1 {
				if mode, ok := call.Args[0].(*ast.SelectorExpr); ok && mode.Sel.Name == "Silent" {
					foundSilentLogger = true
				}
			}
			return true
		})
	}

	for name, found := range want {
		if !found {
			t.Errorf("Database template is missing %s", name)
		}
	}
	if !foundValidation {
		t.Error("OpenDBWithString must validate structurally incomplete database URLs before opening")
	}
	if !foundSilentLogger {
		t.Error("OpenDBWithString must silence GORM until the connection succeeds")
	}
	if !foundDefaultLoggerRestore {
		t.Error("OpenDBWithString must restore GORM's default logger after a successful non-DEBUG connection")
	}
}

func TestMainTemplateReturnsDatabaseAndCLIErrorsWithoutPanicking(t *testing.T) {
	parsed, err := parser.ParseFile(token.NewFileSet(), "main.go", Main, 0)
	if err != nil {
		t.Fatalf("parse Main template: %v", err)
	}

	for _, declaration := range parsed.Decls {
		function, ok := declaration.(*ast.FuncDecl)
		if !ok {
			continue
		}
		closeCalls := 0
		ast.Inspect(function.Body, func(node ast.Node) bool {
			call, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}
			if identifier, ok := call.Fun.(*ast.Ident); ok && identifier.Name == "panic" {
				t.Errorf("generated %s must not panic on startup errors", function.Name.Name)
			}
			if selector, ok := call.Fun.(*ast.SelectorExpr); ok && selector.Sel.Name == "NewDBFromEnvVars" {
				t.Errorf("generated %s must use the error-returning database API", function.Name.Name)
			}
			if selector, ok := call.Fun.(*ast.SelectorExpr); ok && selector.Sel.Name == "Close" {
				closeCalls++
			}
			return true
		})
		if function.Name.Name == "startServer" && closeCalls > 1 {
			t.Errorf("generated startServer closes the database %d times, want once", closeCalls)
		}
	}
}

func TestDatabaseTemplateUsesURLHostnameForPostgres(t *testing.T) {
	source := strings.Split(Database, "var ShardingArray")[0]
	parsed, err := parser.ParseFile(token.NewFileSet(), "database.go", source, 0)
	if err != nil {
		t.Fatalf("parse Database template: %v", err)
	}

	foundHostname := false
	for _, declaration := range parsed.Decls {
		function, ok := declaration.(*ast.FuncDecl)
		if !ok || function.Name.Name != "GetConnectionString" {
			continue
		}
		ast.Inspect(function.Body, func(node ast.Node) bool {
			call, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}
			if selector, ok := call.Fun.(*ast.SelectorExpr); ok && selector.Sel.Name == "Hostname" {
				foundHostname = true
			}
			return true
		})
	}
	if !foundHostname {
		t.Error("Postgres DSN generation must use URL.Hostname() so IPv6 hosts are preserved")
	}
}

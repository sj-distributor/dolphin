package templates

import (
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

func TestResultTypeTemplateAppendsStableIDTieBreaker(t *testing.T) {
	if _, err := parser.ParseFile(token.NewFileSet(), "result-type.go", ResultType, 0); err != nil {
		t.Fatalf("parse ResultType template: %v", err)
	}

	wants := []string{
		`idSortPrefix := strings.ToLower(opts.Alias + ".id ")`,
		`strings.HasPrefix(strings.ToLower(strings.TrimSpace(s)), idSortPrefix)`,
		`sorts = append(sorts, opts.Alias+".id DESC")`,
	}
	for _, want := range wants {
		if !strings.Contains(ResultType, want) {
			t.Errorf("ResultType template is missing stable pagination clause %q", want)
		}
	}
}

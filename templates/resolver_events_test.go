package templates

import (
	"strings"
	"testing"
)

func TestResolverEventsUsesValidJSONStructTags(t *testing.T) {
	if strings.Contains(ResolverEvents, `json:'`) {
		t.Fatal("resolver events template contains invalid single-quoted JSON struct tags")
	}
	for _, expected := range []string{`json:\"name\"`, `json:\"entityId\"`, `json:\"changes\"`} {
		if !strings.Contains(ResolverEvents, expected) {
			t.Fatalf("resolver events template missing %s", expected)
		}
	}
}

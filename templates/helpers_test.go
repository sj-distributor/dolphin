package templates

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunInteractiveInDirUsesRequestedDirectory(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "sentinel"), nil, 0o644); err != nil {
		t.Fatal(err)
	}

	if err := RunInteractiveInDir("test -f sentinel", dir); err != nil {
		t.Fatalf("command did not run in requested directory %q: %v", dir, err)
	}
}

func TestRunInteractiveInDirForwardsStderrAndReturnsFailure(t *testing.T) {
	originalStderr := os.Stderr
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stderr = writer
	t.Cleanup(func() {
		os.Stderr = originalStderr
		_ = reader.Close()
		_ = writer.Close()
	})

	runErr := RunInteractiveInDir("echo child-diagnostic >&2; exit 7", t.TempDir())
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	os.Stderr = originalStderr
	output, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}

	if runErr == nil {
		t.Fatal("RunInteractiveInDir() returned nil for a failing child")
	}
	if !strings.Contains(string(output), "child-diagnostic") {
		t.Fatalf("child stderr was not forwarded, got %q", output)
	}
}

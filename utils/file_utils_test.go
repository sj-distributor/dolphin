package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEnsureDir(t *testing.T) {
	t.Run("create_new", func(t *testing.T) {
		dir := filepath.Join(t.TempDir(), "new_dir")
		if err := EnsureDir(dir); err != nil {
			t.Fatalf("EnsureDir(%q) failed: %v", dir, err)
		}
		info, err := os.Stat(dir)
		if err != nil {
			t.Fatalf("directory not created: %v", err)
		}
		if !info.IsDir() {
			t.Error("EnsureDir created a file instead of a directory")
		}
	})

	t.Run("already_exists", func(t *testing.T) {
		dir := t.TempDir()
		if err := EnsureDir(dir); err != nil {
			t.Fatalf("EnsureDir(%q) on existing dir failed: %v", dir, err)
		}
	})
}

func TestFileExists(t *testing.T) {
	t.Run("exists", func(t *testing.T) {
		f := filepath.Join(t.TempDir(), "test.txt")
		if err := os.WriteFile(f, []byte("test"), 0644); err != nil {
			t.Fatal(err)
		}
		if !FileExists(f) {
			t.Errorf("FileExists(%q) = false, want true", f)
		}
	})

	t.Run("not_exists", func(t *testing.T) {
		if FileExists("/tmp/does_not_exist_12345.txt") {
			t.Error("FileExists returned true for non-existent file")
		}
	})

	t.Run("directory", func(t *testing.T) {
		dir := t.TempDir()
		// FileExists returns false for directories (by design — it checks !info.IsDir())
		if FileExists(dir) {
			t.Errorf("FileExists(%q) = true for directory, want false", dir)
		}
	})
}

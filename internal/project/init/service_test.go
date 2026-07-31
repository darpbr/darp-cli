package init

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInitializeCreatesExpectedProjectStructure(t *testing.T) {
	root := t.TempDir()
	projectDir := filepath.Join(root, "demo-project")
	if err := os.Mkdir(projectDir, 0o755); err != nil {
		t.Fatalf("mkdir project dir: %v", err)
	}

	service := NewService(NewOSFileSystem(), "# lifecycle\n")

	result, err := service.Initialize(projectDir)
	if err != nil {
		t.Fatalf("initialize: %v", err)
	}

	if result.AlreadyInitialized {
		t.Fatalf("expected fresh initialization")
	}

	assertFileContent(t, filepath.Join(projectDir, "darp.yml"), `version: "1.0"

project:
  name: "demo-project"

governance:
  lifecycle: .darp/lifecycle.md

workflows:
  default: implement

skills:
  documentation: .agents/skills/documentation
`)
	assertFileContent(t, filepath.Join(projectDir, ".darp", "lifecycle.md"), "# lifecycle\n")

	for _, directory := range darpDirectories {
		assertDirectoryExists(t, filepath.Join(projectDir, directory))
	}
}

func TestInitializeIsIdempotentAndDoesNotOverwriteFiles(t *testing.T) {
	root := t.TempDir()
	projectDir := filepath.Join(root, "demo-project")
	if err := os.Mkdir(projectDir, 0o755); err != nil {
		t.Fatalf("mkdir project dir: %v", err)
	}

	service := NewService(NewOSFileSystem(), "# first version\n")

	if _, err := service.Initialize(projectDir); err != nil {
		t.Fatalf("first initialize: %v", err)
	}

	if err := os.WriteFile(filepath.Join(projectDir, "darp.yml"), []byte("name: custom\n"), 0o644); err != nil {
		t.Fatalf("write custom darp.yml: %v", err)
	}
	if err := os.WriteFile(filepath.Join(projectDir, ".darp", "lifecycle.md"), []byte("# user changes\n"), 0o644); err != nil {
		t.Fatalf("write custom lifecycle.md: %v", err)
	}

	secondService := NewService(NewOSFileSystem(), "# second version\n")
	result, err := secondService.Initialize(projectDir)
	if err != nil {
		t.Fatalf("second initialize: %v", err)
	}

	if !result.AlreadyInitialized {
		t.Fatalf("expected already initialized result")
	}

	assertFileContent(t, filepath.Join(projectDir, "darp.yml"), "name: custom\n")
	assertFileContent(t, filepath.Join(projectDir, ".darp", "lifecycle.md"), "# user changes\n")
}

func TestInitializeRestoresMissingConfigWithoutChangingExistingDarpFiles(t *testing.T) {
	projectDir := filepath.Join(t.TempDir(), "demo-project")
	if err := os.Mkdir(projectDir, 0o755); err != nil {
		t.Fatalf("mkdir project dir: %v", err)
	}

	service := NewService(NewOSFileSystem(), "# lifecycle\n")
	if _, err := service.Initialize(projectDir); err != nil {
		t.Fatalf("first initialize: %v", err)
	}

	importantSpec := filepath.Join(projectDir, ".darp", "specs", "important.md")
	if err := os.MkdirAll(filepath.Dir(importantSpec), 0o755); err != nil {
		t.Fatalf("mkdir specs: %v", err)
	}
	if err := os.WriteFile(importantSpec, []byte("# Keep this\n"), 0o644); err != nil {
		t.Fatalf("write important spec: %v", err)
	}
	if err := os.Remove(filepath.Join(projectDir, "darp.yml")); err != nil {
		t.Fatalf("remove config: %v", err)
	}

	result, err := service.Initialize(projectDir)
	if err != nil {
		t.Fatalf("restore config: %v", err)
	}
	if result.AlreadyInitialized {
		t.Fatalf("expected repair result")
	}
	assertFileContent(t, importantSpec, "# Keep this\n")
	if _, err := os.Stat(filepath.Join(projectDir, "darp.yml")); err != nil {
		t.Fatalf("expected restored darp.yml: %v", err)
	}
}

func assertFileContent(t *testing.T, path string, want string) {
	t.Helper()

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}

	if string(content) != want {
		t.Fatalf("unexpected content for %s\nwant:\n%s\ngot:\n%s", path, want, string(content))
	}
}

func assertDirectoryExists(t *testing.T, path string) {
	t.Helper()

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat %s: %v", path, err)
	}

	if !info.IsDir() {
		t.Fatalf("expected directory at %s", path)
	}
}

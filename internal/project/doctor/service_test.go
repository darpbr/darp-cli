package doctor

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDiagnoseHealthyProject(t *testing.T) {
	root := healthyProject(t, "1.0")
	result := NewService().Diagnose(root)

	passed, warnings, errors := result.Counts()
	if passed != 7 || warnings != 0 || errors != 0 || result.ExitCode() != 0 {
		t.Fatalf("unexpected healthy result: %#v", result)
	}
}

func TestDiagnoseOlderContractWarns(t *testing.T) {
	result := NewService().Diagnose(healthyProject(t, "0.9"))
	passed, warnings, errors := result.Counts()
	if passed != 6 || warnings != 1 || errors != 0 || result.ExitCode() != 0 {
		t.Fatalf("unexpected warning result: %#v", result)
	}
}

func TestDiagnoseInvalidProjectReportsFailuresAndDoesNotWrite(t *testing.T) {
	root := healthyProject(t, "2.0")
	workflow := filepath.Join(root, ".darp", "workflows", "implement.yaml")
	if err := os.Remove(workflow); err != nil {
		t.Fatalf("remove workflow: %v", err)
	}
	before, err := projectSnapshot(root)
	if err != nil {
		t.Fatalf("snapshot before: %v", err)
	}

	result := NewService().Diagnose(root)

	after, err := projectSnapshot(root)
	if err != nil {
		t.Fatalf("snapshot after: %v", err)
	}
	if before != after {
		t.Fatalf("doctor modified the project\nbefore: %q\nafter: %q", before, after)
	}
	passed, warnings, errors := result.Counts()
	if passed != 5 || warnings != 0 || errors != 2 || result.ExitCode() != 1 {
		t.Fatalf("unexpected failing result: %#v", result)
	}
}

func TestDiagnoseRejectsConfigurationOutsideProject(t *testing.T) {
	root := healthyProject(t, "1.0")
	config := `version: "1.0"
project:
  name: demo
governance:
  lifecycle: /tmp/lifecycle.md
workflows:
  default: implement
skills:
  documentation: .agents/skills/documentation
`
	if err := os.WriteFile(filepath.Join(root, "darp.yml"), []byte(config), 0o644); err != nil {
		t.Fatalf("write config: %v", err)
	}
	result := NewService().Diagnose(root)
	if result.Checks[0].State != Fail {
		t.Fatalf("expected configuration failure, got %#v", result.Checks[0])
	}
}

func TestDiagnoseRejectsSkillTraversal(t *testing.T) {
	root := healthyProject(t, "1.0")
	workflow := filepath.Join(root, ".darp", "workflows", "implement.yaml")
	if err := os.WriteFile(workflow, []byte("name: implement\nsteps:\n  - ../outside\n"), 0o644); err != nil {
		t.Fatalf("write workflow: %v", err)
	}
	result := NewService().Diagnose(root)
	if result.Checks[2].State != Fail {
		t.Fatalf("expected workflow failure, got %#v", result.Checks[2])
	}
}

func TestDiagnoseRejectsSkillMarkdownDirectory(t *testing.T) {
	root := healthyProject(t, "1.0")
	path := filepath.Join(root, ".agents", "skills", "documentation", "SKILL.md")
	if err := os.Remove(path); err != nil {
		t.Fatalf("remove skill file: %v", err)
	}
	if err := os.Mkdir(path, 0o755); err != nil {
		t.Fatalf("create skill directory: %v", err)
	}
	result := NewService().Diagnose(root)
	if result.Checks[3].State != Fail {
		t.Fatalf("expected skills failure, got %#v", result.Checks[3])
	}
}

func healthyProject(t *testing.T, version string) string {
	t.Helper()
	root := t.TempDir()
	for _, directory := range []string{
		".darp/governance", ".darp/workflows", ".agents/skills/documentation", ".darp/templates",
	} {
		if err := os.MkdirAll(filepath.Join(root, directory), 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", directory, err)
		}
	}
	files := map[string]string{
		"darp.yml":                              "version: \"" + version + "\"\nproject:\n  name: demo\ngovernance:\n  lifecycle: .darp/lifecycle.md\nworkflows:\n  default: implement\nskills:\n  documentation: .agents/skills/documentation\n",
		".darp/lifecycle.md":                    "# Lifecycle\n",
		".darp/governance/quality-gates.md":     "# Gates\n",
		".darp/workflows/implement.yaml":        "name: implement\nsteps:\n  - documentation\n",
		".agents/skills/documentation/SKILL.md": "---\nname: documentation\ndescription: test\n---\n# Skill\n",
	}
	for path, content := range files {
		if err := os.WriteFile(filepath.Join(root, path), []byte(content), 0o644); err != nil {
			t.Fatalf("write %s: %v", path, err)
		}
	}
	return root
}

func projectSnapshot(root string) (string, error) {
	var snapshot string
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relative, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		snapshot += relative + ":"
		if !entry.IsDir() {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			snapshot += string(content)
		}
		snapshot += "\n"
		return nil
	})
	return snapshot, err
}

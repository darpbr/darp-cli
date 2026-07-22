package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunInitInitializesProject(t *testing.T) {
	projectDir := t.TempDir()
	previousDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}

	if err := os.Chdir(projectDir); err != nil {
		t.Fatalf("chdir to temp dir: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(previousDir)
	})

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := Run([]string{"init"}, &stdout, &stderr, "# lifecycle\n")
	if exitCode != 0 {
		t.Fatalf("expected success, got exit code %d with stderr %q", exitCode, stderr.String())
	}

	if stderr.Len() != 0 {
		t.Fatalf("expected empty stderr, got %q", stderr.String())
	}

	if !strings.Contains(stdout.String(), "Project initialized") {
		t.Fatalf("expected success output, got %q", stdout.String())
	}

	if _, err := os.Stat(filepath.Join(projectDir, "darp.yaml")); err != nil {
		t.Fatalf("expected darp.yaml to exist: %v", err)
	}
}

func TestRunWithoutCommandReturnsUsageError(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := Run(nil, &stdout, &stderr, "# lifecycle\n")
	if exitCode != 1 {
		t.Fatalf("expected exit code 1, got %d", exitCode)
	}

	if stdout.Len() != 0 {
		t.Fatalf("expected empty stdout, got %q", stdout.String())
	}

	if !strings.Contains(stderr.String(), "DARP CLI - Developer AI Resource Platform") {
		t.Fatalf("expected usage output, got %q", stderr.String())
	}
}

func TestRunHelpVariants(t *testing.T) {
	testCases := [][]string{
		{"help"},
		{"--help"},
		{"-h"},
	}

	for _, args := range testCases {
		t.Run(strings.Join(args, "_"), func(t *testing.T) {
			var stdout bytes.Buffer
			var stderr bytes.Buffer

			exitCode := Run(args, &stdout, &stderr, "# lifecycle\n")
			if exitCode != 0 {
				t.Fatalf("expected exit code 0, got %d", exitCode)
			}

			if stderr.Len() != 0 {
				t.Fatalf("expected empty stderr, got %q", stderr.String())
			}

			if !strings.Contains(stdout.String(), "Useful commands:") {
				t.Fatalf("expected help output, got %q", stdout.String())
			}
		})
	}
}

func TestRunVersionVariants(t *testing.T) {
	testCases := [][]string{
		{"version"},
		{"--version"},
		{"-v"},
	}

	for _, args := range testCases {
		t.Run(strings.Join(args, "_"), func(t *testing.T) {
			var stdout bytes.Buffer
			var stderr bytes.Buffer

			exitCode := RunWithVersion(args, &stdout, &stderr, "# lifecycle\n", "0.1.0")
			if exitCode != 0 {
				t.Fatalf("expected exit code 0, got %d", exitCode)
			}

			if stderr.Len() != 0 {
				t.Fatalf("expected empty stderr, got %q", stderr.String())
			}

			if strings.TrimSpace(stdout.String()) != "darp version 0.1.0" {
				t.Fatalf("unexpected version output %q", stdout.String())
			}
		})
	}
}

func TestRunUnknownCommandShowsHelp(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := Run([]string{"nope"}, &stdout, &stderr, "# lifecycle\n")
	if exitCode != 1 {
		t.Fatalf("expected exit code 1, got %d", exitCode)
	}

	if stdout.Len() != 0 {
		t.Fatalf("expected empty stdout, got %q", stdout.String())
	}

	if !strings.Contains(stderr.String(), `unknown command "nope"`) {
		t.Fatalf("expected unknown command error, got %q", stderr.String())
	}

	if !strings.Contains(stderr.String(), "Useful commands:") {
		t.Fatalf("expected help text with unknown command, got %q", stderr.String())
	}
}

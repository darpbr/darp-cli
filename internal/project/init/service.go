package init

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

const configFileName = "darp.yml"

var darpDirectories = []string{
	".darp",
	".darp/governance",
	".darp/workflows",
	".darp/templates",
	".agents/skills/documentation",
}

// FileSystem abstracts filesystem interactions used by the initialization service.
type FileSystem interface {
	Exists(path string) (bool, error)
	MkdirAll(path string) error
	WriteFile(path string, data []byte) error
	Base(path string) string
}

// Result captures the visible outcome of the initialization flow.
type Result struct {
	AlreadyInitialized bool
	Messages           []string
}

// Service initializes a directory as a DARP project.
type Service struct {
	fs               FileSystem
	lifecycleContent string
}

// NewService creates a new initialization service.
func NewService(fs FileSystem, lifecycleContent string) Service {
	return Service{
		fs:               fs,
		lifecycleContent: lifecycleContent,
	}
}

// Initialize applies the project bootstrap structure to the provided directory.
func (s Service) Initialize(root string) (Result, error) {
	messages := []string{"✔ Initializing project"}

	projectName := strings.TrimSpace(s.fs.Base(root))
	if projectName == "" || projectName == "." || projectName == string(filepath.Separator) {
		return Result{}, errors.New("could not derive project name from current directory")
	}

	changed := false
	configPath := filepath.Join(root, configFileName)
	created, err := s.writeFileIfMissing(configPath, []byte(renderConfig(projectName)))
	if err != nil {
		return Result{}, fmt.Errorf("create darp.yml: %w", err)
	}
	if created {
		messages = append(messages, "✔ Creating darp.yml")
		changed = true
	}

	createdStructure := false
	for _, directory := range darpDirectories {
		created, err := s.createDirectoryIfMissing(filepath.Join(root, directory))
		if err != nil {
			return Result{}, fmt.Errorf("create %s: %w", directory, err)
		}
		createdStructure = createdStructure || created
		changed = changed || created
	}
	if createdStructure {
		messages = append(messages, "✔ Creating .darp structure")
	}

	contracts := []struct{ path, content string }{
		{".darp/lifecycle.md", s.lifecycleContent},
		{".darp/governance/quality-gates.md", "# Quality Gates\n"},
		{".agents/skills/documentation/SKILL.md", "---\nname: documentation\ndescription: Keep project documentation aligned with the implementation.\n---\n\n# Documentation Skill\n"},
		{".darp/workflows/implement.yaml", "name: implement\n\nsteps:\n  - documentation\n"},
	}
	createdContracts := false
	for _, contract := range contracts {
		created, err := s.writeFileIfMissing(filepath.Join(root, contract.path), []byte(contract.content))
		if err != nil {
			return Result{}, fmt.Errorf("write %s: %w", contract.path, err)
		}
		createdContracts = createdContracts || created
		changed = changed || created
	}
	if createdContracts {
		messages = append(messages, "✔ Creating DARP contracts")
	}

	if !changed {
		messages = append(messages, "✔ Project already initialized")
		return Result{AlreadyInitialized: true, Messages: messages}, nil
	}

	messages = append(messages, "✔ Project initialized")

	return Result{
		Messages: messages,
	}, nil
}

func (s Service) createDirectoryIfMissing(path string) (bool, error) {
	exists, err := s.fs.Exists(path)
	if err != nil {
		return false, err
	}
	if exists {
		return false, nil
	}
	return true, s.fs.MkdirAll(path)
}

func (s Service) writeFileIfMissing(path string, data []byte) (bool, error) {
	exists, err := s.fs.Exists(path)
	if err != nil {
		return false, err
	}
	if exists {
		return false, nil
	}
	return true, s.fs.WriteFile(path, data)
}

func renderConfig(projectName string) string {
	return fmt.Sprintf(`version: "1.0"

project:
  name: %q

governance:
  lifecycle: .darp/lifecycle.md

workflows:
  default: implement

skills:
  documentation: .agents/skills/documentation
`, projectName)
}

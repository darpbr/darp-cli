package init

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

const configFileName = "darp.yaml"

var darpDirectories = []string{
	".darp",
	".darp/specs",
	".darp/tasks",
	".darp/docs",
	".darp/templates",
	".darp/examples",
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

	alreadyInitialized, err := s.isAlreadyInitialized(root)
	if err != nil {
		return Result{}, err
	}

	if alreadyInitialized {
		messages = append(messages, "✔ Project already initialized")
		return Result{
			AlreadyInitialized: true,
			Messages:           messages,
		}, nil
	}

	projectName := strings.TrimSpace(s.fs.Base(root))
	if projectName == "" || projectName == "." || projectName == string(filepath.Separator) {
		return Result{}, errors.New("could not derive project name from current directory")
	}

	messages = append(messages, "✔ Creating darp.yaml")
	if err := s.fs.WriteFile(filepath.Join(root, configFileName), []byte(renderConfig(projectName))); err != nil {
		return Result{}, fmt.Errorf("create darp.yaml: %w", err)
	}

	messages = append(messages, "✔ Creating .darp structure")
	for _, directory := range darpDirectories {
		if err := s.fs.MkdirAll(filepath.Join(root, directory)); err != nil {
			return Result{}, fmt.Errorf("create %s: %w", directory, err)
		}
	}

	messages = append(messages, "✔ Copying lifecycle.md")
	if err := s.fs.WriteFile(filepath.Join(root, ".darp", "lifecycle.md"), []byte(s.lifecycleContent)); err != nil {
		return Result{}, fmt.Errorf("write lifecycle.md: %w", err)
	}

	messages = append(messages, "✔ Project initialized")

	return Result{
		Messages: messages,
	}, nil
}

func (s Service) isAlreadyInitialized(root string) (bool, error) {
	for _, candidate := range []string{
		filepath.Join(root, configFileName),
		filepath.Join(root, ".darp"),
	} {
		exists, err := s.fs.Exists(candidate)
		if err != nil {
			return false, err
		}
		if exists {
			return true, nil
		}
	}

	return false, nil
}

func renderConfig(projectName string) string {
	return fmt.Sprintf("name: %s\nversion: 1\nspecVersion: 0.1\n", projectName)
}

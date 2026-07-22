package cli

import (
	"fmt"
	"io"
	"strings"

	projectinit "github.com/darpbr/darp-cli/internal/project/init"
)

const (
	cliName        = "DARP CLI"
	cliDescription = "Developer AI Resource Platform"
	defaultVersion = "dev"
)

type commandHelp struct {
	Name    string
	Summary string
}

var commandCatalog = []commandHelp{
	{Name: "init", Summary: "Initialize a DARP project"},
	{Name: "help", Summary: "Show CLI help"},
	{Name: "version", Summary: "Show CLI version"},
}

// Run executes the CLI command flow and returns a process exit code.
func Run(args []string, stdout io.Writer, stderr io.Writer, lifecycleContent string) int {
	return RunWithVersion(args, stdout, stderr, lifecycleContent, defaultVersion)
}

// RunWithVersion executes the CLI command flow with the provided version string.
func RunWithVersion(args []string, stdout io.Writer, stderr io.Writer, lifecycleContent string, version string) int {
	if len(args) == 0 {
		fmt.Fprint(stderr, renderHelp())
		return 1
	}

	switch args[0] {
	case "help", "-h", "--help":
		fmt.Fprint(stdout, renderHelp())
		return 0
	case "version", "-v", "--version":
		fmt.Fprintln(stdout, renderVersion(version))
		return 0
	case "init":
		service := projectinit.NewService(projectinit.NewOSFileSystem(), lifecycleContent)
		result, err := service.Initialize(".")
		if err != nil {
			fmt.Fprintf(stderr, "error: %v\n", err)
			return 1
		}

		for _, message := range result.Messages {
			fmt.Fprintln(stdout, message)
		}

		return 0
	default:
		fmt.Fprintf(stderr, "unknown command %q\n", args[0])
		fmt.Fprintln(stderr, "")
		fmt.Fprint(stderr, renderHelp())
		return 1
	}
}

func renderHelp() string {
	var builder strings.Builder

	builder.WriteString(cliName)
	builder.WriteString(" - ")
	builder.WriteString(cliDescription)
	builder.WriteString("\n\n")
	builder.WriteString("Manage reusable AI assets and DARP project structure from the command line.\n\n")
	builder.WriteString("Usage:\n")
	builder.WriteString("  darp <command> [arguments]\n\n")
	builder.WriteString("Useful commands:\n")

	for _, command := range commandCatalog {
		builder.WriteString(fmt.Sprintf("  %-8s %s\n", command.Name, command.Summary))
	}

	builder.WriteString("\nOptions:\n")
	builder.WriteString("  -h, --help       Show help\n")
	builder.WriteString("  -v, --version    Show version\n")

	return builder.String()
}

func renderVersion(version string) string {
	return fmt.Sprintf("darp version %s", version)
}

package main

import (
	_ "embed"
	"os"

	"github.com/darpbr/darp-cli/internal/cli"
)

//go:embed .darp/lifecycle.md
var lifecycleContent string

var version = "dev"

func main() {
	os.Exit(cli.RunWithVersion(os.Args[1:], os.Stdout, os.Stderr, lifecycleContent, version))
}

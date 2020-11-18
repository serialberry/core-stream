package main

import (
	cmd "github.com/serialberry/core-stream/internal/cmd/cli"
	"github.com/serialberry/core-stream/internal/cmd/version"
)

func main() {
	cli := cmd.NewCli()
	cli.Add(version.NewVersion("0.0.1"))
	cli.Run()
}

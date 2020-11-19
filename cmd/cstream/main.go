package main

import (
	"github.com/serialberry/core-stream/internal/cmd"
	"github.com/spf13/cobra"
)

var (
	cli *cobra.Command
)

func init() {
	cli = &cobra.Command{
		Use:   "cstream <command> <subcommand> [flags]",
		Short: "Core Stream CLI",
		Long:  "Command line tool for working with connected camera feeds",
	}
}

func main() {
	cli.AddCommand(cmd.VersionCommand("0.0.1"))
	cli.AddCommand(cmd.CaptureCommand())
	cli.Execute()
}

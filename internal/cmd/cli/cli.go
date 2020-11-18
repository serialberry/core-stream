package cli

import (
	"github.com/spf13/cobra"
)

type Cli struct {
	Command *cobra.Command
}

func NewCli() *Cli {
	return &Cli{
		Command: &cobra.Command{
			Use:   "cstream <command> <subcommand> [flags]",
			Short: "Core Stream CLI",
			Long:  "Command line tool for working with connected camera feeds",
		}}
}

// add child command to cli
func (r *Cli) Add(command *cobra.Command) {
	r.Command.AddCommand(command)
}

func (r *Cli) Run() error {
	return r.Command.Execute()
}

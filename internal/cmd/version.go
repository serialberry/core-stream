package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func VersionCommand(version string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "cli version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}
}

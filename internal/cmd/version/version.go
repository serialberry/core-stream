package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewVersion(version string) *cobra.Command {
	return &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}
}

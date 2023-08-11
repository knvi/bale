package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Bale",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Bale v0.0.1")
		},
	}

	return cmd
}
package root

import (
	"github.com/knvi/bale/pkg/cmd/version"
	"github.com/spf13/cobra"
)

func RootCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "bale <command> <subcommand> [flags]",
		Short: "Bale is a tool for templating anything",
	}

	cmd.AddCommand(version.CmdVersion())

	return cmd, nil
}

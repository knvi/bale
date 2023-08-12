package root

import (
	"github.com/knvi/bale/internal/cmd/create"
	"github.com/knvi/bale/internal/cmd/deploy"
	"github.com/knvi/bale/internal/cmd/version"
	"github.com/spf13/cobra"
)

func RootCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "bale <command> <subcommand> [flags]",
		Short: "Bale is a tool for templating anything",
	}

	cmd.AddCommand(version.CmdVersion())
	cmd.AddCommand(deploy.CmdDeploy())
	cmd.AddCommand(create.CmdCreate())

	return cmd, nil
}

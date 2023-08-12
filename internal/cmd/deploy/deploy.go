package deploy

import (
	"fmt"

	"github.com/knvi/bale/internal/templater"
	"github.com/spf13/cobra"
)

type DeployOpts struct {
	Name string
}

func CmdDeploy() *cobra.Command {
	opts := &DeployOpts{}

	cmd := &cobra.Command{
		Use:   "deploy [<name>]",
		Short: "Deploy a template",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				opts.Name = args[0]
			}

			if opts.Name == "" {
				fmt.Println("Name is required")
				return
			}

			templater.DeployTemplate(opts.Name)
		},
	}

	return cmd
}
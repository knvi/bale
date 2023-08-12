package create

import (
	"fmt"

	"github.com/knvi/bale/internal/templater"
	"github.com/spf13/cobra"
)

type CreateOpts struct {
	Name string
}

func CmdCreate() *cobra.Command {
	opts := &CreateOpts{}

	cmd := &cobra.Command{
		Use:   "create [<name>]",
		Short: "Create a template",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				opts.Name = args[0]
			}

			if opts.Name == "" {
				fmt.Println("Name is required")
				return
			}

			templater.CreateTemplate(opts.Name)
		},
	}

	return cmd
}
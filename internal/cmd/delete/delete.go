package delete

import (
	"fmt"

	"github.com/knvi/bale/internal/templater"
	"github.com/spf13/cobra"
)

type CreateOpts struct {
	Name string
}

func CmdDelete() *cobra.Command {
	opts := &CreateOpts{}

	cmd := &cobra.Command{
		Use:   "delete [<name>]",
		Short: "Delete a template",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				opts.Name = args[0]
			}

			if opts.Name == "" {
				fmt.Println("Name is required")
				return
			}

			templater.DeleteTemplate(opts.Name)
		},
	}

	return cmd
}
package create

import (
	"fmt"

	"github.com/knvi/bale/internal/templater"
	"github.com/spf13/cobra"
)

type CreateOpts struct {
	Name string
	FilesFlag bool
	Files []string
}

func CmdCreate() *cobra.Command {
	opts := &CreateOpts{}

	cmd := &cobra.Command{
		Use:   "create <name>",
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
			
			if opts.FilesFlag {
				opts.Files = args[1:]

				if len(opts.Files) == 0 {
					fmt.Println("Files are required")
					return
				}

				options := &templater.CreateOpts{
					Name: opts.Name,
					Files: opts.Files,
				}

				templater.CreateTemplate(options)

				return
			}

			options := &templater.CreateOpts{
				Name: opts.Name,
			}

			templater.CreateTemplate(options)
		},
	}

	cmd.Flags().BoolVarP(&opts.FilesFlag, "files", "f", false, "Add files to template")

	return cmd
}
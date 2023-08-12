package list

import (
	"github.com/knvi/bale/internal/templater"
	"github.com/spf13/cobra"
)

func CmdList() *cobra.Command{
	cmd := &cobra.Command{
		Use: "list",
		Short: "List templates",
		Run: func(cmd *cobra.Command, args []string) {
			tmpls, err := templater.ListTemplates()
			if err != nil {
				panic(err)
			}

			for _, tmpl := range tmpls {
				cmd.Printf("- template %s located at %s\n", tmpl.Name, tmpl.Path)
			}
		},
	}

	return cmd
}
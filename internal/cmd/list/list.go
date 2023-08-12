package list

import (
	"github.com/knvi/bale/internal/templater"
	"github.com/spf13/cobra"
)

func formatFileTemplate(files []string) string {
	var filesStr string

	for _, file := range files {
		filesStr += file + ", "
	}

	return filesStr
}

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
				if tmpl.Path != "" {
					// dir template
					cmd.Printf("- template \"%s\" located at %s\n", tmpl.Name, tmpl.Path)
				} else {
					// file template
					cmd.Printf("- template \"%s\" with files %v\n", tmpl.Name, formatFileTemplate(tmpl.Files))
				}
			}
		},
	}

	return cmd
}
package list

import (
	"github.com/gookit/color"
	"github.com/knvi/bale/internal/templater"
	"github.com/spf13/cobra"
)

func formatFileTemplate(files []string) string {
	var filesStr string
	lGreen := color.FgLightGreen.Render

	for _, file := range files {
		filesStr += lGreen(file) + " "
	}

	return filesStr
}

func CmdList() *cobra.Command{
	red := color.FgRed.Render
	lGreen := color.FgLightGreen.Render

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
					cmd.Printf("- template %s located at %s\n", red(tmpl.Name), lGreen(tmpl.Path))
				} else {
					// file template
					cmd.Printf("- template %s with files %v\n", red(tmpl.Name), formatFileTemplate(tmpl.Files))
				}
			}
		},
	}

	return cmd
}
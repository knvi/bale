package templater

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gookit/color"
	"github.com/knvi/bale/internal/config"
	"github.com/knvi/bale/internal/osutil"
)

// CODE that does stuff Sooooooo nicellyyyyyy

// CREATE A TEMPLATE soooooo nicellyyyyyy
type CreateOpts struct {
	Name  string
	Files []string
}

func CreateTemplate(opts *CreateOpts) {
	// get current path
	path, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting current path: %s\n", err)
		os.Exit(1)
	}

	if opts.Files == nil {
		// no files means that the template is a dir template

		// create template
		tmpl := config.Template{
			Path: path,
			Name: opts.Name,
		}

		config, err := config.LoadConfig()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading config: %s\n", err)
			os.Exit(1)
		}

		config.AddTemplate(&tmpl)

		if err := config.Save(); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving config: %s\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "🚀 Created template %s!\n", opts.Name)

	} else {
		// files means that the template is a file template

		// convert opts.Files to absolute paths
		tFiles := make([]string, len(opts.Files))
		for i, file := range opts.Files {
			tFiles[i] = filepath.Join(path, file)
		}

		// create template
		tmpl := config.Template{
			Name: opts.Name,
			Files: tFiles,
		}

		config, err := config.LoadConfig()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading config: %s\n", err)
			os.Exit(1)
		}

		config.AddTemplate(&tmpl)

		if err := config.Save(); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving config: %s\n", err)
			os.Exit(1)
		}

		color.LightGreen.Printf("🚀 Created template %s!\n", opts.Name)
	}
}

// DEPLOY A TEMPLATE soooooo nicellyyyyyy
func DeployTemplate(name string) {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	tmpl := config.GetTemplate(name)
	if tmpl == nil {
		fmt.Println("Template not found")
		return
	}

	path, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting current path: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Deploying template %s", tmpl.Name)

	if tmpl.Files == nil {
		// dir template
		osutil.CopyFiles(tmpl.Path, path, false)
	} else {
		// file template
		for _, file := range tmpl.Files {
			err := osutil.CopyFile(file, path, false)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error copying file %s: %s\n", file, err)
				os.Exit(1)
			}
		}
	}

	color.LightGreen.Println("🚀 Done!")

}

// LIST TEMPLATES soooooo nicellyyyyyy
func ListTemplates() ([]*config.Template, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	return cfg.GetTemplates(), nil
}

// DELETE A TEMPLATE soooooo nicellyyyyyy
func DeleteTemplate(name string) {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	err = config.DeleteTemplate(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting template: %s\n", err)
		os.Exit(1)
	}

	if err := config.Save(); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving config: %s\n", err)
		os.Exit(1)
	}

	color.LightGreen.Printf("🚀 Deleted template %s!\n", name)
}

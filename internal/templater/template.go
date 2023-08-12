package templater

import (
	"fmt"
	"os"

	"github.com/knvi/bale/internal/config"
	"github.com/knvi/bale/internal/osutil"
)

// CODE that does stuff Sooooooo nicellyyyyyy

// CREATE A TEMPLATE soooooo nicellyyyyyy
func CreateTemplate(name string) {
	// get current path
	path, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting current path: %s\n", err)
		os.Exit(1)
	}

	// create template
	tmpl := config.Template{
		Path: path,
		Name: name,
	}

	config, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %s\n", err)
		os.Exit(1)
	}

	config.AddTemplate(tmpl)

	if err := config.Save(); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving config: %s\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "Created template %s\n", name)
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

	fmt.Println("Deploying template", tmpl.Name)
	fmt.Println("From", tmpl.Path)
	fmt.Println("To", path)

	osutil.CopyFiles(tmpl.Path, path, false)

}
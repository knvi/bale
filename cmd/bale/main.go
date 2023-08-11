package main

import (
	"fmt"
	"os"

	"github.com/knvi/bale/pkg/cmd/root"
)

func main() {
	code := run()
	os.Exit(code)
}

func run() int {
	root, err := root.RootCmd()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}
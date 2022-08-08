package main

import (
	"log"

	"github.com/spf13/cobra/doc"
	"github.com/sunny0826/kubecm/cmd"
)

func main() {
	kubecm := cmd.NewBaseCommand().CobraCmd()
	err := doc.GenMarkdownTree(kubecm, "./docs/en-us/cli/")
	if err != nil {
		log.Fatal(err)
	}
}

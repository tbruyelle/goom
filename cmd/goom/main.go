package main

import (
	"fmt"
	"github.com/tbruyelle/goom"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: goon <command>\n")
		return
	}
	g, err := goom.New()
	if err != nil {
		fmt.Printf("Error while creating goom container: %v\n", err)
		return
	}
	val, err := g.Get(os.Args[1])
	if err != nil {
		fmt.Printf("Error while retrieving key %s: %v\n", os.Args[1], err)
		return
	}
	fmt.Printf("%s has been copied to clipboard.\n", val)
}

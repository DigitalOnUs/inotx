package main

/*
   Behavior: reads an HCL input file
   and produces a json output with the proper syntax
   to be consume by another API adding the consul servers
   required for the architecture
*/

import (
	"flag"
	"fmt"
	"os"

	"github.com/DigitalOnUs/inotx/config"
	"github.com/fatih/color"
)

func main() {
	os.Exit(exec())
}

func exec() int {
	// add restrictions
	var standaloneDB bool

	options := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	options.BoolVar(&standaloneDB, "alonedb", false,
		"if a db servie is found its items will be allocated in a standalone client box")
	options.Parse(os.Args[1:])
	args := options.Args()

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, color.RedString(
			"Please provide a path to file to analyze.\n\n"))
	}

	color.Green("*** inotx @DoU alpha ***")

	for _, cfg := range args {
		root, err := config.ParseFile(cfg)
		if err != nil {
			fmt.Fprint(os.Stderr, color.RedString(
				"Unable to load input file ", err))
			continue
		}

		config.AddConsul(root)
	}

	return 0
}

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
	var outputFormat string
	/* defaults same json -> json or hcl -> hcl
	-> json
	-> hcl
	*/
	options := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	options.BoolVar(&standaloneDB, "alonedb", false,
		"if a db service is found its items will be allocated in a standalone client box")
	// output format
	options.StringVar(&outputFormat, "format", "",
		"output formats [json], [hcl] if no format is specified the same format of the input will be used")

	options.Parse(os.Args[1:])
	args := options.Args()

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, color.RedString(
			"Please provide a path to file to analyze.\n"))
		return 1
	}

	color.Green("*** inotx @DoU alpha ***")

	for _, cfg := range args {
		root, err := config.ParseFile(cfg)
		if err != nil {
			fmt.Fprint(os.Stderr, color.RedString(
				"Unable to load input file %s", err))
			continue
		}

		newDoc, err := config.AddConsul(root)
		if err != nil {
			fmt.Fprintf(os.Stderr, color.RedString(
				"Unable to update spec %s", err))
			continue
		}

		filename, err := config.WriteFile(cfg, outputFormat, newDoc)
		if err != nil {
			fmt.Fprintf(os.Stderr, color.RedString(
				"Error writing outcome file %s", err))
			continue
		}

		color.Green("Output spec %s", filename)

	}

	return 0
}

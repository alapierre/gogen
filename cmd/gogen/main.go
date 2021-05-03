package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/alapierre/gogen/generator"
	"github.com/alapierre/gogen/utils"
	"os"
)

func main() {

	parser := argparse.NewParser("gogen", "Project generator for Go")

	module := parser.String("m", "module", &argparse.Options{Required: true, Help: "Module name"})

	//traceEnabled := parser.Flag("d", "debug", &argparse.Options{Required: false, Help: "Enable trace requests info"})

	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	name, err := generator.ExtractProjectName(*module)

	if err != nil {
		fmt.Printf("module name in wrong format %v\n", err)
		os.Exit(1)
	}

	err = generator.CreateProjectStructure(utils.ToSnakeCase(name))
	if err != nil {
		fmt.Printf("Can;t create folder %s %v\n", name, err)
		os.Exit(1)
	}

	err = generator.GenMain(utils.ToSnakeCase(name))
	if err != nil {
		fmt.Printf("Can't create main %s %v\n", name, err)
		os.Exit(1)
	}

}

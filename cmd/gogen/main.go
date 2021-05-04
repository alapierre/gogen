package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/alapierre/gogen/generator"
	"io/ioutil"
	"os"
	"os/exec"
)

//go:embed res/Dockerfile
var dockerDefaultTemplate string

//go:embed res/groups
var groupsTemplate string

//go:embed res/passwd
var passTemplate string

//go:embed res/Makefile
var makefileTemplate string

func main() {

	parser := argparse.NewParser("gogen", "Project generator for Go")

	module := parser.String("m", "module", &argparse.Options{Required: true, Help: "Module name"})
	generateDocker := parser.Flag("d", "docker", &argparse.Options{Required: false, Help: "Enable Dockerfile generation"})

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

	project := generator.Project{
		Module:       *module,
		OriginalName: name,
		Name:         name,
		Docker:       nil,
	}

	if *generateDocker {
		project.Docker = &generator.Docker{
			Maintainer: "Adrian Lapierre <al@alapierre.io>",
			Expose:     "9098",
		}
	}

	err = generator.CreateProjectStructure(project.Name)
	if err != nil {
		fmt.Printf("Can;t create folder %s %v\n", name, err)
		os.Exit(1)
	}

	err = generator.GenMain(project.Name)
	if err != nil {
		fmt.Printf("Can't create main %s %v\n", name, err)
		os.Exit(1)
	}

	if *generateDocker {
		docker(&project)
	}

	env(&project)
	makefile(&project)
	err = initMod(&project)
	if err != nil {
		fmt.Printf("Can't init go module %s %v\n", name, err)
		os.Exit(1)
	}
}

func makefile(project *generator.Project) {
	f, err := os.Create(project.Name + "/Makefile")
	if err != nil {
		fmt.Printf("Can't create Makefile %s %v\n", project.Name, err)
		return
	}

	//goland:noinspection ALL
	defer f.Close()

	err = generator.FileGenerator(*project, &makefileTemplate, f)
	if err != nil {
		fmt.Printf("Can't write to Makefile %s %v\n", project.Name, err)
		return
	}
}

func env(project *generator.Project) {
	_, err := os.Create(project.Name + "/.env")
	if err != nil {
		fmt.Printf("Can't create .env for %s %v\n", project.Name, err)
		return
	}
}

func docker(project *generator.Project) {

	f, err := os.Create(project.Name + "/Dockerfile")
	if err != nil {
		fmt.Printf("Can't create Dockerfile %s %v\n", project.Name, err)
		return
	}

	//goland:noinspection ALL
	defer f.Close()

	err = generator.FileGenerator(*project, &dockerDefaultTemplate, f)
	if err != nil {
		fmt.Printf("Can't write to dockerfile %s %v\n", project.Name, err)
		return
	}

	err = CopyDockerResources(project.Name)
	if err != nil {
		fmt.Printf("Can't copy docker resurces")
	}
}

func CopyDockerResources(projectName string) error {

	err := os.MkdirAll(projectName+"/resources", os.ModePerm)

	if err != nil {
		return fmt.Errorf("can't create resources dir for project name %s  %v", projectName, err)
	}

	err = ioutil.WriteFile(projectName+"/resources/group", []byte(groupsTemplate), 0644)

	if err != nil {
		return fmt.Errorf("can't save groups file for project name %s  %v", projectName, err)
	}

	err = ioutil.WriteFile(projectName+"/resources/password", []byte(passTemplate), 0644)

	if err != nil {
		return fmt.Errorf("can't save password file for project name %s  %v", projectName, err)
	}
	return nil
}

func initMod(project *generator.Project) error {

	cmdStr := "cd " + project.Name + " && go mod init " + project.Module
	cmd := exec.Command("sh", "-c", cmdStr)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	_, err := cmd.Output()

	if err != nil {
		return fmt.Errorf("initMod: sh -c %s => err:%v", cmdStr, err.Error()+" , "+stderr.String())
	}

	return nil
}

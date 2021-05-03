package generator

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"net/url"
	"os"
	"path"
)

func MakeProjectFolder(modPath string) error {

	project, err := ExtractProjectName(modPath)

	if err != nil {
		return fmt.Errorf("can't create project folder for %s  %v", modPath, err)
	}

	err = os.MkdirAll(project, os.ModePerm)

	if err != nil {
		return err
	}
	return nil
}

func CreateProjectStructure(projectName string) error {

	err := os.MkdirAll(projectName+"/cmd/"+projectName, os.ModePerm)

	if err != nil {
		return fmt.Errorf("can't create cmd dir for project name %s  %v", projectName, err)
	}

	err = os.MkdirAll(projectName+"/service", os.ModePerm)

	if err != nil {
		return fmt.Errorf("can't create service dir for project name %s  %v", projectName, err)
	}

	err = os.MkdirAll(projectName+"/transport/http", os.ModePerm)

	if err != nil {
		return fmt.Errorf("can't create /transport/http dir for project name %s  %v", projectName, err)
	}

	return nil

}

func ExtractProjectName(modPath string) (string, error) {
	myUrl, err := url.Parse(modPath)
	if err != nil {
		return "", fmt.Errorf("can't parse mod path %v", err)
	}
	return path.Base(myUrl.Path), nil
}

func GenMain(projectName string) error {

	f := jen.NewFile("main")
	f.Func().Id("main").Params().Block(
		jen.Qual("fmt", "Println").Call(jen.Lit("Hello, world")),
	)

	err := f.Save(projectName + "/cmd/" + projectName + "/main.go")
	if err != nil {
		return fmt.Errorf("can't save main.go %v", err)
	}
	return nil
}

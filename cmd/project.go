package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

type Project struct {
	PkgName         string
	AbsolutePath    string
	ApplicationName string
}

func (p *Project) Create() error {
	// check if AbsolutePath exists
	if _, err := os.Stat(p.AbsolutePath); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(p.AbsolutePath, 0754); err != nil {
			return err
		}
	}
	// create config/application.yaml
	if _, err := os.Stat(fmt.Sprintf("%s/config", p.AbsolutePath)); os.IsNotExist(err) {
		if err := os.Mkdir(fmt.Sprintf("%s/config", p.AbsolutePath), 0751); err != nil {
			return err
		}
	}

	if _, err := os.Stat(fmt.Sprintf("%s/config/application.yaml", p.AbsolutePath)); os.IsExist(err) {
		if err := os.Rename(fmt.Sprintf("%s/config/application.yaml", p.AbsolutePath),
			fmt.Sprintf("%s/config/application-old.yaml", p.AbsolutePath)); err != nil {
			return err
		}
	}

	appFile, err := os.Create(fmt.Sprintf("%s/config/application.yaml", p.AbsolutePath))
	if err != nil {
		return err
	}
	defer appFile.Close()

	appTemplate := template.Must(template.New("application").Parse(string(ApplicationTemplate())))
	err = appTemplate.Execute(appFile, p)
	if err != nil {
		return err
	}
	// import pkg
	err = exec.Command("go", "get", p.PkgName).Run()
	if err != nil {
		fmt.Println("warn: failed to go get", p.PkgName)
	}
	// create config/application.go
	goFile, err := os.Create(fmt.Sprintf("%s/config/application.go", p.AbsolutePath))
	if err != nil {
		return err
	}
	defer goFile.Close()

	goTemplate := template.Must(template.New("go").Parse(string(ApplicationGoTemplate())))
	err = goTemplate.Execute(goFile, p)
	if err != nil {
		return err
	}
	return nil
}

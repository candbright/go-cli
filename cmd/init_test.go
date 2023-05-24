package cmd

import (
	"os"
	"testing"
)

func getProject() *Project {
	wd, _ := os.Getwd()
	return &Project{
		PkgName:         "github.com/candbright/go-core",
		AbsolutePath:    wd,
		ApplicationName: "test",
	}
}

func TestProject_Create(t *testing.T) {
	err := getProject().Create()
	if err != nil {
		t.Fatal(err)
	}
}

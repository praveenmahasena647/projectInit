package model

import (
	"fmt"
	"os"
	"os/exec"
)

func (m *Model) Exc() error {
	var currentDir, dirErr = os.Getwd() // or I could do something like os.Getenv("GOROOT") for this case
	if dirErr != nil {
		return fmt.Errorf("error during getting path")
	}
	if err := chDir(); err != nil {
		return fmt.Errorf("error during chDir")
	}
	if err := m.packageInit(); err != nil {
		return fmt.Errorf("Error on packageInit")
	}
	if err := m.getHTTPpackage(); err != nil {
		return fmt.Errorf("Error on Getting workDir")
	}
	if err := m.getDBdriver(); err != nil {
		return fmt.Errorf("Error on Getting workDir")
	}
	if err := m.generateTemplate(currentDir); err != nil {
		return fmt.Errorf("Error on Getting workDir")
	}
	if err := m.gitInit(); err != nil {
		return fmt.Errorf("Error on Getting workDir")
	}
	return nil
}

func (m *Model) packageInit() error {
	return exec.Command("go", "mod", "init", m.ProjectName).Run()
}

func (m *Model) getHTTPpackage() error {
	switch m.HTTPpackage {
	case "Standered":
		return nil
	case "fastHTTP":
		return exec.Command(
			"go", "get", "-u", "github.com/qiangxue/fasthttp-routing",
			"&&",
			"go", "get", "-u", "github.com/valyala/fasthttp").Run()
	}
	return exec.Command("go", "get", "-u", m.HTTPpackage).Run()
}

func (m *Model) getDBdriver() error {
	switch m.DBdriver {
	case "gorm-postgres":
		return exec.Command(
			"go", "get", "-u", "gorm.io/gorm",
			"&&",
			"go", "get", "-u", "gorm.io/driver/postgres",
		).Run()
	case "gorm-MySQL":
		return exec.Command(
			"go", "get", "-u", "gorm.io/gorm",
			"&&",
			"go", "get", "-u", "gorm.io/driver/mysql",
		).Run()
	}
	return exec.Command("go", "get", "-u", m.DBdriver).Run()
}

func (m *Model) generateTemplate(root string) error {
	println(root)
	return nil
}

func (m *Model) gitInit() error {
	return exec.Command("git", "init").Run()
}

func chDir() error {
	var wd, wdErr = os.Getwd()
	if wdErr != nil {
		return wdErr
	}
	return os.Chdir(wd)
}

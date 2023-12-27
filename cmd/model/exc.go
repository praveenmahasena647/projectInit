package model

import (
	"fmt"
	"os"
	"os/exec"
)

func (m *Model) Exc() error {
	if err := chDir(); err != nil {
		println(1)
		return fmt.Errorf("error during chDir")
	}
	if err := m.projectInit(); err != nil {
		println(2)
		return fmt.Errorf("Error on Getting workDir")
	}
	if err := m.getHTTPpackage(); err != nil {
		println(3)
		return fmt.Errorf("Error on Getting workDir")
	}
	if err := m.getDBdriver(); err != nil {
		println(4)
		return fmt.Errorf("Error on Getting workDir")
	}
	if err := m.generateTemplate(); err != nil {
		println(5)
		return fmt.Errorf("Error on Getting workDir")
	}
	if err := m.gitInit(); err != nil {
		println(6)
		return fmt.Errorf("Error on Getting workDir")
	}
	return nil
}

func (m *Model) projectInit() error {
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

func (m *Model) generateTemplate() error {
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

package model

import (
	"errors"
	"os"
	"os/exec"
)

func (m *Model) Exc() error {
	if err := chDir(); err != nil {
		return errors.New("Error on Getting workDir")
	}
	if err := m.projectInit(); err != nil {
		return errors.New("Error on Getting workDir")
	}
	if err := m.getHTTPpackage(); err != nil {
		return errors.New("Error on Getting workDir")
	}
	if err := m.getDBdriver(); err != nil {
		return errors.New("Error on Getting workDir")
	}
	if err := m.generateTemplate(); err != nil {
		return errors.New("Error on Getting workDir")
	}
	if err := m.gitInit(); err != nil {
		return errors.New("Error on Getting workDir")
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

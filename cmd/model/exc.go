package model

import (
	"os"
	"os/exec"
)

func (m *Model) Exc() error {
	if err := chDir(); err != nil {
		return err
	}
	if err := m.projectInit(); err != nil {
		return err
	}
	return nil
}

func (m *Model) projectInit() error {
	return exec.Command("go", "mod", "init", m.ProjectName).Run()
}

func chDir() error {
	var wd, wdErr = os.Getwd()
	if wdErr != nil {
		return wdErr
	}
	return os.Chdir(wd)
}

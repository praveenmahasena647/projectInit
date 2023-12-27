package model

import "github.com/charmbracelet/huh"

type Model struct {
	ProjectName string
}

func New() *Model {
	return &Model{}
}

func (m *Model) Update() error {
	var form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("pls Enter your project Name").
				Value(&m.ProjectName),
		),
	)
	return form.Run()
}

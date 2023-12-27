package model

import "github.com/charmbracelet/huh"

type Model struct {
	ProjectName      string
	HTTPpackage      string
	DBdriver         string
	GenerateTemplate bool
	GitInit          bool
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
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Pls Select your HTTP lib || FrameWork").
				Options(
					huh.NewOption("Standered", "Standered"),
					huh.NewOption("Gorlilla/Mux", "github.com/gorilla/mux"),
					huh.NewOption("fastHTTP", "fastHTTP"),
					huh.NewOption("go-Chi", "github.com/go-chi/chi/v5"),
					huh.NewOption("go-Gin", "github.com/gin-gonic/gin"),
					huh.NewOption("Echo", "github.com/labstack/echo/v4"),
				).
				Value(&m.HTTPpackage),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Pls Select your DB driver").
				Options(
					huh.NewOption("MongoDB", "go.mongodb.org/mongo-driver/mongo"),
					huh.NewOption("gorm-postgres", "gorm-postgres"),
					huh.NewOption("gorm-MySQL", "gorm-MySQL"),
					huh.NewOption("Standered-postgresql", "github.com/lib/pq"),
				).
				Value(&m.HTTPpackage),
		),
		huh.NewGroup(
			huh.NewConfirm().
				Title("Generate Template").
				Affirmative("Yes").
				Negative("No").
				Value(&m.GenerateTemplate),
		),
		huh.NewGroup(
			huh.NewConfirm().
				Title("Initiate Git").
				Affirmative("Yes").
				Negative("No").
				Value(&m.GitInit),
		),
	)
	return form.Run()
}

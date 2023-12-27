package cmd

import (
	"github.com/praveenmahasena647/projectInit/cmd/model"
)

func Start() error {
	var m = model.New()
	if err := m.Update(); err != nil {
		return err
	}
	var excErr = m.Exc()
	if excErr != nil {
		return excErr
	}
	return nil
}

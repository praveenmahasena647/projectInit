package cmd

import (
	"fmt"

	"github.com/praveenmahasena647/projectInit/cmd/model"
)

func Start() error {
	var m = model.New()
	if err := m.Update(); err != nil {
		return err
	}
	var excErr = m.Exc()
	if excErr != nil {
		fmt.Println("error came here")
		fmt.Println(excErr)
		return excErr
	}
	return nil
}

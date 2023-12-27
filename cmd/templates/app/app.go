package app

import (
	"fmt"

	"github.com/praveenmahasena647/projectInit/cmd/templates/app/dbs"
	"github.com/praveenmahasena647/projectInit/cmd/templates/app/handlers"
)

func Start() error {
	var dbErr = dbs.Connect()
	if dbErr != nil {
		return fmt.Errorf("Database conection Err")
	}
	var router = handlers.New("")
	var routerErr = router.Init()
	if routerErr != nil {
		return routerErr
	}
	return nil
}

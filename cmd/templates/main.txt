package main

import (
	"log"
	"os"

	"github.com/praveenmahasena647/projectInit/cmd/templates/app"
)

func main() {
	var appErr = app.Start()
	if appErr != nil {
		log.Printf("%v", appErr)
		os.Exit(1)
	}
}

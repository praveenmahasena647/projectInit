package main

import (
	"log"
	"os"

	"github.com/praveenmahasena647/projectInit/cmd"
)

func main() {
	var appErr = cmd.Start()
	if appErr != nil {
		log.Printf("%v", appErr)
		os.Exit(1)
	}
}

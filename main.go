package main

import (
	"fmt"
	"os"

	"github.com/praveenmahasena647/projectInit/cmd"
)

func main() {
	var appErr = cmd.Start()
	if appErr != nil {
		fmt.Printf("%v", appErr)
		os.Exit(1)
	}
}

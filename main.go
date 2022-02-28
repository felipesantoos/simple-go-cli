package main

import (
	"command-line/app"
	"log"
	"os"
)

func main() {
	app := app.Generate()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

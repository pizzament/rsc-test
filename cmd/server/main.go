package main

import (
	"log"
	"os"

	new_app "github.com/pizzament/rsc-test/internal/app"
)

func main() {
	log.Println("app starting")

	app, err := new_app.NewApp(os.Getenv("CONFIG_FILE"))
	if err != nil {
		panic(err)
	}

	if err := app.ListenAndServe(); err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"log"

	"github.com/Galbeyte1/go-user-man-ser/internal/database"
)

type App struct {}

func (app *App) Run() error {
	fmt.Println("Starting Application...")

	var err error
	fmt.Println("Instantiating Database")
	_, err = database.NewDatabase()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Welcome Go User Management Service")
	
	app := App{}
	if err := app.Run(); err != nil {
		log.Fatalln("Error Starting up Applicaition")
		log.Fatalf("Error: %s", err)
	} // Terminate gracefully
}
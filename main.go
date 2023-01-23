package main

import (
	"github.com/andres15alvarez/ticket_manager/app"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app := app.CreateApp()
	app.Server.Run()
}

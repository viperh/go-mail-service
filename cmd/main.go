package main

import app2 "go-mail-service/internal/app"

func main() {
	app := app2.NewApp()
	app.Run()
}

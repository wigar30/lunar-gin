package main

import "lunar/internal/app"

func main() {
	app := app.NewWire()
	app.ListenApp()
}
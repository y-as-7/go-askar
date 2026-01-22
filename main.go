package main

import (
	"github.com/y-as-7/go-askar/app/DashAskar"
	"github.com/y-as-7/go-askar/bootstrap"
)

func main() {
	app := bootstrap.NewApplication()

	// Register DashAskar Admin Panel via Provider
	adminPanel := DashAskar.PanelProvider()
	
	adminPanel.Register(app.Router)

	app.Run()
}

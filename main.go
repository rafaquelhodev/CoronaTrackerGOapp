package main

import (
	"controllers"
	"fmt"
	"log"
	"models"
	"net/http"

	"github.com/gernest/utron"
)

func main() {

	// Start the MVC App
	app, err := utron.NewMVC()
	if err != nil {
		log.Fatal(err)
	}

	// Register Models
	app.Model.Register(&models.Users{})
	app.Model.Register(&models.Clients{})

	// CReate Models tables if they dont exist yet
	app.Model.AutoMigrateAll()

	// Register Controller
	app.AddController(controllers.NewMediumController)
	app.AddController(controllers.NewFindInfectedControllerController)

	// Start the server
	port := fmt.Sprintf(":%d", app.Config.Port)
	log.Fatal(http.ListenAndServe(port, app))
}

package main

import (
	"ice-creams-app/internal/controllers/handlers/protected-routes"
	"ice-creams-app/internal/controllers/handlers/public-routes"
	"ice-creams-app/internal/controllers/router"
	"ice-creams-app/internal/db"
	"ice-creams-app/internal/pkg/logger"
	"ice-creams-app/internal/repositories/icecreams-repo"
	"ice-creams-app/internal/repositories/users-repo"
	authorization "ice-creams-app/internal/services/auth-services"
	products "ice-creams-app/internal/services/icecream-services"
)

func main() {

	log := logger.GetLogger()

	log.Info("Connecting to databases")
	icecreamsData := db.New()
	if err := icecreamsData.Connect("icecreams_catalog"); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	usersData := db.New()
	if err := usersData.Connect("user_data"); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	repoIcecreams := icecreams.NewRepo(icecreamsData.Connection)
	serviceIcecreams := products.NewService(repoIcecreams)
	handlerIcecreams := protected.NewIceCreamHandler(serviceIcecreams)

	repoAuth := users.NewRepo(usersData.Connection)
	serviceAuth := authorization.NewService(repoAuth)
	handlerAuth := public.NewAuthHandler(serviceAuth)

	app := router.New(handlerAuth, handlerIcecreams, icecreamsData.Connection, usersData.Connection)

	r := router.SetupRouter(app)
	log.Info("Starting the server on port 8080")
	r.Run()

}

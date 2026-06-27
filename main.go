package main

import (
	"api-rest/config"
	"api-rest/database"
	"api-rest/routes"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	//Configs
	config.LoadConfigEnv()

	//Database connection
	database.RunDatabase()

	//App Fiber
	app := fiber.New()

	//Api version
	api := app.Group("/api/v1")

	//Routes
	routes.IndexRoute(api)
	routes.UserRoutes(api)
	routes.TaskRoutes(api)
	routes.ActivityRoutes(api)

	//Run server
	log.Println(app.Listen(":" + config.PORT))
}

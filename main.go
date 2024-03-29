package main

import (
	"github.com/gin-gonic/gin"
	"greatcomcatengineering.com/backend/configs"
	"greatcomcatengineering.com/backend/database"
	"greatcomcatengineering.com/backend/routes"
	"log"
)

// @title Great Comcat Engineering API
// @version 1
// @description This is the API for the Great Comcat Engineering project.
// @host localhost:8080
// @BasePath /v0
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	database.ConnectToMongoDB()
	router := gin.Default()
	routes.IntroRoutes(router)
	routes.SwaggerRoutes(router)
	versionControlled := router.Group("/" + configs.AppConfig().App.ApiVersion)
	{
		routes.DefaultRoutes(versionControlled)
		routes.UserRoutes(versionControlled)
		routes.ProductRoutes(versionControlled)
	}
	router.Run(configs.AppConfig().App.Host + configs.AppConfig().App.Port)
	database.DisconnectFromMongoDB()
}

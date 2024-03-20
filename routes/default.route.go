package routes

import (
	"github.com/gin-gonic/gin"
	"greatcomcatengineering.com/backend/configs"
)

// This is for /v0
func DefaultRoutes(group *gin.RouterGroup) {

	group.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the GreatComCatEngineering API",
			"version": configs.AppConfig().App.AppVersion,
			"author":  "Gowthaman Ravindrathas",
		})
	})

}

// This is for /
func IntroRoutes(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Welcome to the GreatComCatEngineering API",
			"version": configs.AppConfig().App.AppVersion,
			"author":  "Gowthaman Ravindrathas",
			"api":     "v0",
		})
	})

}

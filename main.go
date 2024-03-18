package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"greatcomcatengineering.com/backend/services/user"
)

func main() {
	godotenv.Load()
	router := gin.Default()

	router.POST("/user/create", func(c *gin.Context) {
		user.HandleCreateUser(c.Writer, c.Request)
	})

	router.GET("/user/:id", func(c *gin.Context) {
		user.HandleGetUserByEmail(c.Writer, c.Request)
	})

	router.Run(":8080")
}

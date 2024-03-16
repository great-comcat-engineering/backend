package main

import (
	"github.com/gin-gonic/gin"
	login "greatcomcatengineering.com/backend/api/auth/login"
)

func main() {
	router := gin.Default()

	router.GET("/auth/login", func(c *gin.Context) {
		login.Handler(c.Writer, c.Request)
	})
	router.Run(":8080")
}

package routes

import (
	"github.com/gin-gonic/gin"
	"greatcomcatengineering.com/backend/middleware"
	"greatcomcatengineering.com/backend/services/user"
)

func UserRoutes(group *gin.RouterGroup) {
	userGroup := group.Group("/user")
	{
		// Public routes
		userGroup.POST("/create", func(c *gin.Context) {
			user.HandleCreateUser(c)
		})

		userGroup.POST("/login", func(c *gin.Context) {
			user.HandleLogin(c)
		})

		// Default user routes
		userGroup.Use(middleware.JWTAuthMiddleware())
		{
			userGroup.GET("/current", func(c *gin.Context) {
				user.HandleGetCurrentUser(c)
			})

			// Admin user routes
			userGroup.Use(middleware.IsAdmin())
			{
				userGroup.GET("/:email", func(c *gin.Context) {
					user.HandleGetUserByEmail(c)
				})

				userGroup.GET("/all", func(c *gin.Context) {
					user.HandleGetAllUsers(c)
				})
			}
		}
	}
}

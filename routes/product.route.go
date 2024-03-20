package routes

import (
	"github.com/gin-gonic/gin"
	"greatcomcatengineering.com/backend/middleware"
	"greatcomcatengineering.com/backend/services/product"
)

func ProductRoutes(group *gin.RouterGroup) {
	userGroup := group.Group("/product")
	{
		// Public routes
		userGroup.GET("/all", func(c *gin.Context) {
			product.HandleGetAllProducts(c)
		})

		// Default user routes
		userGroup.Use(middleware.JWTAuthMiddleware())
		{

			// Admin user routes
			userGroup.Use(middleware.IsAdmin())
			{
				userGroup.POST("/create", func(c *gin.Context) {
					product.HandleCreateProduct(c)
				})

				userGroup.POST("/createMany", func(c *gin.Context) {
					product.HandleCreateManyProducts(c)
				})
			}
		}
	}
}

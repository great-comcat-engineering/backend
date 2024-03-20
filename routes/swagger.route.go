package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/http-swagger"
	_ "greatcomcatengineering.com/backend/docs"
)

func SwaggerRoutes(router *gin.Engine) {
	router.GET("/swagger/*any", gin.WrapH(httpSwagger.WrapHandler))
}

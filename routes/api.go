package routes

import (
	"go-backbone/src/controllers"
	"go-backbone/src/middlewares"

	"github.com/gin-gonic/gin"
)

func ApiRoute(route *gin.Engine) {
	route.GET("/", controllers.GetServerHealth)

	api := route.Group("/api")
	{
		api.GET("/", controllers.GetApiHealth)
	}

	secure := route.Group("/api").Use(middlewares.Auth())
	{
		secure.GET("/secure", controllers.GetSecureApiHealth)
	}
}

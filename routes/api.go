package routes

import (
	"go-backbone/src/controller"

	"github.com/gin-gonic/gin"
)

func ApiRoute(route *gin.Engine) {
	route.GET("/", controller.GetServerHealth)

	api := route.Group("/api")
	{
		api.GET("/", controller.GetApiHealth)
	}
}

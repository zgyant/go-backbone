package routes

import (
	"go-backbone/src/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiRoute(route *gin.Engine) {
	route.GET("/", func(contx *gin.Context) {
		contx.JSON(http.StatusOK, gin.H{
			"message": utils.Bt(`Server Running at ${env("SERVER_URL")}:${env("SERVER_PORT")}`),
		})
	})

	api := route.Group("/api")
	{
		api.GET("/", func(contx *gin.Context) {
			contx.JSON(http.StatusOK, gin.H{
				"message": utils.Bt(`/api is healthy`),
			})
		})
	}
}

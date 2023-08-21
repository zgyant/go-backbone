package controllers

import (
	"go-backbone/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetServerHealth(contx *gin.Context) {
	contx.JSON(http.StatusOK, gin.H{
		"message": utils.Bt(`Server Running at ${env("SERVER_URL")}:${env("SERVER_PORT")}`),
	})
}

func GetApiHealth(contx *gin.Context) {
	contx.JSON(http.StatusOK, gin.H{
		"message": utils.Bt(`/api is healthy`),
	})
}

func GetSecureApiHealth(contx *gin.Context) {
	contx.JSON(http.StatusOK, gin.H{
		"message": utils.Bt(`secure /api is healthy`),
	})
}

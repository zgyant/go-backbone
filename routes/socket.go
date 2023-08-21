package routes

import (
	"go-backbone/configs"
	"go-backbone/src/middlewares"
	"log"

	"github.com/gin-gonic/gin"

	socketio "github.com/googollee/go-socket.io"
)

func SocketRoute(route *gin.Engine) {
	//initialize socket
	server := configs.SocketInit()

	secure := route.Group("/").Use(middlewares.Auth())
	{
		secure.GET("/socket/*any", gin.WrapH(server))
		secure.POST("/socket/*any", gin.WrapH(server))
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		log.Println("New WebSocket connection:", s.ID())
		return nil
	})
}

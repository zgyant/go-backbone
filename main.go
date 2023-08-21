package main

import (
	"go-backbone/configs"
	"go-backbone/routes"
	"go-backbone/src/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()

	//start db & migrate
	configs.DbInit()
	configs.DbMigrate()

	//initialize api routes
	routes.ApiRoute(route)
	//initialize socket routes
	routes.SocketRoute(route)

	// Construct the address in the format "host:port"
	address := utils.Env("SERVER_URL") + ":" + utils.Env("SERVER_PORT")

	// Start the Gin server
	err := route.Run(address)
	if err != nil {
		panic(err)
	}
}

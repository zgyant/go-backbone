package main

import (
	"go-backbone/routes"
	"go-backbone/src/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()
	routes.ApiRoute(route)

	// Construct the address in the format "host:port"
	address := utils.Env("SERVER_URL") + ":" + utils.Env("SERVER_PORT")

	// Start the Gin server
	route.Run(address)
}

package configs

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
)

var server *socketio.Server

// SocketInit initializes a new Socket Server if not already initialized.
func SocketInit() *socketio.Server {
	log.Println("Initializing Socket...")
	if server == nil {
		log.Println("Spawning New Socket Server...")
		server = socketio.NewServer(nil)
	}

	return server
}

// GetSocketServer returns the initialized Socket Server or nil if not initialized.
func GetSocketServer() *socketio.Server {
	if server == nil {
		log.Println("Socket hasn't been initialized.. please use SocketInit()")
		return nil
	}

	return server
}

// ServeSocket starts serving the Socket Server in a Goroutine.
func ServeSocket() {
	if server == nil {
		log.Println("Socket hasn't been initialized.. please use SocketInit()")
		return
	}

	go func() {
		log.Println("Starting Socket Server...")
		if err := server.Serve(); err != nil {
			log.Println("Socket Server Error:", err)
		}
	}()
}

// CleanSocket stops and closes the Socket Server.
func CleanSocket() {
	if server == nil {
		log.Println("Socket hasn't been initialized.. please use SocketInit()")
		return
	}

	log.Println("Cleaning up Socket Server...")
	server.Close()
}

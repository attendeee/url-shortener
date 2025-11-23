package main

import (
	"os"

	"github.com/attendeee/url-shortener/handlers"
	"github.com/attendeee/url-shortener/server"
	"github.com/attendeee/url-shortener/storage/lite"
)

func main() {

	server.Init()

	lite.Init()

	handlers.Init()

	server.AddRouter(handlers.All)

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		server.RunServer()
	}()

	// Wait until server needs shutdown
	server.GracefulShutdown()
	os.Exit(0)

}

package main

import (
	"os"
	"time"

	"github.com/attendeee/url-shortener/handlers"
	"github.com/attendeee/url-shortener/server"
	"github.com/attendeee/url-shortener/storage/lite"
)

func main() {
	// Todo: add external .toml/.json/.yaml config
	server.Init(server.Config{
		Addr:         "127.0.0.1:3000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 15,
	})

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

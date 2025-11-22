package main

import (
	"net/http"
	"os"
	"time"

	"github.com/attendeee/url-shortener/server"
	"github.com/gorilla/mux"
)

func main() {
	// Todo: add external .toml/.json/.yaml config
	server.Init(server.Config{
		Addr:         "0.0.0.0:3000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 15,
	})

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("I am working"))
	})

	server.AddRouter(r)

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		server.RunServer()
	}()

	// Wait until server needs shutdown
	server.GracefulShutdown()
	os.Exit(0)

}

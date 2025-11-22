package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var gracefulTimeout time.Duration
var server *http.Server

type Config struct {
	Addr string

	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	IdleTimeout  time.Duration
}

// Inits server stuff //
func Init(config Config) {
	log.Println("Init server")
	server = &http.Server{
		Addr: config.Addr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: config.WriteTimeout,
		ReadTimeout:  config.ReadTimeout,
		IdleTimeout:  config.IdleTimeout,
	}
}

func AddRouter(r http.Handler) {
	log.Println("Add routers")
	server.Handler = r
}

func RunServer() error {
	log.Println("Start server")
	err := server.ListenAndServe()
	if err != nil {
		log.Println("Runtime server error: ", err)
	}
	return err
}

func GracefulShutdown() {
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), gracefulTimeout)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)
	log.Println("shutting down")
}

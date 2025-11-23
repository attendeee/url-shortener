package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var server *http.Server

type Config struct {
	Addr string

	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	IdleTimeout  time.Duration

	GracefuleTimeout time.Duration
}

var config Config

func GracefulShutdown() {
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), config.GracefuleTimeout)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)
	log.Println("shutting down")
}

// Inits server stuff //
func Init() {
	log.Println("Init .env")

	err := godotenv.Load(".env")
	if err != nil {
		panic("Unable to load .env")
	}

	loadConfig()

	log.Println("Init server")
	server = &http.Server{
		Addr: config.Addr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: config.WriteTimeout,
		ReadTimeout:  config.ReadTimeout,
		IdleTimeout:  config.IdleTimeout,
	}
}

func RunServer() error {
	log.Println("Start server")
	err := server.ListenAndServe()
	if err != nil {
		log.Println("Runtime server error: ", err)
	}
	return err
}

func loadConfig() {
	config.Addr = os.Getenv("ADDR")

	w, err := strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))
	if err != nil {
		log.Println("No WRITE_TIMEOUT in .env")
		panic(err)
	}
	r, err := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	if err != nil {
		log.Println("No READ_TIMEOUT in .env")
		panic(err)
	}
	i, err := strconv.Atoi(os.Getenv("IDLE_TIMEOUT"))
	if err != nil {
		log.Println("No IDLE_TIMEOUT in .env")
		panic(err)
	}

	config.WriteTimeout = time.Second * time.Duration(w)
	config.ReadTimeout = time.Second * time.Duration(r)
	config.IdleTimeout = time.Second * time.Duration(i)

}

func AddRouter(r http.Handler) {
	log.Println("Add routers")
	server.Handler = r
}

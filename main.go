package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/caconka/go-template/connectors"
	"github.com/caconka/go-template/handlers"
	"github.com/caconka/go-template/services"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	port := 8080
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	server := services.NewServer(addr, 5*time.Second, 10*time.Second, 60*time.Second)

	addRoutes(server)

	fmt.Println(fmt.Sprintf("Server address :%s ", addr))
	err := server.Start()
	if err != nil {
		log.Fatal(fmt.Sprintf("listen and serve err %s", err))
	}

}

func addRoutes(s services.Server) {
	jokeProvider := connectors.NewJokeProvider()
	jokeService := services.NewJokeService(jokeProvider)
	jokeHandler := handlers.NewJokeHandler(jokeService)

	s.AddRoute(http.MethodGet, "/health", handlers.HealthCheck)
	s.AddRoute(http.MethodGet, "/joke", jokeHandler.GetRandomJoke)
}

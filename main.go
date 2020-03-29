package main

import (
	"fmt"
	"os"
	"time"

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
	healthCheckHandler := handlers.HealthCheck
	server.AddRoute("/health", healthCheckHandler)

	fmt.Println(fmt.Sprintf("Server address :%s ", addr))
	err := server.Start()
	if err != nil {
		log.Fatal(fmt.Sprintf("listen and serve err %s", err))
	}

}

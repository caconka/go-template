package services

import (
	"log"
	"net/http"
	"os"
	"time"
)

type server struct {
	router   *http.ServeMux
	logger   *log.Logger
	instance *http.Server
}

type Server interface {
	Start() error
	AddRoute(path string, handler http.HandlerFunc)
}

func NewServer(listenAddr string, readTimeout, writeTimeout, idleTimeout time.Duration) Server {
	router := http.NewServeMux()
	logger := log.New(os.Stdout, "go-template-http: ", log.Ldate)

	s := &http.Server{
		Addr:         listenAddr,
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}

	return &server{
		instance: s,
		logger:   logger,
		router:   router,
	}
}

func (s *server) Start() error {
	return s.instance.ListenAndServe()
}

func (s *server) AddRoute(path string, handler http.HandlerFunc) {
	s.router.HandleFunc(path, handler)
}

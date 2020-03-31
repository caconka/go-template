package services

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

type server struct {
	router   *httprouter.Router
	logger   *log.Logger
	instance *http.Server
}

type Server interface {
	Start() error
	AddRoute(method, path string, hf http.Handler, middlewares ...func(http.Handler) http.Handler)
}

func NewServer(listenAddr string, readTimeout, writeTimeout, idleTimeout time.Duration) Server {
	router := httprouter.New()
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

func (s *server) AddRoute(method, path string, h http.Handler, middlewares ...func(http.Handler) http.Handler) {
	for _, md := range middlewares {
		h = md(h)
	}

	s.router.Handler(method, path, h)
}

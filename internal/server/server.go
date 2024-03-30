package server

import (
	"fmt"
	"main/internal/handlers"
	"main/internal/storage"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Host    string
	Port    string
	Storage storage.Storage
	Router  *mux.Router
}

func New() *Server {
	var s Server
	s.Router = mux.NewRouter()
	s.configureRouter()
	s.Host = "localhost"
	s.Port = "9991"

	return &s
}

func (s *Server) configureRouter() {
	s.Router.HandleFunc("/get", handlers.Get(s.Storage))
	s.Router.HandleFunc("/set", handlers.Set(s.Storage))
}

func (s *Server) Start() error {
	const op = "server.start"
	err := http.ListenAndServe(s.Host+":"+s.Port, s.Router)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

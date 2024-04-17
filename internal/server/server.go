package server

import (
	"fmt"
	"main/internal/databus"
	"main/internal/handlers"
	"main/internal/parcel"
	"main/internal/server/mw"
	"main/internal/storage"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type Server struct {
	Host string
	Port string

	Server  *http.Server
	Storage storage.Storage
	Databus handlers.Databus
	Router  *mux.Router
	Mode    string
}

func New() *Server {
	var s Server
	s.Router = mux.NewRouter()
	s.Databus = databus.New(parcel.Parcel{})
	s.configureRouter()
	s.Host = "localhost"
	s.Port = "9991"

	s.Server = &http.Server{
		Addr:         s.Host + ":" + s.Port,
		Handler:      mw.Logging(s.Router),
		ReadTimeout:  time.Duration(5 * time.Second),
		WriteTimeout: time.Duration(5 * time.Second),
		IdleTimeout:  time.Duration(1 * time.Hour),
	}
	return &s
}

func (s *Server) configureRouter() {
	s.Router.NotFoundHandler = http.HandlerFunc(handlers.DefaultHandler)
	getRouter := s.Router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/service/get", handlers.Get(s.Databus))
	// getRouter.HandleFunc("/service/get/service_name:[a-z]", handlers.GetByName())

	postRouter := s.Router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/service/set", handlers.Set(s.Databus))
	// s.Router.HandleFunc("service/register", handlers.ServiceRegister())
	// s.Router.HandleFunc("user/register", handlers.UserRegister())
	// s.Router.HandleFunc("user/login", handlers.UserLogin())

	// s.Router.HandleFunc("/user/control_panel", handlers.ControlPanel())
}

func (s *Server) Start() error {
	const op = "server.start"
	log.Info().Msg("server is started")
	err := s.Server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

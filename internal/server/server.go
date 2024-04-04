package server

import (
	"fmt"
	"main/internal/databus"
	"main/internal/handlers"
	"main/internal/parcel"
	"main/internal/storage"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type Server struct {
	Host string
	Port string

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

	return &s
}

func (s *Server) configureRouter() {
	s.Router.HandleFunc("/service/get", handlers.Get(s.Databus))
	s.Router.HandleFunc("/service/set", handlers.Set(s.Databus))
	// s.Router.HandleFunc("service/register", handlers.ServiceRegister())
	// s.Router.HandleFunc("user/register", handlers.UserRegister())
	// s.Router.HandleFunc("user/login", handlers.UserLogin())

	// s.Router.HandleFunc("/user/control_panel", handlers.ControlPanel())
}

func (s *Server) Start() error {
	const op = "server.start"
	log.Info().Msg("server is started")
	baseAddr := s.Host + ":" + s.Port
	log.Info().Msg(baseAddr)
	err := http.ListenAndServe(baseAddr, s.Router)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

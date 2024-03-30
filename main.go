package main

import (
	"main/internal/server"
	"main/internal/storage"

	"github.com/rs/zerolog/log"
)

var Storage storage.Storage

func main() {
	const op = "main"
	serv := server.New()

	err := serv.Start()
	if err != nil {
		log.Err(err).Msg(op)
	}
}

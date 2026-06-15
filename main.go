package main

import (
	"github.com/rs/zerolog/log"

	"github.com/heitorbolisw4/api-students/api"
)

func main() {
	server := api.NewServer()

	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Error().Err(err).Msg("Server initialization failed.")
	}

}

package logging

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func WriteToFile() {
	file, err := os.OpenFile("PimpMyEvening.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open log file")
	}
	log.Logger = log.Output(file)
}

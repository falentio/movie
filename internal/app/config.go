package app

import (
	"github.com/kamalshkeir/kenv"
	"github.com/rs/zerolog/log"
)

var config Config

type Config struct {
	Address     string `kenv:"ADDRESS|:8080"`
	DatabaseDsn string `kenv:"DATABASE_DSN|./database/database.db"`
	Scrape      bool   `kenv:"SCRAPE|false"`
	Debug       bool   `kenv:"DEBUG|false"`
}

func LoadConfig() Config {
	if err := kenv.Fill(&config); err != nil {
		log.Fatal().Err(err).Msg("failed to load configuration")
	}
	return config
}

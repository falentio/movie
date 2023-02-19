package main

import (
	"github.com/falentio/movie/internal/app"
	"github.com/rs/zerolog/log"
)

func main() {
	app := &app.Application{
		Config: app.LoadConfig(),
	}
	app.InitLogger()
	if err := app.InitRepository(); err != nil {
		log.Fatal().Err(err).Msg("failed to load repository")
	}
	app.InitHandler()
	app.InitScraper()
	app.Scrape()
	if err := app.Listen(); err != nil {
		log.Fatal().Err(err).Msg("failed start server")
	}
}

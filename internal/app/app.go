package app

import (
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"xorm.io/xorm"

	"github.com/falentio/movie/internal/domain"
	"github.com/falentio/movie/internal/pkg/scraper"
	"github.com/falentio/movie/internal/service/movie"
)

type Application struct {
	Config  Config
	App     *fiber.App
	Repo    domain.Repository
	Scraper *scraper.Scraper
}

func (a *Application) InitLogger() {
	lvl := zerolog.ErrorLevel
	if a.Config.Debug {
		lvl = zerolog.DebugLevel
	}
	log.Logger.Level(lvl)
}

func (a *Application) InitRepository() error {
	engine, err := xorm.NewEngine("sqlite3", a.Config.DatabaseDsn)
	if err != nil {
		return err
	}
	a.Repo.Movie = &movie.MovieRepo{Engine: engine}
	return nil
}

func (a *Application) InitScraper() {
	a.Scraper = scraper.New(a.Repo.Movie)
}

func (a *Application) InitHandler() {
	if a.App == nil {
		a.App = fiber.New()
	}

	movieService := &movie.MovieService{
		Repo: a.Repo.Movie,
	}

	api := fiber.New()
	api.Mount("/movie", movie.MovieRouter{Service: movieService}.App())

	a.App.Mount("/api", api)
	a.App.Static("/database", "./database", fiber.Static{
		Download: true,
	})
}

func (a *Application) Listen() error {
	return a.App.Listen(a.Config.Address)
}

func (a *Application) Scrape() {
	if !a.Config.Scrape {
		return
	}
	ticker := time.NewTicker(time.Hour * 24)
	for {
		a.Scraper.Collect()
		<-ticker.C
	}
}

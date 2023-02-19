package main

import (
	"github.com/falentio/movie/internal/app"
	"github.com/falentio/movie/internal/pkg/scraper"
	"github.com/falentio/movie/internal/service/movie"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

func main() {
	config := app.LoadConfig()
	engine, err := xorm.NewEngine("sqlite3", config.DatabaseDsn)
	if err != nil {
		panic(err)
	}
	repo := movie.MovieRepo{Engine: engine}

	od := scraper.Otakudesu{
		Repo: &repo,
	}
	od.Collect()
}

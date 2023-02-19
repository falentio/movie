package domain

import "context"

type Movie struct {
	Id          int `json:"id" xorm:"autoincr pk <-"` 
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Thumbnail   string `json:"thumbnail"`
	Tags        []string    `json:"tags" xorm:"-"`
	DownloadUrl []*MovieUrl `json:"downloadUrl" xorm:"-"`
}

type MovieUrl struct {
	Id          int `json:"id" xorm:"autoincr pk <-"` 
	MovieId    int `json:"movieId"`
	Server     string `json:"server"`
	Resolution string `json:"resolution"`
	Url        string `json:"url"`
}

type MovieTag struct {
	Id          int `json:"id" xorm:"autoincr pk <-"` 
	MovieId int `json:"movieId"`
	Name    string `json:"name"`
}

type MovieRepository interface {
	CreateMovie(ctx context.Context, m *Movie) error
	CreateMovieTag(ctx context.Context, t ...*MovieTag) error
	CreateMovieUrl(ctx context.Context, u ...*MovieUrl) error
	DeleteMovie(ctx context.Context, id int) error
	SearchMovie(ctx context.Context, m *Movie, page int) ([]*Movie, error)
}

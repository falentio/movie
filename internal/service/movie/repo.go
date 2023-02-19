package movie

import (
	"context"

	"xorm.io/xorm"

	"github.com/falentio/movie/internal/domain"
)

type MovieRepo struct {
	Engine *xorm.Engine
}

func (r *MovieRepo) CreateMovie(ctx context.Context, m *domain.Movie) error {
	_, err := r.Engine.Context(ctx).Insert(m)
	return err
}

func (r *MovieRepo) CreateMovieTag(ctx context.Context, t ...*domain.MovieTag) error {
	_, err := r.Engine.Context(ctx).Insert(t)
	return err
}

func (r *MovieRepo) CreateMovieUrl(ctx context.Context, u ...*domain.MovieUrl) error {
	_, err := r.Engine.Context(ctx).Insert(u)
	return err
}

func (r *MovieRepo) DeleteMovie(ctx context.Context, id int) error {
	_, err := r.Engine.Context(ctx).Delete(&domain.Movie{Id: id})
	return err
}

func (r *MovieRepo) SearchMovie(ctx context.Context, m *domain.Movie, page int) ([]*domain.Movie, error) {
	movies := []*domain.Movie{}
	e := r.Engine.Context(ctx)
	err := e.SQL(`
		SELECT 
			* 
		FROM 
			movie m 
		INNER JOIN 
			(
				SELECT 
					MAX(rowid) id, 
					rank
				FROM 
					movie_fts 
				WHERE 
					movie_fts MATCH ? 
				GROUP BY title
				ORDER BY rank DESC
			) f 
		ON 
			(f.id = m.id)
		LIMIT 10
		OFFSET ?
	`, m.Title, page * 10 - 10).
		Find(&movies)

	for _, m := range movies {
		m.DownloadUrl = make([]*domain.MovieUrl, 0)
		m.Tags = make([]string, 0)
		err := e.
			Table("movie_url").
			Where("movie_id = ?", m.Id).
			Find(&m.DownloadUrl)
		if err != nil {
			return nil, err
		}
		tags := []*domain.MovieTag{}
		err = e.
			Table("movie_tag").
			Where("movie_id = ?", m.Id).
			Find(&tags)
		for _, t := range tags {
			m.Tags = append(m.Tags, t.Name)
		}
		if err != nil {
			return nil, err
		}
	}
	return movies, err
}

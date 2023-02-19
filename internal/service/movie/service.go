package movie

import (
	"context"

	"github.com/falentio/movie/internal/domain"
)

type MovieService struct {
	Repo domain.MovieRepository
}

func (s *MovieService) SearchMovie(ctx context.Context, m *domain.Movie, page int) ([]*domain.Movie, error) {
	return s.Repo.SearchMovie(ctx, m, page)
}

package scraper

import (
	"github.com/falentio/movie/internal/domain"
)

type Scraper struct {
	Otakudesu *Otakudesu
}

func New(r domain.MovieRepository) *Scraper {
	return &Scraper{
		Otakudesu: &Otakudesu{r},
	}
}

func (s *Scraper) Collect() {
	go s.Otakudesu.Collect()
}

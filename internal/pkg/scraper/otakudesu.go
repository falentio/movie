package scraper

import (
	"context"
	"fmt"
	"regexp"

	"github.com/gocolly/colly"
	"github.com/rs/zerolog/log"

	"github.com/falentio/movie/internal/domain"
)

const otakudesuUrl = "https://otakudesu.ltd"
var otakudesuRegex = regexp.MustCompile(fmt.Sprintf("^%s", otakudesuUrl))

type Otakudesu struct {
	Repo domain.MovieRepository
}

func (o *Otakudesu) Collect() {
	log := log.
		With().
		Str("scraper", "otakudesu").
		Logger()
	log.Debug().Msg("starting")

	c := colly.NewCollector()
	c.Async = true
	c.URLFilters = append(c.URLFilters, otakudesuRegex)

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		c.Visit(h.Request.AbsoluteURL(h.Attr("href")))
	})

	c.OnHTML(".download > ul > li", func(h *colly.HTMLElement) {
		reso := h.ChildText("strong")
		m := &domain.Movie{}
		m.Url = h.Request.URL.String()
		m.Description = "otakudesu | otaku desu"
		m.Title = h.DOM.Parents().Find("title").Text()
		m.Thumbnail, _ = h.DOM.Parents().Find(".cukder > .attachment-post-thumbnail.size-post-thumbnail.wp-post-image").Attr("src")

		if err := o.Repo.CreateMovie(context.Background(), m); err != nil {
			log.
				Error().
				Err(err).
				Msg("some error received while inserting to database")
		}

		h.ForEach("a", func(i int, h *colly.HTMLElement) {
			url := &domain.MovieUrl{}
			url.MovieId = m.Id
			url.Resolution = reso
			url.Server = h.Text
			url.Url = h.Attr("href")
			m.DownloadUrl = append(m.DownloadUrl, url)
		})

		if err := o.Repo.CreateMovieUrl(context.Background(), m.DownloadUrl...); err != nil {
			log.
				Error().
				Err(err).
				Msg("some error received while inserting to database")
		}

		log.Debug().Interface("movie", m).Msg("result")
	})

	c.OnScraped(func(r *colly.Response) {
		log.Debug().Str("url", r.Request.URL.String()).Msg("done")
	})

	c.OnRequest(func(r *colly.Request) {
		log.Debug().Str("url", r.URL.String()).Msg("scraping")
	})

	c.Visit(otakudesuUrl)
	c.Wait()
}

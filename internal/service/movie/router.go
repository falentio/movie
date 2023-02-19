package movie

import (
	"github.com/falentio/movie/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type MovieRouter struct {
	Service *MovieService
}

func (r MovieRouter) App() *fiber.App {
	app := fiber.New()
	app.Get("/search", r.SearchMovie)
	return app
}

func (r MovieRouter) SearchMovie(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	m := &domain.Movie{}
	m.Title = c.Query("title")

	movies, err := r.Service.SearchMovie(c.Context(), m, page)
	if err != nil {
		return err
	}
	c.JSON(movies)
	return nil
}

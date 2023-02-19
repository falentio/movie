package movie

import (
	"context"
	"testing"

	"github.com/falentio/movie/internal/domain"
	_ "github.com/mattn/go-sqlite3"

	"xorm.io/xorm"
)

func TestMovieRepo(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite3", "../../../database/testing.db")
	if err != nil {
		t.Fatal(err)
		return
	}
	engine.Exec("delete movie")

	r := &MovieRepo{engine}
	testRepo(t, r)
}

func testRepo(t *testing.T, r *MovieRepo) {
	movies := []*domain.Movie{
		{Id: -1, Title: "minions the rise of gru"},
		{Id: -1, Title: "the rise of king"},
		{Id: -1, Title: "despicable me 3"},
		{Id: -1, Title: "despicable me 1"},
	}
	for _, m := range movies {
		if err := r.CreateMovie(context.Background(), m); err != nil {
			t.Fatalf("%v", err)
		}
		if m.Id == -1 {
			t.Fatal("failed to update id")
		}
		t.Logf("%#+v", m)
	}

	s, err := r.SearchMovie(context.Background(), &domain.Movie{
		Title: "rise",
	}, 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(s) != 2 {
		t.Fatalf("invalid data returned, len %d", len(s))
	}
}

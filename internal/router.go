package internal

import (
	"net/http"
	"time"

	root "github.com/SXsid/cine-lock"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*"},
		AllowedMethods:   []string{"GET", "PUT"},
		AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "Authorization"},
		ExposedHeaders:   []string{"Etag", "Content-Length"},
		AllowCredentials: true,
		MaxAge:           int((12 * time.Hour).Seconds()),
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/v1", v1Routes())
	})

	r.Handle("/*", http.FileServerFS(root.StaticAssests))
	return r
}

func v1Routes() http.Handler {
	r := chi.NewRouter()
	r.Route("/movie", func(r chi.Router) {
		r.Get("/", AllMoviesHandler)
		r.Get("/poll-status", PollSeatStatus)
		r.Patch("/seat-status", ChangeSeatStatus)
	})

	return r
}

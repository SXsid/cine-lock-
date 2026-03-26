package internal

import (
	"net/http"

	root "github.com/SXsid/cine-lock"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServerFS(root.StaticAssests))
	mux.HandleFunc("GET /movies", AllMoviesHandler)
	mux.HandleFunc("GET /poll-status", PollSeatStatus)
	mux.HandleFunc("PATCH /seat-status", ChangeSeatStatus)

	return mux
}

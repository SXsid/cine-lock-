package internal

import (
	"net/http"

	root "github.com/SXsid/cine-lock"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServerFS(root.StaticAssests))
	mux.HandleFunc("/movies", AllMoviesHandler)

	return mux
}

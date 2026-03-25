package internal

import (
	"net/http"
)

var Movies = []Movie{
	{
		Name:   "Dhurandhar",
		Row:    5,
		Coloum: 8,
		Seats:  generateSeats(5, 8),
	},
	{
		Name:   "Dhurandhar:The Revenge",
		Row:    4,
		Coloum: 6,
		Seats:  generateSeats(4, 6),
	},
}

func AllMoviesHandler(w http.ResponseWriter, r *http.Request) {
	WriteJson(w, Movies, http.StatusOK, "")
}

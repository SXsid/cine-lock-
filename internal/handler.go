package internal

import (
	"net/http"
	"strconv"
	"strings"
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

func PollSeatStatus(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("id")
	if value == "" {
		http.Error(w, "Id is required field", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	WriteJson(w, Movies[id].Seats, http.StatusOK, "")
}

func ChangeSeatStatus(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("key")
	idData := r.URL.Query().Get("id")

	statusData := r.URL.Query().Get("status")

	if value == "" || len(value) != 2 || statusData == "" || idData == "" {
		http.Error(w, "id ,key and status is required field and should be combinaton of Row and col index", http.StatusBadRequest)
		return
	}

	status := SeatStatus(statusData)

	if ok := status.IsValid(); !ok {

		http.Error(w, "statsu shoudl SeatStatus type", http.StatusBadRequest)
		return
	}
	data := strings.Split(value, "")

	row, err := strconv.Atoi(data[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	col, err := strconv.Atoi(data[1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	id, err := strconv.Atoi(idData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	Movies[id].Seats[row][col].Status = status
}

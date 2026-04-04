package internal

import (
	"fmt"
	"net/http"
	"strconv"
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
	WriteJSON(w, Movies, http.StatusOK)
}

func PollSeatStatus(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("id")
	if value == "" {
		WriteError(w, "Id is required field", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(value)
	if err != nil {
		WriteError(w, err.Error(), http.StatusBadRequest)
		return

	}
	if id >= len(Movies) {
		WriteError(w, "invalid movie ", http.StatusBadRequest)
		return
	}
	WriteJSON(w, Movies[id].Seats, http.StatusOK)
}

type ChangStatusRequest struct {
	Id     int        `json:"id"`
	Row    int        `json:"row"`
	Col    int        `json:"col"`
	Status SeatStatus `json:"status"`
	UserID string     `json:"userId"`
}

func (r ChangStatusRequest) Validate() error {
	// dummy checks
	if r.Id >= len(Movies) {
		return fmt.Errorf("no movie exists")
	}
	seatData := Movies[r.Id]
	if r.Row >= seatData.Row || r.Col >= seatData.Coloum {
		return fmt.Errorf("no such seat exist")
	}
	seatStatus := SeatStatus(r.Status)
	switch seatStatus {
	case Booked, Hold, Vacant:
		return nil
	default:
		return fmt.Errorf("status shoud be Seat Staus type")

	}
}

func ChangeSeatStatus(w http.ResponseWriter, r *http.Request) {
	res, err := ReadJson[ChangStatusRequest](r)
	if err != nil {
		WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := res.Validate(); err != nil {
		WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}
	// INFO: dumy get the data form the db with a db check if posilbe that seat is vacnat or not
	seat := &Movies[res.Id].Seats[res.Row][res.Col]

	if err := NormalBooking(seat, res.Status, res.UserID); err != nil {
		WriteError(w, err.Error(), http.StatusForbidden)
		return
	}

	writeOK(w, "request successfull..", http.StatusOK)
}

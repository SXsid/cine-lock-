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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := res.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	seat := &Movies[res.Id].Seats[res.Row][res.Col]

	switch res.Status {
	case Hold:
		err = seat.Hold(res.UserID)
	case Booked:
		err = seat.Book(res.UserID)
	default:
		err = seat.Vaccant(res.UserID)

	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	fmt.Fprint(w, "request successfull..")
}

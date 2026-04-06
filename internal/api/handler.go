package api

import (
	"net/http"
	"strconv"

	"github.com/SXsid/cine-lock/internal/domain"
)

type BookingService interface {
	GetSeatData(movieId int) ([][]domain.Seat, error)
	GetAllMovies() []domain.Movie
	GetMovieById(id int) (domain.Movie, error)
	NormalBooking(id, row, col int, action domain.SeatStatus, userId string) error
	BookingUsingMutex(id, row, col int, action domain.SeatStatus, userId string) error
}
type BookingHandler struct {
	bookingService BookingService
}

func NewBookingHandler(bookingService BookingService) *BookingHandler {
	return &BookingHandler{
		bookingService: bookingService,
	}
}

func (b *BookingHandler) AllMoviesHandler(w http.ResponseWriter, r *http.Request) {
	data := b.bookingService.GetAllMovies()
	WriteJSON(w, data, http.StatusOK)
}

func (b *BookingHandler) PollSeatStatus(w http.ResponseWriter, r *http.Request) {
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
	data, err := b.bookingService.GetSeatData(id)
	if err != nil {
		WriteError(w, err.Error(), http.StatusBadRequest)
	}
	WriteJSON(w, data, http.StatusOK)
}

func (b *BookingHandler) ChangeSeatStatus(w http.ResponseWriter, r *http.Request) {
	req, err := ReadJson[ChangStatusRequest](r)
	if err != nil {
		WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := b.bookingService.BookingUsingMutex(req.Id, req.Row, req.Col, req.Status, req.UserID); err != nil {
		WriteError(w, err.Error(), http.StatusForbidden)
		return
	}

	writeOK(w, "request successfull..", http.StatusOK)
}

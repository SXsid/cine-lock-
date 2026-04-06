package service

import (
	"fmt"

	"github.com/SXsid/cine-lock/internal/domain"
)

type BookingService struct {
	movies []domain.Movie
}

func NewBookingService() *BookingService {
	return &BookingService{
		movies: []domain.Movie{
			{
				Name:   "Dhurandhar",
				Row:    5,
				Coloum: 8,
				Seats:  domain.GenerateSeats(5, 8),
			},
			{
				Name:   "Dhurandhar:The Revenge",
				Row:    4,
				Coloum: 6,
				Seats:  domain.GenerateSeats(4, 6),
			},
		},
	}
}

func (b *BookingService) GetAllMovies() []domain.Movie {
	return b.movies
}

func (b *BookingService) GetSeatData(movieId int) ([][]domain.Seat, error) {
	Movies := b.movies
	if movieId >= len(Movies) {
		return nil, fmt.Errorf("no movies exist with this id :%d", movieId)
	}
	return Movies[movieId].Seats, nil
}

func (b *BookingService) GetMovieById(id int) (domain.Movie, error) {
	movies := b.movies
	if id >= len(movies) {
		return domain.Movie{}, fmt.Errorf("no movies exist with this id :%d", id)
	}

	return movies[id], nil
}

func (b *BookingService) NormalBooking(id, row, col int, action domain.SeatStatus, userId string) error {
	if err := b.validate(id, row, col); err != nil {
		return err
	}
	// dummy  db call
	seat := &b.movies[id].Seats[row][col]
	var err error
	switch action {
	case domain.Hold:
		err = seat.Hold(userId)
	case domain.Booked:
		err = seat.Book(userId)
	default:
		err = seat.Vaccant(userId)
	}

	return err
}

// bloackng any other go rouine if one has acqied lock (won't work where dat come form database just for in memroy)
func (b *BookingService) validate(id, row, col int) error {
	Movies := b.movies
	if id >= len(Movies) {
		return fmt.Errorf("no movie exists")
	}
	seatData := Movies[id]
	if row >= seatData.Row || col >= seatData.Coloum {
		return fmt.Errorf("no such seat exist")
	}
	return nil
}

func (b *BookingService) BookingUsingMutex(id, row, col int, action domain.SeatStatus, userId string) error {
	if err := b.validate(id, row, col); err != nil {
		return err
	}
	// db call
	seat := &b.movies[id].Seats[row][col]
	seat.Mu.Lock()
	defer seat.Mu.Unlock()
	var err error
	switch action {
	case domain.Hold:
		err = seat.Hold(userId)
	case domain.Booked:
		err = seat.Book(userId)
	default:
		err = seat.Vaccant(userId)
	}

	return err
}

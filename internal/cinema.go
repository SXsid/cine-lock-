package internal

import "fmt"

type Movie struct {
	Name   string
	Row    int
	Coloum int
	Seats  [][]Seat
}

type SeatStatus string

const (
	Booked  SeatStatus = "booked"
	Hold    SeatStatus = "hold"
	Vacant  SeatStatus = "vacant"
	Selectd SeatStatus = "selected"
)

type Seat struct {
	Name   string
	Status SeatStatus
}

func generateSeats(row, col int) [][]Seat {
	seats := [][]Seat{}
	for i := range row {
		row_seats := []Seat{}
		for j := range col {
			seatName := fmt.Sprintf("%s%d", string(rune('A'+i)), j+1)
			status := Vacant
			row_seats = append(row_seats, Seat{
				Name:   seatName,
				Status: status,
			})

		}
		seats = append(seats, row_seats)
	}

	return seats
}

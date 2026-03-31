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
	Name     string
	Status   SeatStatus
	LockedBy string
}

func (s *Seat) Hold(userId string) error {
	if s.Status != Vacant {
		return fmt.Errorf("seat is not vaccant")
	}
	s.LockedBy = userId
	s.Status = Hold
	return nil
}

func (s *Seat) Book(userId string) error {
	if s.Status != Hold || s.LockedBy != userId {
		return fmt.Errorf("seat can't be booked")
	}
	s.Status = Booked
	return nil
}

func (s *Seat) Vaccant(userId string) error {
	if s.Status != Hold || s.LockedBy != userId {
		return fmt.Errorf("action can't be peformed")
	}
	s.LockedBy = ""
	s.Status = Vacant
	return nil
}

func generateSeats(row, col int) [][]Seat {
	seats := [][]Seat{}
	for i := range row {
		rowSeats := []Seat{}
		for j := range col {
			seatName := fmt.Sprintf("%s%d", string(rune('A'+i)), j+1)
			status := Vacant
			rowSeats = append(rowSeats, Seat{
				Name:   seatName,
				Status: status,
			})

		}
		seats = append(seats, rowSeats)
	}

	return seats
}

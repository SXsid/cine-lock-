package internal

import (
	"fmt"
	"sync"
	"time"
)

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
	mu       sync.Mutex
	Name     string
	Status   SeatStatus
	LockedBy string
}

func (s *Seat) Hold(userId string) error {
	if s.Status != Vacant {
		return fmt.Errorf("seat is not vaccant")
	}
	// to simulate the  db right let say this slep
	// cause if any unprotect go rouine will try to update
	// willsee the value and value will be vacnt and both wil get update the the seat
	// both get seat allocation
	time.Sleep(10 * time.Millisecond)
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

package internal

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestBookingUsingMutex(t *testing.T) {
	seat := Seat{
		Name:     "somehgn",
		Status:   Vacant,
		LockedBy: "",
	}
	var wg sync.WaitGroup
	var sucess int64
	for i := range 100 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if BookingUsingMutex(&seat, Hold, fmt.Sprintf("%d", i)) == nil {
				atomic.AddInt64(&sucess, 1)
			}
		}(i)
	}
	wg.Wait()
	if sucess != 1 {
		t.Errorf("expected 1,got %d", sucess)
	}
}

func TestBookTest(t *testing.T) {
}

func TestNormalBookingRace(t *testing.T) {
	seat := Seat{
		Name:     "somehgn",
		Status:   Vacant,
		LockedBy: "",
	}
	var wg sync.WaitGroup
	var sucess int64
	for i := range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if NormalBooking(&seat, Hold, fmt.Sprintf("%d", i)) == nil {
				atomic.AddInt64(&sucess, 1)
			}
		}()
	}
	wg.Wait()
	if sucess != 1 {
		t.Errorf("expected 1,got %d", sucess)
	}
}

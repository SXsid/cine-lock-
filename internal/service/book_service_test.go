package service

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/SXsid/cine-lock/internal/domain"
)

func TestBookingUsingMutex(t *testing.T) {
	service := NewBookingService()

	var wg sync.WaitGroup
	var sucess int64
	for i := range 100 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if service.BookingUsingMutex(0, 0, 1, domain.Hold, fmt.Sprintf("%d", i)) == nil {
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
	// here wee inject dummy db inreal test
	service := NewBookingService()
	var wg sync.WaitGroup
	var sucess int64
	for i := range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if service.NormalBooking(0, 0, 1, domain.Hold, fmt.Sprintf("%d", i)) == nil {
				atomic.AddInt64(&sucess, 1)
			}
		}()
	}
	wg.Wait()
	if sucess == 1 {
		t.Errorf("expected not 1,got %d", sucess)
	}
}

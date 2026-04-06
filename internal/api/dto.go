package api

import (
	"fmt"

	"github.com/SXsid/cine-lock/internal/domain"
)

type ChangStatusRequest struct {
	Id     int               `json:"id"`
	Row    int               `json:"row"`
	Col    int               `json:"col"`
	Status domain.SeatStatus `json:"status"`
	UserID string            `json:"userId"`
}

func (r ChangStatusRequest) Validate() error {
	// dummy checks
	seatStatus := domain.SeatStatus(r.Status)
	switch seatStatus {
	case domain.Booked, domain.Hold, domain.Vacant:
		return nil
	default:
		return fmt.Errorf("status shoud be Seat Staus type")

	}
}

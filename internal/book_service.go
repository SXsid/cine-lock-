package internal

// race conditon risk
func NormalBooking(seat *Seat, action SeatStatus, userId string) error {
	var err error
	switch action {
	case Hold:
		err = seat.Hold(userId)
	case Booked:
		err = seat.Book(userId)
	default:
		err = seat.Vaccant(userId)
	}
	return err
}

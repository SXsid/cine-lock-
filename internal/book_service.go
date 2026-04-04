package internal

// race conditon risk
// just procedd with payment and book the seat
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

// bloackng any other go rouine if one has acqied lock
func BookingUsingMutex(seat *Seat, action SeatStatus, userId string) error {
	seat.mu.Lock()
	defer seat.mu.Unlock()
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

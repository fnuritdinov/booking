package models

import (
	"booking-service/booking"
	"booking-service/pkg/errors"
)

type Booking struct {
	ID      int
	UserID  int64
	MovieID int64
	Status  booking.Booking
}

func (b *Booking) Validate() error {
	if b.UserID < 1 || b.MovieID < 1 || b.Status == booking.BookingStatus("") {
		return errors.ErrValidate
	}

	return nil
}

const (
	UnspecifiedStatus = "0"
	PendingStatus     = "1"
	ConfirmedStatus   = "2"
	CanceledStatus    = "3"
)

package models

import (
	"github.com/fnuritdinov/booking/pkg/errors"
	"github.com/fnuritdinov/proto/bookingPr"
)

type Booking struct {
	ID      int
	UserID  int64
	MovieID int64
	Status  booking.Booking
}

func (b *Booking) Validate() error {
	if b.UserID < 1 || b.MovieID < 1 || b.Status.Status == 0 {
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

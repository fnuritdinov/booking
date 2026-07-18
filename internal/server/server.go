package server

import (
	"context"
	"errors"

	"github.com/fnuritdinov/booking/internal/models"
	"github.com/fnuritdinov/booking/internal/service"
	errs "github.com/fnuritdinov/booking/pkg/errors"
	"github.com/fnuritdinov/booking/pkg/logger"
	"github.com/fnuritdinov/proto/bookingPr"
	"go.uber.org/zap"
)

type Server struct {
	booking.UnimplementedBookingServiceServer
	service service.Service
	logger  logger.Logger
}

func New(service service.Service, logger logger.Logger) *Server {
	return &Server{
		service: service,
		logger:  logger,
	}
}

func (s *Server) CreateBooking(ctx context.Context, req *booking.CreateBookingRequest) (*booking.CreateBookingResponse, error) {

	b, err := s.service.Create(ctx, models.Booking{
		UserID:  req.UserId,
		MovieID: req.MovieId,
	})
	if err != nil {
		if errors.Is(err, errs.ErrValidate) {
			return nil, errs.ErrBadRequest
		}
		s.logger.Error("error from s.service.Create",
			zap.Error(err))
		return nil, err
	}

	return &booking.CreateBookingResponse{
		Booking: &booking.Booking{
			Id:      int64(b.ID),
			UserId:  b.UserID,
			MovieId: b.MovieID,
			Status:  b.Status,
		},
	}, nil

}

func (s *Server) GetBooking(ctx context.Context, req *booking.GetBookingRequest) (*booking.GetBookingResponse, error) {

	id := req.Id

	b, err := s.service.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, errs.ErrNotFound
		}
		s.logger.Error("error from s.service.GetByID",
			zap.Error(err))
		return nil, err
	}

	var resp booking.GetBookingResponse
	resp.Booking = &booking.Booking{
		UserId:  b.UserID,
		MovieId: b.MovieID,
		Status:  b.Status,
	}
	return &resp, nil
}

func (s *Server) GetUserBookings(ctx context.Context, req *booking.GetUserBookingsRequest) (*booking.GetUserBookingsResponse, error) {

	userID := req.UserId

	myBookings, err := s.service.GetUserBookings(ctx, userID)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, errs.ErrNotFound
		}
		s.logger.Error("error from s.service.GetUserBookings",
			zap.Error(err))
		return nil, err
	}

	bookings := make([]*booking.Booking, 0, len(myBookings))

	for _, b := range myBookings {
		bookings = append(bookings, &booking.Booking{
			Id:      int64(b.ID),
			UserId:  b.UserID,
			MovieId: b.MovieID,
		})
	}

	resp := booking.GetUserBookingsResponse{
		Bookings: bookings,
	}

	return &resp, nil
}

func (s *Server) CancelBooking(ctx context.Context, req *booking.CancelBookingRequest) (*booking.CancelBookingResponse, error) {

	bookingID := req.Id

	err := s.service.Cancel(ctx, bookingID)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, errs.ErrNotFound
		}
		s.logger.Error("error from s.service.Cancel",
			zap.Error(err))
		return nil, err
	}

	return &booking.CancelBookingResponse{}, nil
}

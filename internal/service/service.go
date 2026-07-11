package service

import (
	"booking-service/internal/gateway/movie"
	"booking-service/internal/gateway/user"
	"booking-service/internal/models"
	"booking-service/internal/repository"
	"booking-service/pkg/errors"
	"context"
)

type Service struct {
	repo repository.Repository
	mvG  movie.Gateway
	usG  user.UserGateway
}

func New(repo repository.Repository, mvG movie.Gateway, usG user.UserGateway) Service {
	return Service{
		repo: repo,
		mvG:  mvG,
		usG:  usG,
	}
}

func (s *Service) Create(ctx context.Context, req models.Booking) (models.Booking, error) {
	if req.MovieID < 1 && req.UserID < 1 {
		return models.Booking{}, errors.ErrBadRequest
	}

	b, err := s.repo.Create(ctx, req)
	if err != nil {
		return models.Booking{}, err
	}

	return b, nil
}

func (s *Service) GetByID(ctx context.Context, id int64) (models.Booking, error) {

	if id < 1 {
		return models.Booking{}, errors.ErrValidate
	}

	b, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return models.Booking{}, err
	}

	return b, nil
}

func (s *Service) GetUserBookings(ctx context.Context, userID int64) ([]models.Booking, error) {

	if userID < 1 {
		return nil, errors.ErrValidate
	}

	bookings, err := s.repo.GetUserBookings(ctx, userID)
	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func (s *Service) Cancel(ctx context.Context, bookingID int64) error {

	if bookingID < 1 {
		return errors.ErrValidate
	}

	err := s.repo.Cancel(ctx, bookingID)
	if err != nil {
		return err
	}

	return nil
}

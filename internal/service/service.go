package service

import (
	"context"

	"github.com/fnuritdinov/booking/internal/gateway/movie"
	"github.com/fnuritdinov/booking/internal/gateway/user"
	"github.com/fnuritdinov/booking/internal/models"
	"github.com/fnuritdinov/booking/internal/repository"
	"github.com/fnuritdinov/booking/pkg/errors"
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

func (s *Service) Create(ctx context.Context, request *models.Booking) (*models.Booking, error) {

	err := request.Validate()
	if err != nil {
		return &models.Booking{}, errors.ErrValidate
	}

	b, err := s.repo.Create(ctx, request)
	if err != nil {
		return &models.Booking{}, err
	}

	return b, nil
}

func (s *Service) GetByID(ctx context.Context, id int64) (*models.Booking, error) {

	if id < 1 {
		return &models.Booking{}, errors.ErrValidate
	}

	b, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return &models.Booking{}, err
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

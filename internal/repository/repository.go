package repository

import (
	"context"
	"errors"

	"github.com/fnuritdinov/booking/internal/models"
	errs "github.com/fnuritdinov/booking/pkg/errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) Repository {
	return Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, req models.Booking) (models.Booking, error) {

	const query = `
	INSERT INTO booking (user_id, movie_id)
	VALUES ($1, $2)
	 RETURNING id, user_id, movie_id, status;`

	var b models.Booking

	err := r.db.QueryRow(ctx, query, req.UserID, req.MovieID).Scan(
		&b.ID,
		&b.UserID,
		&b.MovieID,
		&b.Status)
	if err != nil {
		return models.Booking{}, err
	}
	return b, nil
}

func (r *Repository) GetByID(ctx context.Context, id int64) (models.Booking, error) {

	var b models.Booking

	const query = `SELECT user_id, movie_id, status 
					FROM booking
					WHERE id = $1`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&b.UserID,
		&b.MovieID,
		&b.Status)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Booking{}, errs.ErrNotFound
		}
		return models.Booking{}, err
	}

	return b, nil
}

func (r *Repository) GetUserBookings(ctx context.Context, userID int64) ([]models.Booking, error) {

	const query = `SELECT id, movie_id, status 
			FROM booking 
			WHERE user_id = $1`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var bookings []models.Booking

	for rows.Next() {
		var booking models.Booking

		err = rows.Scan(
			&booking.ID,
			&booking.MovieID,
			&booking.Status)
		if err != nil {
			return nil, err
		}

		bookings = append(bookings, booking)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *Repository) Cancel(ctx context.Context, bookingID int64) error {

	const query = `UPDATE booking 
					SET status = $2 
					WHERE id = $1`

	result, err := r.db.Exec(ctx, query, bookingID, models.CanceledStatus)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errs.ErrNotFound
	}

	return nil
}

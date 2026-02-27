package storage

import (
	"GymMembership-api/internal/models"
	"context"
)

func (s *Storage) CreateBooking(ctx context.Context, booking *models.Booking) error {
	query := `INSERT INTO bookings (user_id , class_id) VALUES ($1 , $2)`
	_, err := s.pool.Exec(ctx, query,
		booking.UserId,
		booking.ClassId,
	)
	return err
}

func (s *Storage) GetBookingByUserId(ctx context.Context, userid int64) ([]*models.Booking, error) {
	query := `SELECT id, user_id, class_id, status, created_at FROM bookings WHERE user_id = $1`

	rows, err := s.pool.Query(ctx, query, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bookings := []*models.Booking{}

	for rows.Next() {
		booking := &models.Booking{}
		err := rows.Scan(
			&booking.Id,
			&booking.UserId,
			&booking.ClassId,
			&booking.Status,
			&booking.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}
	return bookings, err
}

func (s *Storage) DeleteBooking(ctx context.Context, id int64) error {
	query := `DELETE FROM bookings WHERE id = $1`
	_, err := s.pool.Exec(ctx, query, id)
	return err
}

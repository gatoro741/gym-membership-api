package storage

import (
	"GymMembership-api/internal/models"
	"context"
)

func (s *Storage) CreateUser(ctx context.Context, user models.User) error {
	query := `INSERT INTO users (name , email , password_hash , role) VALUES ($1 , $2 , $3 , $4)`
	_, err := s.pool.Exec(ctx, query,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.Role,
	)
	return err
}

func (s *Storage) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id, name, email, password_hash, role, created_at FROM users WHERE email = $1`
	user := &models.User{}
	err := s.pool.QueryRow(ctx, query, email).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
	)
	return user, err
}

func (s *Storage) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	query := `SELECT id, name, email, password_hash, role, created_at FROM users WHERE id = $1`
	user := &models.User{}
	err := s.pool.QueryRow(ctx, query, id).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
	)
	return user, err
}

package service

import (
	"GymMembership-api/internal/models"
	"context"
)

type Storage interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type Service struct {
	storage Storage
}

func New(storage Storage) *Service {
	return &Service{storage: storage}
}

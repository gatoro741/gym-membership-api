package handlers

import (
	"GymMembership-api/internal/models"
	"context"
)

type Service interface {
	Register(ctx context.Context, email string, password string, name string) (*models.User, error)
	Login(ctx context.Context, email string, password string) (string, error)
	BuyMembership(ctx context.Context, planId int, userId int64) error
	GetMyMembership(ctx context.Context, userId int64) (*models.UserMembership, error)
	BookClass(ctx context.Context, userId int64, classId int64) (*models.Booking, error)
	CancelBooking(ctx context.Context, bookingId int64, userId int64) error
	GetMyBookings(ctx context.Context, userId int64) ([]*models.Booking, error)
	CreateClass(ctx context.Context, class *models.Class) error
	GetAllClasses(ctx context.Context) ([]*models.Class, error)
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{service: service}
}

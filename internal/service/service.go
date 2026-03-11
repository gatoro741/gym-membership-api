package service

import (
	"GymMembership-api/internal/models"
	"context"
)

type Storage interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetPlanById(ctx context.Context, id int) (*models.MembershipPlan, error)
	CreateUserMembershipPlan(ctx context.Context, userMembership models.UserMembership) error
	GetMembershipByUserId(ctx context.Context, userId int64) (*models.UserMembership, error)
	CreateClass(ctx context.Context, class *models.Class) error
	GetAllClasses(ctx context.Context) ([]*models.Class, error)
	GetClassById(ctx context.Context, id int64) (*models.Class, error)
	CreateBooking(ctx context.Context, userId int64, classId int64) (*models.Booking, error)
	IncrementOccupied(ctx context.Context, classId int64) error
	DeleteBooking(ctx context.Context, id int64, userId int64) error
	GetBookingByUserId(ctx context.Context, userid int64) ([]*models.Booking, error)
}

type Service struct {
	storage Storage
}

func New(storage Storage) *Service {
	return &Service{storage: storage}
}

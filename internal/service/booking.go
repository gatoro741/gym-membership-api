package service

import (
	"GymMembership-api/internal/models"
	"context"
	"errors"
	"log"
)

func (s *Service) BookClass(ctx context.Context, userId int64, classId int64) (*models.Booking, error) {

	userMembership, err := s.storage.GetMembershipByUserId(ctx, userId)
	if err != nil {
		log.Printf("GetMembershipByUserId: failed to get membership : %v", err)
		return nil, err
	}
	if !userMembership.IsActive {
		return nil, errors.New("no active membership")
	}

	class, err := s.storage.GetClassById(ctx, classId)
	if err != nil {
		log.Printf("GetClassById: failed to get class : %v", err)
		return nil, err
	}
	if class.Capacity <= class.Occupied {
		return nil, errors.New("class is full")
	}

	booking, err := s.storage.CreateBooking(ctx, userId, classId)
	if err != nil {
		log.Printf("CreateBooking: failed to create booking : %v", err)
		return nil, err
	}

	err = s.storage.IncrementOccupied(ctx, classId)
	if err != nil {
		log.Printf("IncrementOccupied: failed to inc occupied : %v", err)
		return nil, err
	}
	return booking, nil

}

func (s *Service) CancelBooking(ctx context.Context, bookingId int64, userId int64) error {
	err := s.storage.DeleteBooking(ctx, bookingId, userId)
	if err != nil {
		log.Printf("DeleteBooking: failed to delete: %v", err)
		return err
	}
	return nil
}

func (s *Service) GetMyBookings(ctx context.Context, userId int64) ([]*models.Booking, error) {
	bookings, err := s.storage.GetBookingByUserId(ctx, userId)
	if err != nil {
		log.Printf("GetBookingByUserId: failed to get bookings: %v", err)
		return nil, err
	}
	return bookings, nil
}

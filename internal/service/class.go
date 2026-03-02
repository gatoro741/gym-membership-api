package service

import (
	"GymMembership-api/internal/models"
	"context"
	"log"
)

func (s *Service) CreateClass(ctx context.Context, class *models.Class) error {
	err := s.storage.CreateClass(ctx, class)
	if err != nil {
		log.Printf("CreateClass: failed to create class : %v", err)
		return err
	}
	return nil
}

func (s *Service) GetAllClasses(ctx context.Context) ([]*models.Class, error) {
	classes, err := s.storage.GetAllClasses(ctx)
	if err != nil {
		log.Printf("GetAllClasses: failed to get all classes : %v", err)
		return nil, err
	}
	return classes, nil
}

package service

import (
	"GymMembership-api/internal/models"
	"context"
	"log"

	"github.com/alexedwards/argon2id"
)

func (s *Service) Register(ctx context.Context, email string, password string, name string) (models.User, error) {

	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)

	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return models.User{}, err
	}
	user := models.User{
		Name:         name,
		Email:        email,
		PasswordHash: hash,
		Role:         "client",
	}
	err = s.storage.CreateUser(ctx, user)
	if err != nil {
		log.Printf("failed to create user: %v", err)
		return user, err
	}
	return user, err

}

func (s *Service) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := s.storage.GetUserByEmail(ctx, email)
	if err != nil {
		log.Printf("failed to find user's email: %v", err)
		return "", err
	}
	match, err := argon2id.ComparePasswordAndHash(password, user.PasswordHash)
	if match {
		return GenerateToken(user.Id, user.Role)

	}
	return "", err

}

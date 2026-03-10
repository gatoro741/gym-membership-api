package handlers

import (
	"time"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MembershipRequest struct {
	PlanId int `json:"plan_id"`
}

type BookingRequest struct {
	ClassId   int64 `json:"class_id"`
	BookingId int64 `json:"booking_id"`
}

type ClassRequest struct {
	Title       string    `json:"title"`
	StartTime   time.Time `json:"start_time"`
	TrainerName string    `json:"trainer_name"`
	Capacity    int       `json:"capacity"`
}

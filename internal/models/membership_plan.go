package models

import "time"

type MembershipPlan struct {
	Id           int64     `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Price        float64   `json:"price" db:"price"`
	DurationDays int       `json:"duration_days" db:"duration_days"`
	VisitsLimit  int       `json:"visits_limit" db:"visits_limit"`
	IsActive     bool      `json:"is_active" db:"is_active"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

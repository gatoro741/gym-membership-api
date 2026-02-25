package models

import "time"

type UserMembership struct {
	Id         int64     `json:"id" db:"id"`
	UserId     int64     `json:"user_id" db:"user_id"`
	PlanId     int       `json:"plan_id" db:"plan_id"`
	StartDate  time.Time `json:"start_date" db:"start_date"`
	EndDate    time.Time `json:"end_date" db:"end_date"`
	VisitsLeft *int      `json:"visits_left" db:"visits_left"`
	IsActive   bool      `json:"is_active" db:"is_active"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

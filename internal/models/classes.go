package models

import "time"

type Class struct {
	Id          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	StartTime   time.Time `json:"start_time" db:"start_time"`
	TrainerName string    `json:"trainer_name" db:"trainer_name"`
	Capacity    int       `json:"capacity" db:"capacity"`
	Occupied    int       `json:"occupied" db:"occupied"`
}

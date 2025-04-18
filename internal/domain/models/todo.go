package models

import "time"

type Task struct {
	ID          uint32
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

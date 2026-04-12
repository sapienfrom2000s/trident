package models

import "time"

type Job struct {
	EventId     int
	Status      string
	CreatedAt   time.Time
	ScheduledAt time.Time

	// Using pointers to support null values
	StartedAt     *time.Time
	ExecutionTime *time.Duration
}

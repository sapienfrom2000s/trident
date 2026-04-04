package server

import "time"

type Job struct {
	eventId       int
	status        string
	createdAt     time.Time
	scheduledAt   time.Time
	startedAt     time.Time
	executionTime time.Duration
}

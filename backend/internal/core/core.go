package core

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

type JobOptions struct {
	Timeout time.Duration
}

type HeartBeat struct {
	JobId int
	time  time.Time
}

type NormalizedEvent struct {
	RepoName  string
	CommitSha string
	Branch    string
	Author    string
	Provider  string
}

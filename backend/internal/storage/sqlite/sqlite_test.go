package sqlite_test

import (
	"testing"
	"time"

	"github.com/sapienfrom2000s/trident/backend/internal/server"
	"github.com/sapienfrom2000s/trident/backend/internal/storage/sqlite"
	"github.com/sapienfrom2000s/trident/backend/internal/webhook"
)

func TestEventStorer(t *testing.T) {
	validEvent := webhook.NormalizedEvent{
		RepoName:  "octocat/Hello-World",
		CommitSha: "9fceb02d0ae5",
		Branch:    "main",
		Author:    "Monalisa Octocat",
		Provider:  "github",
	}

	branchNotAvailableInEvent := webhook.NormalizedEvent{
		RepoName:  "octocat/Hello-World",
		CommitSha: "9fceb02d0ae5",
		Author:    "Monalisa Octocat",
		Provider:  "github",
	}

	emptyEvent := webhook.NormalizedEvent{}

	tests := []struct {
		name            string
		normalizedEvent webhook.NormalizedEvent
		want            bool
	}{
		{
			name:            "Valid Event",
			normalizedEvent: validEvent,
			want:            true,
		},
		{
			name:            "Invalid Event: Branch Not Available",
			normalizedEvent: branchNotAvailableInEvent,
			want:            false,
		},
		{
			name:            "Invalid Event: Empty Event",
			normalizedEvent: emptyEvent,
			want:            false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sqlite.StoreEvent(tt.normalizedEvent)
			if got != tt.want {
				t.Errorf("Got: %v, Want: %v. Error: %v", got, tt.want, err)
			}
		})
	}
}

func TestStoreJob(t *testing.T) {
	validJob := server.Job{
		EventId:       1,
		Status:        "waiting",
		CreatedAt:     time.Now(),
		ScheduledAt:   time.Now(),
		StartedAt:     nil,
		ExecutionTime: nil,
	}

	emptyStatusJob := server.Job{
		EventId:       1,
		Status:        "",
		CreatedAt:     time.Now(),
		ScheduledAt:   time.Now(),
		StartedAt:     nil,
		ExecutionTime: nil,
	}

	tests := []struct {
		name string
		job  server.Job
		want bool
	}{
		{
			name: "Valid Job",
			job:  validJob,
			want: true,
		},
		{
			name: "Invalid Job - status is empty",
			job:  emptyStatusJob,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sqlite.StoreJob(validJob)
			if got != tt.want {
				t.Errorf("Got %v, Want %v. Error: %v", got, tt.want, err)
			}
		})
	}
}

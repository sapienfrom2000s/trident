package sqlite_test

import (
	"testing"
	"time"

	"github.com/sapienfrom2000s/trident/backend/internal/core/models"
	"github.com/sapienfrom2000s/trident/backend/internal/storage"
	"github.com/sapienfrom2000s/trident/backend/internal/storage/sqlite"
)

func TestEventStorer(t *testing.T) {
	validEvent := models.Event{
		RepoName:  "octocat/Hello-World",
		CommitSha: "9fceb02d0ae5",
		Branch:    "main",
		Author:    "Monalisa Octocat",
		Provider:  "github",
	}

	branchNotAvailableInEvent := models.Event{
		RepoName:  "octocat/Hello-World",
		CommitSha: "9fceb02d0ae5",
		Author:    "Monalisa Octocat",
		Provider:  "github",
	}

	emptyEvent := models.Event{}

	tests := []struct {
		name  string
		Event models.Event
		want  bool
	}{
		{
			name:  "Valid Event",
			Event: validEvent,
			want:  true,
		},
		{
			name:  "Invalid Event: Branch Not Available",
			Event: branchNotAvailableInEvent,
			want:  false,
		},
		{
			name:  "Invalid Event: Empty Event",
			Event: emptyEvent,
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sqlite storage.Storer = sqlite.New(nil)
			got, err := sqlite.StoreEvent(tt.Event)
			if got != tt.want {
				t.Errorf("Got: %v, Want: %v. Error: %v", got, tt.want, err)
			}
		})
	}
}

func TestStoreJob(t *testing.T) {
	validJob := models.Job{
		EventId:       1,
		Status:        "waiting",
		CreatedAt:     time.Now(),
		ScheduledAt:   time.Now(),
		StartedAt:     nil,
		ExecutionTime: nil,
	}

	emptyStatusJob := models.Job{
		EventId:       1,
		Status:        "",
		CreatedAt:     time.Now(),
		ScheduledAt:   time.Now(),
		StartedAt:     nil,
		ExecutionTime: nil,
	}

	tests := []struct {
		name string
		job  models.Job
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
			var sqlite storage.Storer = sqlite.New(nil)
			got, err := sqlite.StoreJob(tt.job)
			if got != tt.want {
				t.Errorf("Got %v, Want %v. Error: %v", got, tt.want, err)
			}
		})
	}
}

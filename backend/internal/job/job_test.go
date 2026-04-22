package job_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sapienfrom2000s/trident/backend/internal/core/models"
	"github.com/sapienfrom2000s/trident/backend/internal/job"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateJobHandler(t *testing.T) {
	tests := []struct {
		name            string
		payload         []byte
		cloneRepoErr    error
		wantStatusCode  int
		wantJobsCount   int64
		wantEventsCount int64
	}{
		{
			name: "spawn a job",
			payload: []byte(`{
				"triggeredBy": "Mark",
				"repoName": "octocat/Hello-World",
				"branch": "main",
				"authorName": "Monalisa Octocat"
			}`),
			wantStatusCode:  http.StatusOK,
			wantJobsCount:   1,
			wantEventsCount: 1,
		},
		{
			name: "reject invalid payload",
			payload: []byte(`{
				"triggeredBy": "Mark",
			}`),
			wantStatusCode:  http.StatusBadRequest,
			wantJobsCount:   0,
			wantEventsCount: 0,
		},
		{
			name: "clone failure",
			payload: []byte(`{
				"triggeredBy": "Mark",
				"repoName": "octocat/Hello-World",
				"branch": "main",
				"authorName": "Monalisa Octocat"
			}`),
			cloneRepoErr:    errors.New("clone failed"),
			wantStatusCode:  http.StatusInternalServerError,
			wantJobsCount:   0,
			wantEventsCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := setupTestDB(t)

			req := httptest.NewRequest(http.MethodPost, "/jobs", bytes.NewBuffer(tt.payload))
			req.Header.Add("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			h := job.Handler{
				DB: db,
				CloneRepo: func(token, branch, repoName, destination string) error {
					return tt.cloneRepoErr
				},
			}

			h.CreateJobHandler(rec, req)

			if rec.Result().StatusCode != tt.wantStatusCode {
				t.Errorf("Expected Status Code: %v, Got: %v", tt.wantStatusCode, rec.Result().StatusCode)
			}

			var jobsCount, eventsCount int64
			db.Model(&models.Job{}).Count(&jobsCount)
			db.Model(&models.Event{}).Count(&eventsCount)

			if jobsCount != tt.wantJobsCount {
				t.Errorf("Expected Job Count: %v, Got %v", tt.wantJobsCount, jobsCount)
			}

			if eventsCount != tt.wantEventsCount {
				t.Errorf("Expected Events Count: %v, Got %v", tt.wantEventsCount, eventsCount)
			}
		})
	}
}

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}

	err = db.AutoMigrate(&models.Event{}, &models.Job{})
	if err != nil {
		t.Fatalf("failed to migrate db. err: %v", err)
	}

	t.Cleanup(func() {
		sqlDB, err := db.DB()
		if err == nil {
			sqlDB.Close()
		}
	})

	return db
}

func TestRunJob(t *testing.T) {
	type tests struct {
		name    string
		params  []string
		wantErr bool
	}

	suite := []tests{
		{
			name:    "Single Command",
			params:  []string{"echo Hello World"},
			wantErr: false,
		},
		{
			name:    "Bad Command",
			params:  []string{"eche Hello World"},
			wantErr: true,
		},
		{
			name:    "Multiple Commands",
			params:  []string{"echo Hello World", "ls"},
			wantErr: false,
		},
	}

	for _, test := range suite {
		t.Run(test.name, func(t *testing.T) {
			err := job.RunJob(test.params)

			gotErr := err != nil

			if test.wantErr != gotErr {
				t.Errorf("Want %v, got %v", test.wantErr, gotErr)
			}
		})
	}
}

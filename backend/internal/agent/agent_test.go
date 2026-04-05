package agent_test

import (
	"testing"

	"github.com/sapienfrom2000s/trident/backend/internal/agent"
)

func TestRunJob(t *testing.T) {
	tests := []struct {
		name     string
		repoName string
		branch   string
		jobId    int
		wantErr  bool
	}{
		{
			name:     "Valid Job",
			repoName: "foo/bar",
			branch:   "bla",
			jobId:    1,
			wantErr:  false,
		},
		{
			name:     "Invalid Job: Repo Name is absent",
			repoName: "",
			branch:   "bla",
			jobId:    1,
			wantErr:  true,
		},
		{
			name:     "Invalid Job: Branch Name is absent",
			repoName: "foo/bar",
			branch:   "",
			jobId:    1,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := agent.RunJob(tt.repoName, tt.branch, tt.jobId)
			gotErr := err != nil

			if gotErr != tt.wantErr {
				t.Errorf("Got: %v, Want: %v", gotErr, tt.wantErr)

			}
		})
	}
}

func TestSendHeartBeat(t *testing.T) {
	tests := []struct {
		name    string
		jobId   int
		wantErr bool
	}{
		{
			name:    "Send HeartBeat",
			jobId:   1,
			wantErr: false,
		},
		{
			name:    "Invalid Job Id",
			jobId:   0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := agent.SendHeartBeat(tt.jobId)
			gotErr := err != nil

			if gotErr != tt.wantErr {
				t.Errorf("Got: %v, Want: %v", gotErr, tt.wantErr)

			}
		})
	}
}

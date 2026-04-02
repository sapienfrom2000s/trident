package github_test

import (
	"testing"

	"github.com/sapienfrom2000s/trident/backend/internal/webhook"
	"github.com/sapienfrom2000s/trident/backend/internal/webhook/github"
)

func TestParseGithubPayload(t *testing.T) {
	validPayload := []byte(`
		{
			"after": "9fceb02d0ae5",
			"ref": "refs/heads/main",
			"repository": {
		    "full_name": "octocat/Hello-World",
		  },
		  "head_commit": {
		    "author": {
		      "name": "Monalisa Octocat",
		    }
		  }
		}
    `)

	nonJsonPayload := []byte(`someplaintext`)
	fieldsAbsentInPayload := []byte(`{}`)

	tests := []struct {
		name    string
		payload []byte
		want    webhook.NormalizedEvent
		err     bool
	}{
		{
			name:    "valid json",
			payload: validPayload,
			want: webhook.NormalizedEvent{
				RepoName:  "octocat/Hello-World",
				CommitSha: "9fceb02d0ae5",
				Branch:    "main",
				Author:    "Monalisa Octocat",
				Provider:  "github",
			},
			err: false,
		},
		{
			name:    "invalid json payload",
			payload: nonJsonPayload,
			err:     true,
		},
		{
			name:    "required fields absent",
			payload: fieldsAbsentInPayload,
			err:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := github.ParseEvent(tt.payload)
			errBool := (err != nil)
			if tt.err != errBool {
				t.Errorf("Error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

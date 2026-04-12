package github_test

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sapienfrom2000s/trident/backend/internal/core/models"
	"github.com/sapienfrom2000s/trident/backend/internal/webhook"
	"github.com/sapienfrom2000s/trident/backend/internal/webhook/github"
)

func signPayload(secret string, payload []byte) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	return fmt.Sprintf("sha256=%s", hex.EncodeToString(mac.Sum(nil)))
}

func TestValidateSignature(t *testing.T) {
	secret := "mysecret"
	payload := []byte(`{"after":"9fceb02d0ae5"}`)

	validHeaders := http.Header{}
	validHeaders.Set("X-Hub-Signature-256", signPayload(secret, payload))

	invalidSigHeaders := http.Header{}
	invalidSigHeaders.Set("X-Hub-Signature-256", "sha256=invalidsignature")

	tests := []struct {
		name    string
		payload []byte
		headers http.Header
		secret  string
		wantErr bool
	}{
		{
			name:    "valid signature",
			payload: payload,
			headers: validHeaders,
			secret:  secret,
			wantErr: false,
		},
		{
			name:    "invalid signature",
			payload: payload,
			headers: invalidSigHeaders,
			secret:  secret,
			wantErr: true,
		},
		{
			name:    "missing signature header",
			payload: payload,
			headers: http.Header{},
			secret:  secret,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := github.ValidateSignature(tt.payload, tt.headers, tt.secret)
			gotErr := err != nil
			if gotErr != tt.wantErr {
				t.Errorf("got err=%v, want err=%v: %v", gotErr, tt.wantErr, err)
			}
		})
	}
}

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

	pushHeaders := http.Header{}
	pushHeaders.Set("X-GitHub-Event", "push")

	tests := []struct {
		name    string
		payload []byte
		headers http.Header
		want    models.Event
		err     bool
	}{
		{
			name:    "valid json",
			payload: validPayload,
			headers: pushHeaders,
			want: models.Event{
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
			headers: pushHeaders,
			err:     true,
		},
		{
			name:    "required fields absent",
			payload: fieldsAbsentInPayload,
			headers: pushHeaders,
			err:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := github.ParseEvent(tt.payload, tt.headers)
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

func TestGithubWebhookHandler(t *testing.T) {
	validPayload := []byte(`
		{
			"after": "9fceb02d0ae5",
			"ref": "refs/heads/main",
			"repository": {
		    "full_name": "octocat/Hello-World"
		  },
		  "head_commit": {
		    "author": {
		      "name": "Monalisa Octocat"
		    }
		  }
		}
    `)
	body := bytes.NewBuffer(validPayload)
	req := httptest.NewRequest(http.MethodPost, "/webhook/github", body)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	var gh webhook.Webhook = &github.Handler{
		ValidateSignature: func(b []byte, h http.Header, s string) error { return nil },
	}

	gh.WebhookHandler(rec, req)
	t.Run("Check for Status Code 200", func(t *testing.T) {
		if rec.Result().StatusCode != 200 {
			t.Errorf("Expected: 200, Got: %v", rec.Result().StatusCode)
		}
	})
}

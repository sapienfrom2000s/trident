package server_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sapienfrom2000s/trident/backend/internal/server"
)

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
	// mock Github secret verification
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	server.GithubWebhookHandler(rec, req)
	t.Run("Check for Status Code 200", func(t *testing.T) {
		if rec.Result().StatusCode != 200 {
			t.Errorf("Expected: 200, Got: %v", rec.Result().StatusCode)
		}
	})
	// also check if this increases job count by 1 in db
}

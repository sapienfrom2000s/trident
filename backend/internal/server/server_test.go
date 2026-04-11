package server_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sapienfrom2000s/trident/backend/internal/server"
)

func TestAgentHeartBeat(t *testing.T) {
	payload := []byte(`
		{
			"timestamp": "some-time",
		}
    `)
	body := bytes.NewBuffer(payload)
	// As of now security issue to accept heartbeat w/o any authentication
	req := httptest.NewRequest(http.MethodPost, "/heartbeat", body)
	rec := httptest.NewRecorder()

	server.HeartBeatHandler(rec, req)
	t.Run("Healthy Hearbeat", func(t *testing.T) {
		statusCode := rec.Result().StatusCode
		if statusCode != 200 {
			t.Errorf("Expected: 200, Got: %v", statusCode)
		}
	})
}

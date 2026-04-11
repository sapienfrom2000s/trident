package github

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sapienfrom2000s/trident/backend/internal/core"
)

type Handler struct {
	ValidateSignature func([]byte, http.Header, string) error
}

type GitHubPushEvent struct {
	After      string `json:"after"`
	Ref        string `json:"ref"`
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
	HeadCommit struct {
		Author struct {
			Name string `json:"name"`
		} `json:"author"`
	} `json:"head_commit"`
}

func (h *Handler) WebhookHandler(w http.ResponseWriter, r *http.Request) {
	var pushEvent GitHubPushEvent = GitHubPushEvent{}
	pushEventInBytes, err := io.ReadAll(r.Body)
	if err != nil {
		json.Unmarshal(pushEventInBytes, &pushEvent)
	}
}

func ValidateSignature(b []byte, headers http.Header, secret string) error {
	sig := headers.Get("X-Hub-Signature-256")
	if sig == "" {
		return fmt.Errorf("missing X-Hub-Signature-256 header")
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(b)
	expected := fmt.Sprintf("sha256=%s", hex.EncodeToString(mac.Sum(nil)))

	if !hmac.Equal([]byte(sig), []byte(expected)) {
		return fmt.Errorf("signature mismatch")
	}
	return nil
}

func ParseEvent(b []byte, headers http.Header) (core.NormalizedEvent, error) {
	return core.NormalizedEvent{}, nil
}

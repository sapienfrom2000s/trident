package github

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sapienfrom2000s/trident/backend/internal/core/models"
	"gorm.io/gorm"
)

type Handler struct {
	ValidateSignature func([]byte, http.Header, string) error
	DB                *gorm.DB
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
	pushEventInBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusBadRequest)
		return
	}

	event, err := ParseEvent(pushEventInBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	createEvent := h.DB.Create(&event)
	if createEvent.Error != nil {
		http.Error(w, "failed to create Event", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
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

func ParseEvent(b []byte) (models.Event, error) {
	var pushEvent GitHubPushEvent

	err := json.Unmarshal(b, &pushEvent)
	if err != nil {
		return models.Event{}, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Validate required fields
	if pushEvent.Repository.FullName == "" || pushEvent.After == "" ||
		pushEvent.Ref == "" || pushEvent.HeadCommit.Author.Name == "" {
		return models.Event{}, fmt.Errorf("required fields missing in payload")
	}

	event := models.Event{
		RepoName:  pushEvent.Repository.FullName,
		CommitSha: pushEvent.After,
		Branch:    pluckBranchFromRef(pushEvent.Ref),
		Author:    pushEvent.HeadCommit.Author.Name,
		Provider:  "github",
	}

	return event, nil
}

func pluckBranchFromRef(ref string) string {
	const refPrefix = "refs/heads/"
	return strings.TrimPrefix(ref, refPrefix)
}

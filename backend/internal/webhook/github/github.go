package github

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/sapienfrom2000s/trident/backend/internal/core"
)

type Handler struct {
	ValidateSignature func([]byte, http.Header, string) error
}

func (h *Handler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
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

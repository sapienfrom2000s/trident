package webhook

import "net/http"

// It's an overkill for a single storage backend.
// Doing this for learning purposes
type Webhook interface {
	WebhookHandler(w http.ResponseWriter, r *http.Request)
}

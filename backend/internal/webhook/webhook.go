package webhook

import "net/http"

// It's an overkill for a single storage backend.
// Doing this for learning purposes
type Webhook interface {
	HandleWebhook(w http.ResponseWriter, r *http.Request)
}

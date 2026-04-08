package storage

import (
	"github.com/sapienfrom2000s/trident/backend/internal/core"
)

// It's an overkill for a single storage backend.
// Doing this for learning purposes
type Storer interface {
	StoreEvent(core.NormalizedEvent) (bool, error)
	StoreJob(core.Job) (bool, error)
}

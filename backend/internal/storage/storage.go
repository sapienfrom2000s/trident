package storage

import (
	"github.com/sapienfrom2000s/trident/backend/internal/core/models"
)

// It's an overkill for a single storage backend.
// Doing this for learning purposes
type Storer interface {
	StoreEvent(models.Event) (bool, error)
	StoreJob(models.Job) (bool, error)
}

package sqlite

import (
	"fmt"

	"github.com/sapienfrom2000s/trident/backend/internal/core"
)

func StoreEvent(n core.NormalizedEvent) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

func StoreJob(j core.Job) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

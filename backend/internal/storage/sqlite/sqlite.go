package sqlite

import (
	"fmt"

	"github.com/sapienfrom2000s/trident/backend/internal/server"
	"github.com/sapienfrom2000s/trident/backend/internal/webhook"
)

func StoreEvent(n webhook.NormalizedEvent) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

func StoreJob(j server.Job) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

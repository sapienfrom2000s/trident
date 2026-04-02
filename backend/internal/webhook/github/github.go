package github

import "github.com/sapienfrom2000s/trident/backend/internal/webhook"

func ParseEvent(b []byte) (webhook.NormalizedEvent, error) {
	return webhook.NormalizedEvent{}, nil
}

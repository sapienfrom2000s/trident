package github

import (
	"testing"
)

var githubPayload = []byte(`
		"X-GitHub-Hook-ID": "some-hook-id",
		"X-GitHub-Event": "some-event",
		"X-Github-Delivery": "some-delivery-id",
		"X-Hub-Signature": "some-signature-id",
		"
	`)

func TestParseGithubPayload(t *testing.T) {

}

type Event interface {
	Parse(b []byte) (NormalizedEvent, error)
	Validate(n NormalizedEvent) error
}

type NormalizedEvent struct {
	repo_url     string
	repo_name    string
	commit_sha   string
	branch       string
	triggered_by string
	event_type   string
	provider     string
	metadata     map[string]string
}


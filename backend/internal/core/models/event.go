package models

type Event struct {
	RepoName  string
	CommitSha string
	Branch    string
	Author    string
	Provider  string
}

package webhook

import "errors"

var ErrMissingCommitSha = errors.New("missing commit sha")
var ErrMissingRepoName = errors.New("missing repo name")
var ErrMissingBranchName = errors.New("missing branch name")
var Author = errors.New("missing author name")
var Provider = errors.New("missing provider name")

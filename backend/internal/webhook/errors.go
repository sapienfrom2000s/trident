package webhook

import "errors"

var ErrMissingCommitSha = errors.New("missing commit sha")
var ErrMissingRepoName = errors.New("missing repo name")
var ErrMissingBranchName = errors.New("missing branch name")
var ErrMissingAuthorName = errors.New("missing author name")
var ErrMissingProviderName = errors.New("missing provider name")

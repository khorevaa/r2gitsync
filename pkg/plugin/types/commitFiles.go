package types

import (
	"github.com/khorevaa/r2gitsync/internal/manager/types"
	"time"
)

//CommitFiles(v8end V8Endpoint, dir string, author RepositoryAuthor, when time.Time, comment string) error

type CommitFilesSubscriber struct {
	Before BeforeCommitFilesFn
	On     OnCommitFilesFn
	After  AfterCommitFilesFn
}
type (
	BeforeCommitFilesFn func(v8end V8Endpoint, workdir string, author types.RepositoryAuthor, when time.Time, comment string) error
	OnCommitFilesFn     func(v8end V8Endpoint, workdir string, author types.RepositoryAuthor, when *time.Time, comment *string, stdHandler *bool) error
	AfterCommitFilesFn  BeforeCommitFilesFn
)

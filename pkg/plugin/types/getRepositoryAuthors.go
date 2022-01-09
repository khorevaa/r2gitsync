package types

import (
	"github.com/khorevaa/r2gitsync/internal/manager/types"
)

//GetRepositoryAuthors(v8end V8Endpoint, dir string, filenme string) (map[string]RepositoryAuthor, error)

type GetRepositoryAuthorsSubscriber struct {
	Before BeforeGetRepositoryAuthorsFn
	On     OnGetRepositoryAuthorsFn
	After  AfterGetRepositoryAuthorsFn
}
type (
	BeforeGetRepositoryAuthorsFn func(v8end V8Endpoint, workdir string, filename string) error
	OnGetRepositoryAuthorsFn     func(v8end V8Endpoint, workdir string, filename string, stdHandler *bool) (map[string]types.RepositoryAuthor, error)
	AfterGetRepositoryAuthorsFn  func(v8end V8Endpoint, workdir string, authors *types.RepositoryAuthorsList) error
)

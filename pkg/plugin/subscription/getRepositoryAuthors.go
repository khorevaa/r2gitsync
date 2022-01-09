package subscription

import (
	"github.com/khorevaa/r2gitsync/internal/manager/types"
	. "github.com/khorevaa/r2gitsync/pkg/plugin/types"
)

var _ GetRepositoryAuthorsHandler = (*getRepositoryAuthorsHandler)(nil)

type GetRepositoryAuthorsHandler interface {
	SubscribeHandler
	Subscribe(sub GetRepositoryAuthorsSubscriber)

	Before(v8end V8Endpoint, workdir string, filename string) error
	On(v8end V8Endpoint, workdir string, filename string, stdHandler *bool) (map[string]types.RepositoryAuthor, error)
	After(v8end V8Endpoint, workdir string, authors *types.RepositoryAuthorsList) error
}

type getRepositoryAuthorsHandler struct {
	before []BeforeGetRepositoryAuthorsFn
	on     []OnGetRepositoryAuthorsFn
	after  []AfterGetRepositoryAuthorsFn
}

func (h *getRepositoryAuthorsHandler) Subscribe(sub GetRepositoryAuthorsSubscriber) {

	if sub.Before != nil {
		h.before = append(h.before, sub.Before)
	}

	if sub.On != nil {
		h.on = append(h.on, sub.On)
	}

	if sub.After != nil {
		h.after = append(h.after, sub.After)
	}
}

func (h *getRepositoryAuthorsHandler) Before(v8end V8Endpoint, workdir string, filename string) error {

	for _, fn := range h.before {

		err := fn(v8end, workdir, filename)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *getRepositoryAuthorsHandler) On(v8end V8Endpoint, workdir string, filename string, stdHandler *bool) (map[string]types.RepositoryAuthor, error) {

	for _, fn := range h.on {

		rv, err := fn(v8end, workdir, filename, stdHandler)

		if err != nil {
			return rv, err
		}
	}

	return map[string]types.RepositoryAuthor{}, nil
}

func (h *getRepositoryAuthorsHandler) After(v8end V8Endpoint, workdir string, authors *types.RepositoryAuthorsList) error {

	for _, fn := range h.after {

		err := fn(v8end, workdir, authors)

		if err != nil {
			return err
		}
	}

	return nil
}

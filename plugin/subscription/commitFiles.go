package subscription

import (
	"github.com/khorevaa/r2gitsync/manager/types"
	. "github.com/khorevaa/r2gitsync/plugin/types"
	"time"
)

var _ CommitFilesHandler = (*commitFilesHandler)(nil)

type CommitFilesHandler interface {
	SubscribeHandler

	Subscribe(sub CommitFilesSubscriber)

	Before(v8end V8Endpoint, workdir string, author types.RepositoryAuthor, when time.Time, comment string) error
	On(v8end V8Endpoint, workdir string, author types.RepositoryAuthor, when *time.Time, comment *string, stdHandler *bool) error
	After(v8end V8Endpoint, workdir string, author types.RepositoryAuthor, when time.Time, comment string) error
}

type commitFilesHandler struct {
	before []BeforeCommitFilesFn
	on     []OnCommitFilesFn
	after  []AfterCommitFilesFn
}

func (h *commitFilesHandler) Subscribe(sub CommitFilesSubscriber) {

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

func (h *commitFilesHandler) Before(v8end V8Endpoint, workdir string, author types.RepositoryAuthor, when time.Time, comment string) error {

	for _, fn := range h.before {

		err := fn(v8end, workdir, author, when, comment)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *commitFilesHandler) On(v8end V8Endpoint, workdir string, author types.RepositoryAuthor, when *time.Time, comment *string, stdHandler *bool) error {

	for _, fn := range h.on {

		err := fn(v8end, workdir, author, when, comment, stdHandler)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *commitFilesHandler) After(v8end V8Endpoint, workdir string, author types.RepositoryAuthor, when time.Time, comment string) error {

	for _, fn := range h.after {

		err := fn(v8end, workdir, author, when, comment)

		if err != nil {
			return err
		}
	}

	return nil
}

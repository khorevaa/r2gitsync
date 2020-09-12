package subscription

import . "github.com/khorevaa/r2gitsync/plugin/types"

var _ GetRepositoryHistoryHandler = (*getRepositoryHistoryHandler)(nil)

type getRepositoryHistoryHandler struct {
	before []BeforeDumpConfigFn
	on     []OnDumpConfigFn
	after  []AfterDumpConfigFn
}

func (h *getRepositoryHistoryHandler) Subscribe(sub Subscriber) {
	panic("implement me")
}

type GetRepositoryHistoryHandler interface {
	SubscribeHandler
	Before(v8end V8Endpoint, workdir string, temp string, number int64) error
	On(v8end V8Endpoint, workdir string, temp string, number int64, standartHandler *bool) error
	After(v8end V8Endpoint, workdir string, temp string, number int64) error
}

func (h *getRepositoryHistoryHandler) Before(v8end V8Endpoint, workdir string, temp string, version int64) error {

	for _, fn := range h.before {

		err := fn(v8end, workdir, temp, version)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *getRepositoryHistoryHandler) On(v8end V8Endpoint, workdir string, temp string, version int64, standartHandler *bool) error {

	for _, fn := range h.on {

		err := fn(v8end, workdir, temp, version, standartHandler)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *getRepositoryHistoryHandler) After(v8end V8Endpoint, workdir string, temp string, version int64) error {

	for _, fn := range h.after {

		err := fn(v8end, workdir, temp, version)

		if err != nil {
			return err
		}
	}

	return nil
}

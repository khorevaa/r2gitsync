package subscription

import (
	"github.com/khorevaa/r2gitsync/internal/manager/types"
	. "github.com/khorevaa/r2gitsync/pkg/plugin/types"
)

var _ GetRepositoryHistoryHandler = (*getRepositoryHistoryHandler)(nil)

type getRepositoryHistoryHandler struct {
	before []BeforeGetRepositoryHistoryFn
	on     []OnGetRepositoryHistoryFn
	after  []AfterGetRepositoryHistoryFn
}

func (h *getRepositoryHistoryHandler) Subscribe(sub GetRepositoryHistorySubscriber) {

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

type GetRepositoryHistoryHandler interface {
	SubscribeHandler
	Subscribe(sub GetRepositoryHistorySubscriber)

	Before(v8end V8Endpoint, dir string, NBegin int) error
	On(v8end V8Endpoint, dir string, NBegin int, stdHandler *bool) ([]types.RepositoryVersion, error)
	After(v8end V8Endpoint, dir string, NBegin int, rv *types.RepositoryVersionsList) error

	//Start(v8end V8Endpoint, workdir string, temp string, number int) error
	//On(v8end V8Endpoint, workdir string, temp string, number int, standartHandler *bool) error
	//Finish(v8end V8Endpoint, workdir string, temp string, number int) error
}

func (h *getRepositoryHistoryHandler) Before(v8end V8Endpoint, dir string, NBegin int) error {

	for _, fn := range h.before {

		err := fn(v8end, dir, NBegin)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *getRepositoryHistoryHandler) On(v8end V8Endpoint, dir string, NBegin int, stdHandler *bool) ([]types.RepositoryVersion, error) {

	for _, fn := range h.on {

		rv, err := fn(v8end, dir, NBegin, stdHandler)

		if err != nil {
			return rv, err
		}
	}

	return []types.RepositoryVersion{}, nil
}

func (h *getRepositoryHistoryHandler) After(v8end V8Endpoint, dir string, NBegin int, rv *types.RepositoryVersionsList) error {

	for _, fn := range h.after {

		err := fn(v8end, dir, NBegin, rv)

		if err != nil {
			return err
		}
	}

	return nil
}

var _ ConfigureRepositoryVersionsHandler = (*configureRepositoryVersionsHandler)(nil)

type configureRepositoryVersionsHandler struct {
	on []OnConfigureRepositoryVersionsFn
}

func (h *configureRepositoryVersionsHandler) Subscribe(sub ConfigureRepositoryVersionsSubscriber) {

	if sub.On != nil {
		h.on = append(h.on, sub.On)
	}

}

type ConfigureRepositoryVersionsHandler interface {
	SubscribeHandler
	Subscribe(sub ConfigureRepositoryVersionsSubscriber)
	On(v8end V8Endpoint, versions *types.RepositoryVersionsList, NCurrent, NNext, NMax *int) error
}

func (h *configureRepositoryVersionsHandler) On(v8end V8Endpoint, versions *types.RepositoryVersionsList, NCurrent, NNext, NMax *int) error {

	for _, fn := range h.on {

		err := fn(v8end, versions, NCurrent, NNext, NMax)

		if err != nil {
			return err
		}
	}

	return nil
}

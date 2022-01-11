package subscription

import (
	. "github.com/khorevaa/r2gitsync/pkg/plugin/types"
)

var _ UpdateCfgHandler = (*updateCfgHandler)(nil)

type UpdateCfgHandler interface {
	SubscribeHandler
	Subscribe(sub UpdateCfgSubscriber)
	Before(v8end V8Endpoint, workdir string, number int) error
	On(v8end V8Endpoint, workdir string, number int, standartHandler *bool) error
	After(v8end V8Endpoint, workdir string, number int) error

	BeforeFn(v8end V8Endpoint, workdir string, number int) func() error
}

type updateCfgHandler struct {
	before []BeforeUpdateCfgFn
	on     []OnUpdateCfgFn
	after  []AfterUpdateCfgFn
}

func (h *updateCfgHandler) Count() int {
	return len(h.on) + len(h.after) + len(h.before)
}

func (h *updateCfgHandler) BeforeFn(v8end V8Endpoint, workdir string, number int) func() error {
	return func() error {
		return h.Before(v8end, workdir, number)
	}
}

func (h *updateCfgHandler) Subscribe(sub UpdateCfgSubscriber) {

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

func (h *updateCfgHandler) Before(v8end V8Endpoint, workdir string, version int) error {

	for _, fn := range h.before {

		err := fn(v8end, workdir, version)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *updateCfgHandler) On(v8end V8Endpoint, workdir string, version int, stdHandler *bool) error {

	for _, fn := range h.on {

		err := fn(v8end, workdir, version, stdHandler)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *updateCfgHandler) After(v8end V8Endpoint, workdir string, version int) error {

	for _, fn := range h.after {

		err := fn(v8end, workdir, version)

		if err != nil {
			return err
		}
	}

	return nil
}

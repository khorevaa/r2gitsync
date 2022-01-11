package subscription

import (
	. "github.com/khorevaa/r2gitsync/pkg/plugin/types"
)

var _ ClearWorkdirHandler = (*clearWorkdirHandler)(nil)

type ClearWorkdirHandler interface {
	SubscribeHandler

	Subscribe(sub ClearWorkdirSubscriber)

	Before(v8end V8Endpoint, workdir, temp string) error
	On(v8end V8Endpoint, workdir, temp string, stdHandler *bool) error
	After(v8end V8Endpoint, workdir, temp string) error
}

type clearWorkdirHandler struct {
	before []BeforeClearWorkdirFn
	on     []OnClearWorkdirFn
	after  []AfterClearWorkdirFn
}

func (h *clearWorkdirHandler) Count() int {
	return len(h.on) + len(h.after) + len(h.before)
}

func (h *clearWorkdirHandler) Subscribe(sub ClearWorkdirSubscriber) {

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

func (h *clearWorkdirHandler) Before(v8end V8Endpoint, workdir, temp string) error {

	for _, fn := range h.before {

		err := fn(v8end, workdir, temp)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *clearWorkdirHandler) On(v8end V8Endpoint, workdir, temp string, stdHandler *bool) error {

	for _, fn := range h.on {

		err := fn(v8end, workdir, temp, stdHandler)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *clearWorkdirHandler) After(v8end V8Endpoint, workdir, temp string) error {

	for _, fn := range h.after {

		err := fn(v8end, workdir, temp)

		if err != nil {
			return err
		}
	}

	return nil
}

var _ MoveToWorkdirHandler = (*moveToWorkdirHandler)(nil)

type MoveToWorkdirHandler interface {
	SubscribeHandler

	Subscribe(sub MoveToWorkdirSubscriber)

	Before(v8end V8Endpoint, workdir, temp string) error
	On(v8end V8Endpoint, workdir, temp string, stdHandler *bool) error
	After(v8end V8Endpoint, workdir, temp string) error
}

type moveToWorkdirHandler struct {
	before []BeforeMoveToWorkdirFn
	on     []OnMoveToWorkdirFn
	after  []AfterMoveToWorkdirFn
}

func (h *moveToWorkdirHandler) Count() int {
	return len(h.on) + len(h.after) + len(h.before)
}

func (h *moveToWorkdirHandler) Subscribe(sub MoveToWorkdirSubscriber) {

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

func (h *moveToWorkdirHandler) Before(v8end V8Endpoint, workdir, temp string) error {

	for _, fn := range h.before {

		err := fn(v8end, workdir, temp)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *moveToWorkdirHandler) On(v8end V8Endpoint, workdir, temp string, stdHandler *bool) error {

	for _, fn := range h.on {

		err := fn(v8end, workdir, temp, stdHandler)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *moveToWorkdirHandler) After(v8end V8Endpoint, workdir, temp string) error {

	for _, fn := range h.after {

		err := fn(v8end, workdir, temp)

		if err != nil {
			return err
		}
	}

	return nil
}

package subscription

import (
	. "github.com/khorevaa/r2gitsync/plugin/types"
)

var _ ReadVersionFileHandler = (*readVersionFileHandler)(nil)

type ReadVersionFileHandler interface {
	SubscribeHandler

	Subscribe(sub ReadVersionFileSubscriber)

	Before(v8end V8Endpoint, workdir, filename string) error
	On(v8end V8Endpoint, workdir, filename string, stdHandler *bool) (int64, error)
	After(v8end V8Endpoint, workdir, filename string, number *int64) error
}

type readVersionFileHandler struct {
	before []BeforeReadVersionFileFn
	on     []OnReadVersionFileFn
	after  []AfterReadVersionFileFn
}

func (h *readVersionFileHandler) Subscribe(sub ReadVersionFileSubscriber) {

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

func (h *readVersionFileHandler) Before(v8end V8Endpoint, workdir, filename string) error {

	for _, fn := range h.before {

		err := fn(v8end, workdir, filename)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *readVersionFileHandler) On(v8end V8Endpoint, workdir, filename string, stdHandler *bool) (int64, error) {

	for _, fn := range h.on {

		n, err := fn(v8end, workdir, filename, stdHandler)

		if err != nil {
			return n, err
		}
	}

	return 0, nil
}

func (h *readVersionFileHandler) After(v8end V8Endpoint, workdir, filename string, number *int64) error {

	for _, fn := range h.after {

		err := fn(v8end, workdir, filename, number)

		if err != nil {
			return err
		}
	}

	return nil
}

var _ WriteVersionFileHandler = (*writeVersionFileHandler)(nil)

type WriteVersionFileHandler interface {
	SubscribeHandler

	Subscribe(sub WriteVersionFileSubscriber)

	Before(v8end V8Endpoint, workdir string, number int64, filename string) error
	On(v8end V8Endpoint, workdir string, number int64, filename string, stdHandler *bool) error
	After(v8end V8Endpoint, workdir string, number int64, filename string) error
}

type writeVersionFileHandler struct {
	before []BeforeWriteVersionFileFn
	on     []OnWriteVersionFileFn
	after  []AfterWriteVersionFileFn
}

func (h *writeVersionFileHandler) Subscribe(sub WriteVersionFileSubscriber) {

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

func (h *writeVersionFileHandler) Before(v8end V8Endpoint, workdir string, number int64, filename string) error {

	for _, fn := range h.before {

		err := fn(v8end, workdir, number, filename)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *writeVersionFileHandler) On(v8end V8Endpoint, workdir string, number int64, filename string, stdHandler *bool) error {

	for _, fn := range h.on {

		err := fn(v8end, workdir, number, filename, stdHandler)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *writeVersionFileHandler) After(v8end V8Endpoint, workdir string, number int64, filename string) error {

	for _, fn := range h.after {

		err := fn(v8end, workdir, number, filename)

		if err != nil {
			return err
		}
	}

	return nil
}

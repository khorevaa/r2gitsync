package subscription

import . "github.com/khorevaa/r2gitsync/plugin/types"

var _ DumpConfigToFilesHandler = (*dumpConfigToFilesHandler)(nil)

type dumpConfigToFilesHandler struct {
	before []BeforeDumpConfigFn
	on     []OnDumpConfigFn
	after  []AfterDumpConfigFn
}

func (h *dumpConfigToFilesHandler) Subscribe(sub DumpConfigToFilesSubscriber) {

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

type DumpConfigToFilesHandler interface {
	SubscribeHandler

	Subscribe(sub DumpConfigToFilesSubscriber)
	Before(v8end V8Endpoint, workdir string, temp string, number int, update *bool) error
	On(v8end V8Endpoint, workdir string, temp string, number int, update *bool, stdHandler *bool) error
	After(v8end V8Endpoint, workdir string, temp string, number int, update bool) error
}

func (h *dumpConfigToFilesHandler) Before(v8end V8Endpoint, workdir string, temp string, version int, update *bool) error {

	for _, fn := range h.before {

		err := fn(v8end, workdir, temp, version, update)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *dumpConfigToFilesHandler) On(v8end V8Endpoint, workdir string, temp string, version int, update *bool, stdHandler *bool) error {

	for _, fn := range h.on {

		err := fn(v8end, workdir, temp, version, update, stdHandler)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *dumpConfigToFilesHandler) After(v8end V8Endpoint, workdir string, temp string, version int, update bool) error {

	for _, fn := range h.after {

		err := fn(v8end, workdir, temp, version, update)

		if err != nil {
			return err
		}
	}

	return nil
}

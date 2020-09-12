package subscription

import . "github.com/khorevaa/r2gitsync/plugin/types"

var _ DumpConfigToFilesHandler = (*dumpConfigToFilesHandler)(nil)

type dumpConfigToFilesHandler struct {
	before []BeforeDumpConfigFn
	on     []OnDumpConfigFn
	after  []AfterDumpConfigFn
}

func (h *dumpConfigToFilesHandler) Subscribe(sub UpdateCfgSubscriber) {
	panic("implement me")
}

type DumpConfigToFilesHandler interface {
	SubscribeHandler

	Subscribe(sub UpdateCfgSubscriber)
	Before(v8end V8Endpoint, workdir string, temp string, number int64) error
	On(v8end V8Endpoint, workdir string, temp string, number int64, standartHandler *bool) error
	After(v8end V8Endpoint, workdir string, temp string, number int64) error
}

func (h *dumpConfigToFilesHandler) Before(v8end V8Endpoint, workdir string, temp string, version int64) error {

	for _, fn := range h.before {

		err := fn(v8end, workdir, temp, version)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *dumpConfigToFilesHandler) On(v8end V8Endpoint, workdir string, temp string, version int64, standartHandler *bool) error {

	for _, fn := range h.on {

		err := fn(v8end, workdir, temp, version, standartHandler)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *dumpConfigToFilesHandler) After(v8end V8Endpoint, workdir string, temp string, version int64) error {

	for _, fn := range h.after {

		err := fn(v8end, workdir, temp, version)

		if err != nil {
			return err
		}
	}

	return nil
}

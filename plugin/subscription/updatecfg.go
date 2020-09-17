package subscription

import . "github.com/khorevaa/r2gitsync/plugin/types"

var _ UpdateCfgHandler = (*updateCfgHandler)(nil)

type UpdateCfgHandler interface {
	SubscribeHandler
	Subscribe(sub UpdateCfgSubscriber)
	Before(v8end V8Endpoint, workdir string, number int64) error
	On(v8end V8Endpoint, workdir string, number int64, standartHandler *bool) error
	After(v8end V8Endpoint, workdir string, number int64) error
}

type updateCfgHandler struct {
	before []BeforeUpdateCfgFn
	on     []OnUpdateCfgFn
	after  []AfterUpdateCfgFn
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

func (h *updateCfgHandler) Before(v8end V8Endpoint, workdir string, version int64) error {

	for _, fn := range h.before {

		err := fn(v8end, workdir, version)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *updateCfgHandler) On(v8end V8Endpoint, workdir string, version int64, standartHandler *bool) error {

	for _, fn := range h.on {

		err := fn(v8end, workdir, version, standartHandler)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *updateCfgHandler) After(v8end V8Endpoint, workdir string, version int64) error {

	for _, fn := range h.after {

		err := fn(v8end, workdir, version)

		if err != nil {
			return err
		}
	}

	return nil
}

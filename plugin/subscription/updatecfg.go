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

func (b *updateCfgHandler) Subscribe(sub UpdateCfgSubscriber) {

	updateCfg := sub

	if updateCfg.Before != nil {
		b.before = append(b.before, updateCfg.Before)
	}

	if updateCfg.On != nil {
		b.on = append(b.on, updateCfg.On)
	}

	if updateCfg.After != nil {
		b.after = append(b.after, updateCfg.After)
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

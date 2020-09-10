package subscription

var _ UpdateCfgHandler = (*updateCfgHandler)(nil)

type UpdateCfgHandler interface {
	SubscribeHandler
	Before(v8end V8Endpoint, workdir string, number int64) error
	On(v8end V8Endpoint, workdir string, number int64, standartHandler *bool) error
	After(v8end V8Endpoint, workdir string, number int64) error
}

type updateCfgHandler struct {
	before []BeforeUpdateCfgFn
	on     []OnUpdateCfgFn
	after  []AfterUpdateCfgFn
}

func (b *updateCfgHandler) Handle(event EventType, handler interface{}) {

	switch event {
	case BeforeEvent:

		fn := handler.(BeforeUpdateCfgFn)
		b.before = append(b.before, fn)

	case OnEvent:

		fn := handler.(OnUpdateCfgFn)
		b.on = append(b.on, fn)

	case AfterEvent:

		fn := handler.(AfterUpdateCfgFn)
		b.after = append(b.after, fn)

	default:
		panic("plugins: unsupported event type")
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

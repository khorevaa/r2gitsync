package Subscription

func (b *UpdateCfgHandlers) Handle(event eventType, handler interface{}) {

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

type UpdateCfgHandlers struct {
	before []BeforeUpdateCfgFn
	on     []OnUpdateCfgFn
	after  []AfterUpdateCfgFn
}

func (h *UpdateCfgHandlers) Before(v8end V8Endpoint, workdir string, version int64) error {

	for _, fn := range h.before {

		err := fn(v8end, workdir, version)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *UpdateCfgHandlers) On(v8end V8Endpoint, workdir string, version int64, standartHandler *bool) error {

	for _, fn := range h.on {

		err := fn(v8end, workdir, version, standartHandler)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h *UpdateCfgHandlers) After(v8end V8Endpoint, workdir string, version int64) error {

	for _, fn := range h.after {

		err := fn(v8end, workdir, version)

		if err != nil {
			return err
		}
	}

	return nil
}

package subscription

var _ DumpConfigToFilesHandler = (*dumpConfigToFilesHandler)(nil)

type dumpConfigToFilesHandler struct {
	before []BeforeDumpConfigFn
	on     []OnDumpConfigFn
	after  []AfterDumpConfigFn
}

type DumpConfigToFilesHandler interface {
	SubscribeHandler
	Before(v8end V8Endpoint, workdir string, temp string, number int64) error
	On(v8end V8Endpoint, workdir string, temp string, number int64, standartHandler *bool) error
	After(v8end V8Endpoint, workdir string, temp string, number int64) error
}

func (b *dumpConfigToFilesHandler) Handle(event EventType, handler interface{}) {

	switch event {
	case BeforeEvent:

		fn := handler.(BeforeDumpConfigFn)
		b.before = append(b.before, fn)

	case OnEvent:

		fn := handler.(OnDumpConfigFn)
		b.on = append(b.on, fn)

	case AfterEvent:

		fn := handler.(AfterDumpConfigFn)
		b.after = append(b.after, fn)

	default:
		panic("plugins: unsupported event type")
	}

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

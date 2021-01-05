package types

type UpdateCfgSubscriber struct {
	On     OnUpdateCfgFn
	Before BeforeUpdateCfgFn
	After  AfterUpdateCfgFn
}

type (
	BeforeUpdateCfgFn func(v8end V8Endpoint, workdir string, version int) error
	OnUpdateCfgFn     func(v8end V8Endpoint, workdir string, version int, stdHandler *bool) error
	AfterUpdateCfgFn  BeforeUpdateCfgFn
)

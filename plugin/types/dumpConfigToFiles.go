package types

//DumpConfigToFiles(v8end V8Endpoint, update bool, dir string, dir2 string, number int64) error

type DumpConfigToFilesSubscriber struct {
	On     OnDumpConfigFn
	Before BeforeDumpConfigFn
	After  AfterDumpConfigFn
}

type (
	BeforeDumpConfigFn func(v8end V8Endpoint, workdir string, temp string, number int64, update *bool) error
	OnDumpConfigFn     func(v8end V8Endpoint, workdir string, temp string, number int64, update *bool, stdHandler *bool) error
	AfterDumpConfigFn  func(v8end V8Endpoint, workdir string, temp string, number int64, update bool) error
)

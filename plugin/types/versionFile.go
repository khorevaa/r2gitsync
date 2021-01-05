package types

//WriteVersionFile(v8end V8Endpoint, dir string, number int, filename string) error

type WriteVersionFileSubscriber struct {
	Before BeforeWriteVersionFileFn
	On     OnWriteVersionFileFn
	After  AfterWriteVersionFileFn
}
type (
	BeforeWriteVersionFileFn func(v8end V8Endpoint, workdir string, number int, filename string) error
	OnWriteVersionFileFn     func(v8end V8Endpoint, workdir string, number int, filename string, stdHandler *bool) error
	AfterWriteVersionFileFn  BeforeWriteVersionFileFn
)

type ReadVersionFileSubscriber struct {
	Before BeforeReadVersionFileFn
	On     OnReadVersionFileFn
	After  AfterReadVersionFileFn
}
type (
	BeforeReadVersionFileFn func(v8end V8Endpoint, workdir string, filename string) error
	OnReadVersionFileFn     func(v8end V8Endpoint, workdir string, filename string, stdHandler *bool) (int, error)
	AfterReadVersionFileFn  func(v8end V8Endpoint, workdir string, filename string, number *int) error
)

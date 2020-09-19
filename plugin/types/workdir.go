package types

//ClearWorkDir(v8end V8Endpoint, dir string, tempDir string) error
//MoveToWorkDir(v8end V8Endpoint, dir string, tempDir string) error

type ClearWorkdirSubscriber struct {
	Before BeforeClearWorkdirFn
	On     OnClearWorkdirFn
	After  AfterClearWorkdirFn
}

type (
	BeforeClearWorkdirFn func(v8end V8Endpoint, workdir string, temp string) error
	OnClearWorkdirFn     func(v8end V8Endpoint, workdir string, temp string, stdHandler *bool) error
	AfterClearWorkdirFn  BeforeClearWorkdirFn
)

type MoveToWorkdirSubscriber struct {
	Before BeforeMoveToWorkdirFn
	On     OnMoveToWorkdirFn
	After  AfterMoveToWorkdirFn
}

type (
	BeforeMoveToWorkdirFn BeforeClearWorkdirFn
	OnMoveToWorkdirFn     OnClearWorkdirFn
	AfterMoveToWorkdirFn  BeforeClearWorkdirFn
)

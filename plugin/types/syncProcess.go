package types

//StartSyncProcess(v8end V8Endpoint, dir string)
//FinishSyncProcess(v8end V8Endpoint, dir string)

type (
	SyncProcessSubscriber struct {
		Start  StartSyncProcessFn
		Finish FinishSyncProcessFn
	}

	StartSyncProcessFn  func(v8end V8Endpoint, workdir string)
	FinishSyncProcessFn func(v8end V8Endpoint, workdir string, err *error)
)

package types

//StartSyncVersion(v8end V8Endpoint, workdir string, tempdir string, number int64) error
//FinishSyncVersion(v8end V8Endpoint, workdir string, tempdir string, number int64, err *error)

type (
	SyncVersionSubscriber struct {
		Start  StartSyncVersionFn
		Finish FinishSyncVersionFn
	}

	StartSyncVersionFn  func(v8end V8Endpoint, workdir string, tempdir string, number int)
	FinishSyncVersionFn func(v8end V8Endpoint, workdir string, tempdir string, number int, err *error)
)

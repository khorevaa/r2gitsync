package subscription

import (
	"sync"

	. "github.com/khorevaa/r2gitsync/pkg/plugin/types"
)

type SubscribeManager struct {
	mu                          sync.Mutex
	UpdateCfg                   UpdateCfgHandler
	DumpConfigToFiles           DumpConfigToFilesHandler
	GetRepositoryHistory        GetRepositoryHistoryHandler
	ConfigureRepositoryVersions ConfigureRepositoryVersionsHandler
	GetRepositoryAuthors        GetRepositoryAuthorsHandler
	SyncVersion                 SyncVersionHandler
	SyncProcess                 SyncProcessHandler
	CommitFiles                 CommitFilesHandler
	ReadVersionFile             ReadVersionFileHandler
	WriteVersionFile            WriteVersionFileHandler
	ClearWorkdir                ClearWorkdirHandler
	MoveToWorkdir               MoveToWorkdirHandler
	subscribers                 []Plugin
	count                       int
}

func (sm *SubscribeManager) Count() int {
	return sm.UpdateCfg.Count() +
		sm.DumpConfigToFiles.Count() +
		sm.GetRepositoryHistory.Count() +
		sm.ConfigureRepositoryVersions.Count() +
		sm.GetRepositoryAuthors.Count() +
		sm.SyncVersion.Count() +
		sm.SyncProcess.Count() +
		sm.CommitFiles.Count() +
		sm.ReadVersionFile.Count() +
		sm.WriteVersionFile.Count() +
		sm.ClearWorkdir.Count() +
		sm.MoveToWorkdir.Count()
}
func (sm *SubscribeManager) Subscribe(sub Subscriber) {

	handlers := sub.Handlers()

	for _, handler := range handlers {

		switch h := handler.(type) {
		case UpdateCfgSubscriber:
			sm.UpdateCfg.Subscribe(h)
		case DumpConfigToFilesSubscriber:
			sm.DumpConfigToFiles.Subscribe(h)
		case GetRepositoryHistorySubscriber:
			sm.GetRepositoryHistory.Subscribe(h)
		case GetRepositoryAuthorsSubscriber:
			sm.GetRepositoryAuthors.Subscribe(h)
		case ConfigureRepositoryVersionsSubscriber:
			sm.ConfigureRepositoryVersions.Subscribe(h)
		case SyncVersionSubscriber:
			sm.SyncVersion.Subscribe(h)
		case SyncProcessSubscriber:
			sm.SyncProcess.Subscribe(h)
		case CommitFilesSubscriber:
			sm.CommitFiles.Subscribe(h)
		case ReadVersionFileSubscriber:
			sm.ReadVersionFile.Subscribe(h)
		case WriteVersionFileSubscriber:
			sm.WriteVersionFile.Subscribe(h)
		case ClearWorkdirSubscriber:
			sm.ClearWorkdir.Subscribe(h)
		case MoveToWorkdirSubscriber:
			sm.MoveToWorkdir.Subscribe(h)
		}

	}
}

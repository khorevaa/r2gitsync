package subscription

import (
	. "github.com/khorevaa/r2gitsync/pkg/plugin/types"
)

type SubscribeHandler interface {
	Count() int
}

func NewSubscribeManager() *SubscribeManager {
	return &SubscribeManager{

		UpdateCfg:                   &updateCfgHandler{},
		DumpConfigToFiles:           &dumpConfigToFilesHandler{},
		GetRepositoryHistory:        &getRepositoryHistoryHandler{},
		GetRepositoryAuthors:        &getRepositoryAuthorsHandler{},
		ConfigureRepositoryVersions: &configureRepositoryVersionsHandler{},
		SyncProcess:                 &syncProcessHandler{},
		SyncVersion:                 &syncversionHandler{},
		ClearWorkdir:                &clearWorkdirHandler{},
		MoveToWorkdir:               &moveToWorkdirHandler{},
		CommitFiles:                 &commitFilesHandler{},
		ReadVersionFile:             &readVersionFileHandler{},
		WriteVersionFile:            &writeVersionFileHandler{},
	}
}

type Plugin interface {
	Subscribe() Subscriber
}

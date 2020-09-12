package subscription

import (
	"github.com/khorevaa/r2gitsync/context"
	. "github.com/khorevaa/r2gitsync/plugin/types"
)

type SubscribeHandler interface {
}

func NewSubscribeManager() *SubscribeManager {
	return &SubscribeManager{

		UpdateCfg:             &updateCfgHandler{},
		DumpConfigToFiles:     &dumpConfigToFilesHandler{},
		GetRepositoryHistoryH: &getRepositoryHistoryHandler{},
	}
}

type Plugin interface {
	Subscriber() Subscriber
	InitContext(ctx context.Context)
}

package Subscription

type SubscribeHandler interface {
	Handle(event eventType, handler interface{})
	//Empty() bool
}

func NewSubscribeManager() *SubscribeManager {

	return &SubscribeManager{

		UpdateCfg:             &updateCfgHandler{},
		DumpConfigToFiles:     &dumpConfigToFilesHandler{},
		GetRepositoryHistoryH: &getRepositoryHistoryHandler{},
	}

}

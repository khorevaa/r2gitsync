package subscription

type SubscribeHandler interface {
	Handle(event EventType, handler interface{})
	//Empty() bool
}

func NewSubscribeManager() *SubscribeManager {

	return &SubscribeManager{

		UpdateCfg:             &updateCfgHandler{},
		DumpConfigToFiles:     &dumpConfigToFilesHandler{},
		GetRepositoryHistoryH: &getRepositoryHistoryHandler{},
	}

}

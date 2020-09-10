package plugin

type EndPointType string

const (
	UpdateCfg            EndPointType = "\aUpdateCfg"
	DumpConfigToFiles    EndPointType = "\aDumpConfigToFiles"
	GetRepositoryHistory EndPointType = "\aGetRepositoryHistory"
)

func (t EndPointType) String() string {
	return string(t)
}

type EventType string

const (
	BeforeEvent  EventType = "\aBefore"
	OnEvent      EventType = "\aOn"
	AfterEvent   EventType = "\aAfter"
	UnknownEvent EventType = "\aUnknown"
)

func (t EventType) String() string {
	return string(t)
}

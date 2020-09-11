package subscription

import (
	"github.com/khorevaa/r2gitsync/context"
	"github.com/v8platform/designer/repository"
	v8 "github.com/v8platform/v8"
)

type V8Endpoint interface {
	Infobase() *v8.Infobase
	Repository() *repository.Repository
	Extention() string
	Options() []interface{}
}

type Plugin interface {
	Init(sm *SubscribeManager) error
	InitContext(ctx context.Context)
}

type EndPointType string

const (
	UpdateCfg            EndPointType = "\aUpdateCfg"
	DumpConfigToFiles    EndPointType = "\aDumpConfigToFiles"
	GetRepositoryHistory EndPointType = "\aGetRepositoryHistory"
)

type EventType string

const (
	BeforeEvent  EventType = "\aBefore"
	OnEvent      EventType = "\aOn"
	AfterEvent   EventType = "\aAfter"
	UnknownEvent EventType = "\aUnknown"
)

type (
	BeforeUpdateCfgFn func(v8end V8Endpoint, workdir string, version int64) error
	OnUpdateCfgFn     func(v8end V8Endpoint, workdir string, version int64, standartHandler *bool) error
	AfterUpdateCfgFn  BeforeUpdateCfgFn
)

type (
	BeforeDumpConfigFn func(v8end V8Endpoint, workdir string, temp string, number int64) error
	OnDumpConfigFn     func(v8end V8Endpoint, workdir string, temp string, number int64, standartHandler *bool) error
	AfterDumpConfigFn  BeforeDumpConfigFn
)

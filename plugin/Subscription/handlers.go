package Subscription

import (
	"github.com/v8platform/designer/repository"
	v8 "github.com/v8platform/v8"
)

type V8Endpoint interface {
	Infobase() *v8.Infobase
	Repository() *repository.Repository
	Extention() string
	Options() []interface{}
}

type endPointType string

const (
	UpdateCfg            endPointType = "\aUpdateCfg"
	DumpConfigToFiles    endPointType = "\aDumpConfigToFiles"
	GetRepositoryHistory endPointType = "\aGetRepositoryHistory"
)

type eventType string

const (
	BeforeEvent  eventType = "\aBefore"
	OnEvent      eventType = "\aOn"
	AfterEvent   eventType = "\aAfter"
	UnknownEvent eventType = "\aUnknown"
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

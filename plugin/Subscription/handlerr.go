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

type UpdateCfgEndpoint string

const (

	// Handler: func(*Message)
	BeforeUpdateCfg UpdateCfgEndpoint = "\aBeforeUpdateCfg"
	OnUpdateCfg                       = "\aOnUpdateCfg"
	AfterUpdateCfg                    = "\aAfterUpdateCfg"
)

func (h UpdateCfgEndpoint) String() string {
	return string(h)
}

type eventType string

const (
	BeforeEvent  eventType = "\aBefore"
	OnEvent      eventType = "\aOn"
	AfterEvent   eventType = "\aAfter"
	UnknownEvent eventType = "\nUnknown"
)

type (
	BeforeUpdateCfgFn func(v8end V8Endpoint, workdir string, version int64) error
	OnUpdateCfgFn     func(v8end V8Endpoint, workdir string, version int64, standartHandler *bool) error
	AfterUpdateCfgFn  func(v8end V8Endpoint, workdir string, version int64) error
)

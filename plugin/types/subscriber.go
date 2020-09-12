package types

import (
	"github.com/v8platform/designer/repository"
	v8 "github.com/v8platform/v8"
)

type Subscriber interface {
	Handlers() []interface{}
}

type UpdateCfgSubscriber struct {
	On     OnUpdateCfgFn
	Before BeforeUpdateCfgFn
	After  AfterUpdateCfgFn
}

type V8Endpoint interface {
	Infobase() *v8.Infobase
	Repository() *repository.Repository
	Extention() string
	Options() []interface{}
}

type (
	BeforeDumpConfigFn func(v8end V8Endpoint, workdir string, temp string, number int64) error
	OnDumpConfigFn     func(v8end V8Endpoint, workdir string, temp string, number int64, standartHandler *bool) error
	AfterDumpConfigFn  BeforeDumpConfigFn
)

type (
	BeforeUpdateCfgFn func(v8end V8Endpoint, workdir string, version int64) error
	OnUpdateCfgFn     func(v8end V8Endpoint, workdir string, version int64, standartHandler *bool) error
	AfterUpdateCfgFn  BeforeUpdateCfgFn
)

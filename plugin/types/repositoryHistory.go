package types

import (
	"github.com/khorevaa/r2gitsync/manager/types"
)

type GetRepositoryHistorySubscriber struct {
	On     OnGetRepositoryHistoryFn
	Before BeforeGetRepositoryHistoryFn
	After  AfterGetRepositoryHistoryFn
}

type ConfigureRepositoryVersionsSubscriber struct {
	On OnConfigureRepositoryVersionsFn
}

type OnConfigureRepositoryVersionsFn func(v8end V8Endpoint, versions *types.RepositoryVersionsList, NCurrent, NNext, NMax *int) error

type (
	BeforeGetRepositoryHistoryFn func(v8end V8Endpoint, dir string, NBegin int) error
	OnGetRepositoryHistoryFn     func(v8end V8Endpoint, dir string, NBegin int, stdHandler *bool) ([]types.RepositoryVersion, error)
	AfterGetRepositoryHistoryFn  func(v8end V8Endpoint, dir string, NBegin int, rv *types.RepositoryVersionsList) error
)

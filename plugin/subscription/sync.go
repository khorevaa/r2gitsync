package subscription

import (
	. "github.com/khorevaa/r2gitsync/plugin/types"
)

type SyncProcessHandler interface {
	SubscribeHandler
	Subscribe(sub SyncProcessSubscriber)

	Start(v8end V8Endpoint, dir string)
	Finish(v8end V8Endpoint, dir string, err *error)
}

var _ SyncProcessHandler = (*syncProcessHandler)(nil)

type syncProcessHandler struct {
	start  []StartSyncProcessFn
	finish []FinishSyncProcessFn
}

func (h *syncProcessHandler) Subscribe(sub SyncProcessSubscriber) {

	if sub.Start != nil {
		h.start = append(h.start, sub.Start)
	}

	if sub.Finish != nil {
		h.finish = append(h.finish, sub.Finish)
	}
}

func (h *syncProcessHandler) Start(v8end V8Endpoint, dir string) {

	for _, fn := range h.start {

		fn(v8end, dir)
	}
}

func (h *syncProcessHandler) Finish(v8end V8Endpoint, dir string, err *error) {

	for _, fn := range h.finish {

		fn(v8end, dir, err)
	}

}

type SyncVersionHandler interface {
	SubscribeHandler
	Subscribe(sub SyncVersionSubscriber)

	Start(v8end V8Endpoint, dir, temp string, number int64)
	Finish(v8end V8Endpoint, dir, temp string, number int64, err *error)
}

var _ SyncVersionHandler = (*syncversionHandler)(nil)

type syncversionHandler struct {
	start  []StartSyncVersionFn
	finish []FinishSyncVersionFn
}

func (h *syncversionHandler) Subscribe(sub SyncVersionSubscriber) {

	if sub.Start != nil {
		h.start = append(h.start, sub.Start)
	}

	if sub.Finish != nil {
		h.finish = append(h.finish, sub.Finish)
	}
}

func (h *syncversionHandler) Start(v8end V8Endpoint, dir, temp string, number int64) {

	for _, fn := range h.start {

		fn(v8end, dir, temp, number)
	}
}

func (h *syncversionHandler) Finish(v8end V8Endpoint, dir, temp string, number int64, err *error) {

	for _, fn := range h.finish {

		fn(v8end, dir, temp, number, err)
	}

}

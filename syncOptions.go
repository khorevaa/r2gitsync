package main

import v8 "github.com/v8platform/v8"

type SyncOption func(*SyncOptions)

type SyncOptions struct {
	tempDir     string
	v8Path      string
	v8version   string
	licTryCount int
	domainEmail string
	infobase    v8.Infobase
	plugins     *PluginsManager // TODO добавить менеджер плагинов

}

func WithInfobase(connString, user, password string) SyncOption {

	return func(o *SyncOptions) {

		if len(connString) == 0 {
			return
		}

		o.infobase = syncInfobase(connString, user, password)

	}

}

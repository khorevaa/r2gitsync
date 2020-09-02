package main

import v8 "github.com/v8platform/v8"

type SyncOption func(*SyncOptions)

type SyncOptions struct {
	tempDir         string
	v8Path          string
	v8version       string
	licTryCount     int
	domainEmail     string
	infobase        v8.Infobase
	infobaseCreated bool
	increment       bool
	limit           struct {
		MinVersion int
		MaxVersion int
		Count      int
	}
	plugins *PluginsManager // TODO добавить менеджер плагинов
}

type ib struct {
	User             string
	Password         string
	ConnectionString string
}

func WithInfobaseConfig(infobase ib) SyncOption {

	return WithInfobase(infobase.ConnectionString, infobase.User, infobase.Password)

}

func WithInfobase(connString, user, password string) SyncOption {

	return func(o *SyncOptions) {

		if len(connString) == 0 {
			o.infobase = v8.NewTempIB()
			return
		}

		o.infobase = syncInfobase(connString, user, password)
		o.infobaseCreated = true
	}

}

func WithPlugins(manager *PluginsManager) SyncOption {

	return func(o *SyncOptions) {

		if manager == nil {
			return
		}

		o.plugins = manager

	}

}

func WithV8version(v8version string) SyncOption {

	return func(o *SyncOptions) {

		if len(v8version) == 0 {
			return
		}

		o.v8version = v8version

	}

}

func WithDomainEmail(email string) SyncOption {

	return func(o *SyncOptions) {

		if len(email) == 0 {
			return
		}

		o.domainEmail = email

	}

}

func WithTempDir(path string) SyncOption {

	return func(o *SyncOptions) {

		if len(path) == 0 {
			return
		}

		o.tempDir = path

	}

}

func WithLicTryCount(count int) SyncOption {

	return func(o *SyncOptions) {

		if count == 0 {
			return
		}

		o.licTryCount = count

	}

}

func WithIncrementSync(increment bool) SyncOption {

	return func(o *SyncOptions) {

		o.increment = increment

	}

}

func WithV8Path(path string) SyncOption {

	return func(o *SyncOptions) {

		if len(path) == 0 {
			return
		}

		o.v8Path = path

	}

}

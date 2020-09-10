package manager

import (
	"github.com/khorevaa/r2gitsync/plugin/subscription"
	"github.com/v8platform/runner"
	v8 "github.com/v8platform/v8"
)

type PlSm struct {
	UpdateCfg             subscription.UpdateCfgHandler
	DumpConfigToFiles     subscription.DumpConfigToFilesHandler
	GetRepositoryHistoryH subscription.GetRepositoryHistoryHandler
}

type Option func(*Options)

type Options struct {
	tempDir          string
	v8Path           string
	v8version        string
	licTryCount      int
	domainEmail      string
	infobase         v8.Infobase
	infobaseCreated  bool
	disableIncrement bool
	limit            struct {
		MinVersion int
		MaxVersion int
		Count      int
	}
	plugins *PlSm // TODO добавить менеджер плагинов
}

type ib struct {
	User             string
	Password         string
	ConnectionString string
}

func (o *Options) Options() []runner.Option {

	return []runner.Option{
		v8.WithPath(o.v8Path),
		v8.WithVersion(o.v8version),
	}

}

func WithInfobaseConfig(infobase ib) Option {

	return WithInfobase(infobase.ConnectionString, infobase.User, infobase.Password)

}

func WithInfobase(connString, user, password string) Option {

	return func(o *Options) {

		if len(connString) == 0 {
			o.infobase = v8.NewTempIB()
			return
		}

		o.infobase = syncInfobase(connString, user, password)
		o.infobaseCreated = true
	}

}

func WithPlugins(manager *PlSm) Option {

	return func(o *Options) {

		if manager == nil {
			return
		}

		o.plugins = manager

	}

}

func WithDisableIncrement(disable bool) Option {

	return func(o *Options) {

		o.disableIncrement = disable

	}

}

func WithLimit(limit struct {
	MinVersion int
	MaxVersion int
	Count      int
}) Option {

	return func(o *Options) {

		o.limit = limit

	}

}

func WithV8version(v8version string) Option {

	return func(o *Options) {

		if len(v8version) == 0 {
			return
		}

		o.v8version = v8version

	}

}

func WithDomainEmail(email string) Option {

	return func(o *Options) {

		if len(email) == 0 {
			return
		}

		o.domainEmail = email

	}

}

func WithTempDir(path string) Option {

	return func(o *Options) {

		if len(path) == 0 {
			return
		}

		o.tempDir = path

	}

}

func WithLicTryCount(count int) Option {

	return func(o *Options) {

		if count == 0 {
			return
		}

		o.licTryCount = count

	}

}

func WithIncrementSync(increment bool) Option {

	return func(o *Options) {

		o.increment = increment

	}

}

func WithV8Path(path string) Option {

	return func(o *Options) {

		if len(path) == 0 {
			return
		}

		o.v8Path = path

	}

}

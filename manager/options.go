package manager

import (
	"github.com/khorevaa/r2gitsync/log"
	"github.com/khorevaa/r2gitsync/plugin/subscription"
	"github.com/v8platform/runner"
	v8 "github.com/v8platform/v8"
)

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
	plugins          *subscription.SubscribeManager // TODO добавить менеджер плагинов
	logger           log.Logger
	ForceInit        bool
}

type ib struct {
	User             string
	Password         string
	ConnectionString string
}

func (o *Options) DomainEmail() string {

	if len(o.domainEmail) == 0 {
		return "localhost"
	}
	return o.domainEmail
}

func (o *Options) Options() []runner.Option {

	return []runner.Option{
		v8.WithPath(o.v8Path),
		v8.WithVersion(o.v8version),
		v8.WithCommonValues([]string{"/DisableStartupDialogs", "/DisableStartupMessages"}),
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

func WithPlugins(manager *subscription.SubscribeManager) Option {

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

func WithForceInit(force bool) Option {

	return func(o *Options) {

		o.ForceInit = force

	}

}

func WithLogger(logger log.Logger) Option {

	return func(o *Options) {

		o.logger = logger

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

func WithV8Path(path string) Option {

	return func(o *Options) {

		if len(path) == 0 {
			return
		}

		o.v8Path = path

	}

}

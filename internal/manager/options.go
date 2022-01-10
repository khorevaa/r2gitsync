package manager

import (
	"github.com/khorevaa/r2gitsync/pkg/plugin/subscription"
	"github.com/v8platform/api"
)

type Option func(*Options)

type Options struct {
	TempDir          string
	DisableIncrement bool

	MinVersion       int
	MaxVersion       int
	LimitVersions    int
	InfobaseConnect  string
	InfobaseUser     string
	InfobasePassword string
	DomainEmail      string

	opts      []interface{}
	V8Path    string
	V8version string

	Plugins *subscription.SubscribeManager

	LicTryCount int
}

func (o *Options) Infobase() (*v8.Infobase, error) {

	return v8.ParseConnectionString(o.InfobaseConnect)

}

func (o *Options) Options() []interface{} {

	if len(o.opts) == 0 {
		o.opts = []interface{}{
			v8.WithPath(o.V8Path),
			v8.WithVersion(o.V8version),
			v8.WithCommonValues("/DisableStartupDialogs", "/DisableStartupMessages"),
			v8.WithCredentials(o.InfobaseUser, o.InfobasePassword),
		}
	}

	return o.opts

}

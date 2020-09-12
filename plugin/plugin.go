package plugin

import (
	"github.com/hashicorp/go-multierror"
	"github.com/khorevaa/r2gitsync/cmd/flags"
	"github.com/khorevaa/r2gitsync/context"
	"github.com/khorevaa/r2gitsync/plugin/subscription"
	. "github.com/khorevaa/r2gitsync/plugin/types"
	"github.com/pkg/errors"
)

// Plugin is the interface for plugins to micro. It differs from go-micro in that it's for
// the micro API, Web, Sidecar, CLI. It's a method of building middleware for the HTTP side.
type Symbol interface {

	// Global Flags
	Flags() []flags.Flag
	// Sub-commands
	Commands() []string

	// Name of the plugin
	String() string
	Desc() string
	Version() string
	ShortVersion() string
	Name() string
	Init() Plugin
}

type Plugin interface {
	Subscriber() Subscriber
	InitContext(ctx context.Context)
}

// Manager is the plugin manager which stores plugins and allows them to be retrieved.
// This is used by all the components of micro.
type Manager interface {
	Plugins() []Symbol
	EnabledPlugins() []Symbol
	Register(plugin Symbol) error
	IsRegistered(plugin Symbol) bool
	RegisterFlags(command string, cmd command, ctx context.Context)
	IsEnabled(name string) bool
	Enable(name string) error
	Disable(name string) error
}

// Plugins lists the global plugins
func Plugins() []Symbol {
	return defaultManager.Plugins()
}

// Register registers a global plugins
func Register(names ...Symbol) error {

	mErr := &multierror.Error{}

	for _, name := range names {
		err := defaultManager.Register(name)

		mErr = multierror.Append(mErr, errors.Wrapf(err, "plugin <%s>", name))
	}

	return mErr.ErrorOrNil()
}

// Enable a global plugins
func Enable(names ...string) error {

	mErr := &multierror.Error{}

	for _, name := range names {
		err := defaultManager.Enable(name)
		mErr = multierror.Append(mErr, err)
	}

	return mErr.ErrorOrNil()
}

// Disable a global plugins
func Disable(names ...string) error {

	mErr := &multierror.Error{}

	for _, name := range names {
		err := defaultManager.Disable(name)
		mErr = multierror.Append(mErr, err)
	}

	return mErr.ErrorOrNil()
}

// IsRegistered check plugin whether registered global.
// Notice plugin is not check whether is nil
func IsRegistered(plugin Symbol) bool {
	return defaultManager.IsRegistered(plugin)
}

func RegistryFlags(name string, cmd command, ctx context.Context) {

	defaultManager.RegisterFlags(name, cmd, ctx)

}

func SubscribeManager() *subscription.SubscribeManager {
	return defaultManager.sm
}

// NewManager creates a new plugin manager
func NewManager() Manager {
	return newManager()
}

func Subscription(handlers ...interface{}) Subscriber {

	return subscriber{
		handlers: handlers,
	}
}

func Subscribe() error {

	return defaultManager.Subscribe()
}

func SendContext(ctx context.Context) {
	defaultManager.SendContext(ctx)
}
func IsEnabled(name string) bool {
	return defaultManager.IsEnabled(name)
}

func LoadPlugins(dir string) error {

	loader := NewPluginsLoader(dir)

	err := loader.LoadPlugins(false)

	if err != nil {
		return err
	}

	return Register(loader.Plugins()...)

}

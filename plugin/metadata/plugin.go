package metadata

import (
	"crypto"
	"github.com/khorevaa/r2gitsync/cmd/flags"
	"github.com/khorevaa/r2gitsync/context"
	. "github.com/khorevaa/r2gitsync/plugin/types"
	"time"
)

// PluginSymbol is the interface for plugins
type PluginSymbol interface {

	// Global Flags
	Flags() []flags.Flag
	// Sub-Modules
	Modules() []string

	// Name of the plugin
	String() string
	Desc() string
	Version() string
	ShortVersion() string
	Name() string
	Init() Plugin
}

//Plugin основной интерфейс плагинов
type Plugin interface {
	Subscribe(ctx context.Context) Subscriber
}

type PluginMetadata struct {
	Name        string
	Version     string
	LongVersion string
	Desc        string
	Modules     []string

	Pkg PkgMetadata

	BuildAt   time.Time
	updateAt  time.Time
	createdAt time.Time
	Hash      crypto.Hash
}

func NewPluginMetadata(sym PluginSymbol, pkg PkgMetadata) PluginMetadata {
	return PluginMetadata{
		Name:        sym.Name(),
		Version:     sym.ShortVersion(),
		LongVersion: sym.Version(),
		Desc:        sym.Desc(),
		Modules:     sym.Modules(),
		Pkg:         pkg,
		BuildAt:     time.Now(), // TODO Сделать получение из данных пакета
		updateAt:    time.Time{},
		createdAt:   time.Time{},
		Hash:        pkg.hash,
	}
}

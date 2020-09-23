package plugin

type PluginStorage interface {
	RegistryPlugins(m *manager) error
	DisablePlugin(name ...string)
	EnablePlugin(name ...string)
	ListPlugins() []plugin

	Save() error
}

type StoragePlugin struct {
	Symbol
	File   string
	Enable bool
}

type FilePluginStorage struct {
}

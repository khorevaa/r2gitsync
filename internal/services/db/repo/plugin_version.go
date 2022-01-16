package repo

type PluginVersion struct {
	UuidModel
	PluginID  string `gorm:"TYPE:uuid REFERENCES plugins;uniqueIndex:idx_plugin_id_version"`
	Plugin    Plugin
	Version   string `gorm:"size:50;uniqueIndex:idx_plugin_id_version"`
	Changelog string
	Assets    []*Asset `gorm:"polymorphic:Owner;"`
}

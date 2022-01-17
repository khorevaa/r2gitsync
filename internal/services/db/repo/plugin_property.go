package repo

type PluginPropertyType string

type PluginProperty struct {
	UuidModel
	PluginUuid   string `gorm:"TYPE:uuid REFERENCES plugins;index;uniqueIndex:idx_plugin_uuid_version_uuid_name"`
	VersionUuid  string `gorm:"TYPE:uuid REFERENCES plugin_versions;index;uniqueIndex:idx_plugin_uuid_version_uuid_name"`
	Name         string `gorm:"uniqueIndex:idx_plugin_uuid_version_uuid_name"`
	Type         PluginPropertyType
	DefaultValue *string
	Required     string
}

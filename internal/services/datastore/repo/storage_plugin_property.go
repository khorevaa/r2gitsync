package repo

type StoragePluginProperty struct {
	UuidModel
	StoragePluginUuid  string `gorm:"TYPE:uuid REFERENCES storage_plugins;index;uniqueIndex:idx_plugin_property_uuid_name"`
	PluginPropertyUuid string `gorm:"TYPE:uuid REFERENCES plugin_properties;index;uniqueIndex:idx_plugin_property_uuid_name"`
	Name               string `gorm:"uniqueIndex:idx_plugin_property_uuid_name"`
	Value              *string
}

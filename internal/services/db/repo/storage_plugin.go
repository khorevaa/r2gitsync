package repo

type StoragePlugin struct {
	PluginID        uint `gorm:"TYPE:integer REFERENCES plugins;index"`  // Ссылка на плагин
	StorageID       uint `gorm:"TYPE:integer REFERENCES storages;index"` // Ссылка на репозиторий
	PluginVersionID uint `gorm:"TYPE:integer REFERENCES plugin_versions;index"`
}

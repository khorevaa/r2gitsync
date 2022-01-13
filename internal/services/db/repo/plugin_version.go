package repo

import "gorm.io/gorm"

type PluginVersion struct {
	gorm.Model
	PluginID  uint `gorm:"TYPE:integer REFERENCES plugins;uniqueIndex:idx_plugin_id_version"`
	Plugin    Plugin
	Version   string `gorm:"size:50;uniqueIndex:idx_plugin_id_version"`
	Changelog string
	Assets    []*Asset `gorm:"polymorphic:Owner;"`
}

package dto

type PluginVersion struct {
	Uuid       string
	PluginUuid string
	Version    string
	Changelog  string
	Assets     []*Asset
	Properties []*PluginProperty
}

type PluginVersions []*PluginVersion

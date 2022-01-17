package dto

type PluginPropertyType string

type PluginProperty struct {
	Uuid         string
	PluginUuid   string
	VersionUuid  string
	Name         string
	Type         PluginPropertyType
	DefaultValue *string
	Required     string
}

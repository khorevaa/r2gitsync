package dto

type StoragePlugin struct {
	Uuid              string // Идентификатор записи
	PluginUuid        string // Ссылка на плагин
	StorageUuid       string // Ссылка на репозиторий
	PluginVersionUuid string // Ссылка на версию плагина
}

type StoragePlugins []*StoragePlugin

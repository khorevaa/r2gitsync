package repo

import (
	"context"

	"github.com/khorevaa/r2gitsync/internal/dto"
	"gorm.io/gorm"
)

type StoragePlugin struct {
	UuidModel
	PluginUuid        string                   `gorm:"TYPE:uuid REFERENCES plugins;uniqueIndex:idx_plugin_uuid_storage_uuid"`  // Ссылка на плагин
	StorageUuid       string                   `gorm:"TYPE:uuid REFERENCES storages;uniqueIndex:idx_plugin_uuid_storage_uuid"` // Ссылка на репозиторий
	PluginVersionUuid string                   `gorm:"TYPE:uuid REFERENCES plugin_versions;index"`
	Properties        []*StoragePluginProperty `gorm:""`
}

type IStoragePluginRepository interface {
	Fetch(ctx context.Context) (dto.StoragePlugins, error)
	GetByStorageUuid(ctx context.Context, uuid string) (dto.StoragePlugins, error)
	GetByUuid(ctx context.Context, uuid string) (*dto.StoragePlugin, error)
	Store(ctx context.Context, dtm *dto.StoragePlugin) (*dto.StoragePlugin, error)
	Update(ctx context.Context, uuid string, dtm *dto.StoragePlugin) (*dto.StoragePlugin, error)
	Delete(ctx context.Context, uuid string) error
}

func NewStoragePluginRepository(db *gorm.DB) IStoragePluginRepository {
	return &StoragePluginRepository{db: db}
}

type StoragePluginRepository struct {
	db *gorm.DB
}

func (s StoragePluginRepository) GetByStorageUuid(ctx context.Context, uuid string) (dto.StoragePlugins, error) {
	// TODO implement me
	panic("implement me")
}

func (s StoragePluginRepository) Fetch(ctx context.Context) (dto.StoragePlugins, error) {
	// TODO implement me
	panic("implement me")
}

func (s StoragePluginRepository) GetByUuid(ctx context.Context, uuid string) (*dto.StoragePlugin, error) {
	// TODO implement me
	panic("implement me")
}

func (s StoragePluginRepository) Store(ctx context.Context, dtm *dto.StoragePlugin) (*dto.StoragePlugin, error) {
	// TODO implement me
	panic("implement me")
}

func (s StoragePluginRepository) Update(ctx context.Context, uuid string, dtm *dto.StoragePlugin) (*dto.StoragePlugin, error) {
	// TODO implement me
	panic("implement me")
}

func (s StoragePluginRepository) Delete(ctx context.Context, uuid string) error {
	// TODO implement me
	panic("implement me")
}

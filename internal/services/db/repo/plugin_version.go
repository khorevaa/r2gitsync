package repo

import (
	"context"

	"github.com/khorevaa/r2gitsync/internal/dto"
	"gorm.io/gorm"
)

type PluginVersion struct {
	UuidModel
	PluginUuid string `gorm:"TYPE:uuid REFERENCES plugins;uniqueIndex:idx_plugin_uuid_version"`
	Version    string `gorm:"size:50;uniqueIndex:idx_plugin_uuid_version"`
	Changelog  string
	Assets     []*Asset          `gorm:"polymorphic:Owner;"`
	Properties []*PluginProperty `gorm:""`
}
type IPluginVersionRepository interface {
	Fetch(ctx context.Context) (dto.PluginVersions, error)
	GetById(ctx context.Context, uuid string) (*dto.PluginVersion, error)
	Store(ctx context.Context, dtm *dto.PluginVersion) (*dto.PluginVersion, error)
	Update(ctx context.Context, uuid string, dtm *dto.PluginVersion) (*dto.PluginVersion, error)
	Delete(ctx context.Context, uuid string) error
}

func NewPluginVersionRepository(db *gorm.DB) IPluginVersionRepository {
	return &PluginVersionRepository{db: db}
}

type PluginVersionRepository struct {
	db *gorm.DB
}

func (p PluginVersionRepository) Fetch(ctx context.Context) (dto.PluginVersions, error) {
	// TODO implement me
	panic("implement me")
}

func (p PluginVersionRepository) GetById(ctx context.Context, uuid string) (*dto.PluginVersion, error) {
	// TODO implement me
	panic("implement me")
}

func (p PluginVersionRepository) Store(ctx context.Context, dtm *dto.PluginVersion) (*dto.PluginVersion, error) {
	// TODO implement me
	panic("implement me")
}

func (p PluginVersionRepository) Update(ctx context.Context, uuid string, dtm *dto.PluginVersion) (*dto.PluginVersion, error) {
	// TODO implement me
	panic("implement me")
}

func (p PluginVersionRepository) Delete(ctx context.Context, uuid string) error {
	// TODO implement me
	panic("implement me")
}

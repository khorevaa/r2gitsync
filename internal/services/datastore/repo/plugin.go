package repo

import (
	"context"

	"github.com/khorevaa/r2gitsync/internal/dto"
	"gorm.io/gorm"
)

type Plugin struct {
	UuidModel
	Name        string `gorm:"size:50;uniqueIndex"`
	Description string
}

type IPluginRepository interface {
	Fetch(ctx context.Context) (dto.Plugins, error)
	GetByUuid(ctx context.Context, uuid string) (*dto.Plugin, error)
	Store(ctx context.Context, dtm *dto.Plugin) (*dto.Plugin, error)
	Update(ctx context.Context, uuid string, dtm *dto.Plugin) (*dto.Plugin, error)
	Delete(ctx context.Context, uuid string) error
}

func NewPluginRepository(db *gorm.DB) IPluginRepository {
	return &PluginRepository{db: db}
}

type PluginRepository struct {
	db *gorm.DB
}

func (p PluginRepository) Fetch(ctx context.Context) (dto.Plugins, error) {
	// TODO implement me
	panic("implement me")
}

func (p PluginRepository) GetByUuid(ctx context.Context, uuid string) (*dto.Plugin, error) {
	// TODO implement me
	panic("implement me")
}

func (p PluginRepository) Store(ctx context.Context, dtm *dto.Plugin) (*dto.Plugin, error) {
	// TODO implement me
	panic("implement me")
}

func (p PluginRepository) Update(ctx context.Context, uuid string, dtm *dto.Plugin) (*dto.Plugin, error) {
	// TODO implement me
	panic("implement me")
}

func (p PluginRepository) Delete(ctx context.Context, uuid string) error {
	// TODO implement me
	panic("implement me")
}

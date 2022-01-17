package repo

import (
	"context"

	"github.com/khorevaa/r2gitsync/internal/dto"
	"gorm.io/gorm"
)

type Asset struct {
	UuidModel

	OwnerID   string
	OwnerType string

	Filename string
	Size     uint
}

type IAssetRepository interface {
	Fetch(ctx context.Context) (dto.Assets, error)
	GetByUuid(ctx context.Context, uuid string) (*dto.Asset, error)
	Store(ctx context.Context, dtm *dto.Asset) (*dto.Asset, error)
	Update(ctx context.Context, uuid string, dtm *dto.Asset) (*dto.Asset, error)
	Delete(ctx context.Context, uuid string) error
}

func NewAssetRepository(db *gorm.DB) IAssetRepository {
	return &AssetRepository{db: db}
}

type AssetRepository struct {
	db *gorm.DB
}

func (a AssetRepository) Fetch(ctx context.Context) (dto.Assets, error) {
	// TODO implement me
	panic("implement me")
}

func (a AssetRepository) GetByUuid(ctx context.Context, uuid string) (*dto.Asset, error) {
	// TODO implement me
	panic("implement me")
}

func (a AssetRepository) Store(ctx context.Context, dtm *dto.Asset) (*dto.Asset, error) {
	// TODO implement me
	panic("implement me")
}

func (a AssetRepository) Update(ctx context.Context, uuid string, dtm *dto.Asset) (*dto.Asset, error) {
	// TODO implement me
	panic("implement me")
}

func (a AssetRepository) Delete(ctx context.Context, uuid string) error {
	// TODO implement me
	panic("implement me")
}

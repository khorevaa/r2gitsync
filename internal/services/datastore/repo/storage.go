package repo

import (
	"context"

	"github.com/khorevaa/r2gitsync/internal/dto"
	"gorm.io/gorm"
)

type Storage struct {
	UuidModel
	ConnectionString string
	Type             dto.StorageType
	Develop          bool
	Extension        *string
	ParentUuid       *uint
	Parent           *Storage

	ProjectCode string `gorm:"TYPE:uuid REFERENCES projects;index;references:Code"`
}

type IStorageRepository interface {
	Fetch(ctx context.Context) (dto.Storages, error)
	GetByUuid(ctx context.Context, uuid string) (*dto.Storage, error)
	Store(ctx context.Context, dtm *dto.Storage) (*dto.Storage, error)
	Update(ctx context.Context, uuid string, dtm *dto.Storage) (*dto.Storage, error)
	Delete(ctx context.Context, uuid string) error
}

func NewStorageRepository(db *gorm.DB) IStorageRepository {
	return &StorageRepository{db: db}
}

type StorageRepository struct {
	db *gorm.DB
}

func (s StorageRepository) Fetch(ctx context.Context) (dto.Storages, error) {
	// TODO implement me
	panic("implement me")
}

func (s StorageRepository) GetByUuid(ctx context.Context, uuid string) (*dto.Storage, error) {
	// TODO implement me
	panic("implement me")
}

func (s StorageRepository) Store(ctx context.Context, dtm *dto.Storage) (*dto.Storage, error) {
	// TODO implement me
	panic("implement me")
}

func (s StorageRepository) Update(ctx context.Context, uuid string, dtm *dto.Storage) (*dto.Storage, error) {
	// TODO implement me
	panic("implement me")
}

func (s StorageRepository) Delete(ctx context.Context, uuid string) error {
	// TODO implement me
	panic("implement me")
}

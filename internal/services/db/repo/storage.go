package repo

import (
	"context"

	"github.com/khorevaa/r2gitsync/internal/dto"
	"gorm.io/gorm"
)

type IStorageRepository interface {
	Fetch(ctx context.Context) (dto.Storages, error)
	GetById(ctx context.Context, id uint) (*dto.Storage, error)
	Store(ctx context.Context, dtm *dto.Storage) (*dto.Storage, error)
	Update(ctx context.Context, id uint, dtm *dto.Storage) (*dto.Storage, error)
	Delete(ctx context.Context, id uint) error
}

func NewStorageRepository(db *gorm.DB) IStorageRepository {
	return &StorageRepository{db: db}
}

type Storage struct {
	gorm.Model
	ConnectionString string
	Type             dto.StorageType
	Develop          bool
	Extension        *uint
	ParentID         *uint
	Parent           *Storage

	ProjectCode string  // `gorm:"TYPE:string REFERENCES projects;index"`
	Project     Project `gorm:"foreignKey:ProjectCode; references:Code; constraint:OnUpdate:CASCADE,OnDelete:DELETE;"`
}

type StorageRepository struct {
	db *gorm.DB
}

func (s StorageRepository) Fetch(ctx context.Context) (dto.Storages, error) {
	// TODO implement me
	panic("implement me")
}

func (s StorageRepository) GetById(ctx context.Context, id uint) (*dto.Storage, error) {
	// TODO implement me
	panic("implement me")
}

func (s StorageRepository) Store(ctx context.Context, dtm *dto.Storage) (*dto.Storage, error) {
	// TODO implement me
	panic("implement me")
}

func (s StorageRepository) Update(ctx context.Context, id uint, dtm *dto.Storage) (*dto.Storage, error) {
	// TODO implement me
	panic("implement me")
}

func (s StorageRepository) Delete(ctx context.Context, id uint) error {
	// TODO implement me
	panic("implement me")
}

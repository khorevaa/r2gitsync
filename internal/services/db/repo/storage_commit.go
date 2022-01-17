package repo

import (
	"context"
	"time"

	"github.com/khorevaa/r2gitsync/internal/dto"
	"gorm.io/gorm"
)

type StorageCommit struct {
	UuidModel
	StorageUuid          string `gorm:"TYPE:uuid REFERENCES storages;index;uniqueIndex:idx_storage_uuid_number"`
	Number               uint   `gorm:"uniqueIndex:idx_storage_uuid_number"`
	ConfigurationVersion string
	Author               string
	Description          string
	CommitAt             time.Time `gorm:"index"`
	Tag                  string
	TagDesc              string
}

type IStorageCommitRepository interface {
	Fetch(ctx context.Context) (dto.StorageCommits, error)
	GetByUuid(ctx context.Context, uuid string) (*dto.StorageCommit, error)
	Store(ctx context.Context, dtm *dto.StorageCommit) (*dto.StorageCommit, error)
	Update(ctx context.Context, uuid string, dtm *dto.StorageCommit) (*dto.StorageCommit, error)
	Delete(ctx context.Context, uuid string) error
}

func NewStorageCommitRepository(db *gorm.DB) IStorageCommitRepository {
	return &StorageCommitRepository{db: db}
}

type StorageCommitRepository struct {
	db *gorm.DB
}

func (s StorageCommitRepository) Fetch(ctx context.Context) (dto.StorageCommits, error) {
	// TODO implement me
	panic("implement me")
}

func (s StorageCommitRepository) GetByUuid(ctx context.Context, uuid string) (*dto.StorageCommit, error) {
	// TODO implement me
	panic("implement me")
}

func (s StorageCommitRepository) Store(ctx context.Context, dtm *dto.StorageCommit) (*dto.StorageCommit, error) {
	// TODO implement me
	panic("implement me")
}

func (s StorageCommitRepository) Update(ctx context.Context, uuid string, dtm *dto.StorageCommit) (*dto.StorageCommit, error) {
	// TODO implement me
	panic("implement me")
}

func (s StorageCommitRepository) Delete(ctx context.Context, uuid string) error {
	// TODO implement me
	panic("implement me")
}

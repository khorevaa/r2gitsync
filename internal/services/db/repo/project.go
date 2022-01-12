package repo

import (
	"context"

	"github.com/khorevaa/r2gitsync/internal/dto"
	"gorm.io/gorm"
)

type IProjectRepository interface {
	Fetch(ctx context.Context) (dto.Project, error)
	GetById(ctx context.Context, id uint) (*dto.Project, error)
	Store(ctx context.Context, dtm *dto.Project) (*dto.Project, error)
	Update(ctx context.Context, id uint, dtm *dto.Project) (*dto.Project, error)
	Delete(ctx context.Context, id uint) error
}

func NewProjectRepository(db *gorm.DB) IProjectRepository {
	return &ProjectRepository{db: db}
}

type Project struct {
	gorm.Model
	Code        string `gorm:"size:10;uniqueIndex"`
	Name        string
	Description string
}

type ProjectRepository struct {
	db *gorm.DB
}

func (p ProjectRepository) Fetch(ctx context.Context) (dto.Project, error) {
	// TODO implement me
	panic("implement me")
}

func (p ProjectRepository) GetById(ctx context.Context, id uint) (*dto.Project, error) {
	// TODO implement me
	panic("implement me")
}

func (p ProjectRepository) Store(ctx context.Context, dtm *dto.Project) (*dto.Project, error) {
	// TODO implement me
	panic("implement me")
}

func (p ProjectRepository) Update(ctx context.Context, id uint, dtm *dto.Project) (*dto.Project, error) {
	// TODO implement me
	panic("implement me")
}

func (p ProjectRepository) Delete(ctx context.Context, id uint) error {
	// TODO implement me
	panic("implement me")
}

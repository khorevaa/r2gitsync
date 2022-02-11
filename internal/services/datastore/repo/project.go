package repo

import (
	"context"

	"github.com/khorevaa/r2gitsync/internal/dto"
	"github.com/khorevaa/r2gitsync/internal/services/db"
)

type IProjectRepository interface {
	Fetch(ctx context.Context) (dto.Projects, error)
	GetById(ctx context.Context, uuid string) (*dto.Project, error)
	Store(ctx context.Context, dtm *dto.Project) (*dto.Project, error)
	Update(ctx context.Context, uuid string, dtm *dto.Project) (*dto.Project, error)
	Delete(ctx context.Context, uuid string) error
}

func NewProjectRepository(db *db.Client) IProjectRepository {
	return &ProjectRepository{db: db}
}

type ProjectRepository struct {
	db *db.Client
}

func (p ProjectRepository) Fetch(ctx context.Context) (dto.Projects, error) {

	projects, err := p.db.Project.Query().
		WithMasterStorage().
		WithDevelopStorage().
		WithStorages().
		All(ctx)
	if err != nil {
		return dto.Projects{}, err
	}

	return (dto.Projects{}).FromEnt(projects), nil
}

func (p ProjectRepository) GetById(ctx context.Context, uuid string) (*dto.Project, error) {
	// TODO implement me
	panic("implement me")
}

func (p ProjectRepository) Store(ctx context.Context, dtm *dto.Project) (*dto.Project, error) {
	// TODO implement me
	panic("implement me")
}

func (p ProjectRepository) Update(ctx context.Context, uuid string, dtm *dto.Project) (*dto.Project, error) {
	// TODO implement me
	panic("implement me")
}

func (p ProjectRepository) Delete(ctx context.Context, uuid string) error {
	// TODO implement me
	panic("implement me")
}

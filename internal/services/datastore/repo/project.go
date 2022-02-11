package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/khorevaa/r2gitsync/internal/dto"
	"github.com/khorevaa/r2gitsync/internal/services/db"
	"github.com/khorevaa/r2gitsync/internal/services/db/project"
)

type IProjectRepository interface {
	Fetch(ctx context.Context) (dto.Projects, error)
	GetById(ctx context.Context, uuid uuid.UUID) (*dto.Project, error)
	Store(ctx context.Context, dtm *dto.Project) (*dto.Project, error)
	Update(ctx context.Context, uuid uuid.UUID, dtm *dto.Project) (*dto.Project, error)
	Delete(ctx context.Context, uuid uuid.UUID) error
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

func (p ProjectRepository) GetById(ctx context.Context, uuid uuid.UUID) (*dto.Project, error) {
	// TODO implement me
	panic("implement me")
}

func (p ProjectRepository) Store(ctx context.Context, dtm *dto.Project) (*dto.Project, error) {
	dbm, err := p.db.Project.Create().
		// SetID(uuid.New()).
		SetCode(dtm.Code).
		SetDescription(dtm.Description).
		SetName(dtm.Name).
		SetType(project.Type(dtm.Type)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return (&dto.Project{}).FromEnt(dbm), nil
}

func (p ProjectRepository) Update(ctx context.Context, uuid uuid.UUID, dtm *dto.Project) (*dto.Project, error) {
	dbm, err := p.db.Project.UpdateOneID(uuid).
		SetCode(dtm.Code).
		SetDescription(dtm.Description).
		SetName(dtm.Name).
		SetType(project.Type(dtm.Type)).
		SetNillableMasterStorageID(dtm.MasterStorageID).
		SetNillableDevelopStorageID(dtm.DevelopStorageID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return (&dto.Project{}).FromEnt(dbm), nil
}

func (p ProjectRepository) Delete(ctx context.Context, uuid uuid.UUID) error {
	// TODO implement me
	panic("implement me")
}

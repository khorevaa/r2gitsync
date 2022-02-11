package logic

import (
	"context"

	"github.com/google/uuid"
	"github.com/khorevaa/r2gitsync/internal/di"
	"github.com/khorevaa/r2gitsync/internal/dto"
)

type IProjectsLogic interface {
	GetProjects(ctx context.Context) (dto.Projects, error)
	CreateProject(ctx context.Context, project *dto.Project) (*dto.Project, error)
	UpdateProject(ctx context.Context, uuid uuid.UUID, project *dto.Project) (*dto.Project, error)
}

func NewProjectsLogic(di di.IAppDeps) IProjectsLogic {
	return &ProjectsLogic{di: di}
}

type ProjectsLogic struct {
	di di.IAppDeps
}

func (p *ProjectsLogic) UpdateProject(ctx context.Context, uuid uuid.UUID, project *dto.Project) (*dto.Project, error) {
	return p.di.DB().Projects.Update(ctx, uuid, project)
}

func (p *ProjectsLogic) CreateProject(ctx context.Context, project *dto.Project) (*dto.Project, error) {
	return p.di.DB().Projects.Store(ctx, project)
}

func (p *ProjectsLogic) GetProjects(ctx context.Context) (dto.Projects, error) {

	projects, err := p.di.DB().Projects.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

package logic

import (
	"context"

	"github.com/khorevaa/r2gitsync/internal/di"
	"github.com/khorevaa/r2gitsync/internal/dto"
)

type IProjectsLogic interface {
	GetProjects(ctx context.Context) (dto.Projects, error)
}

func NewProjectsLogic(di di.IAppDeps) IProjectsLogic {
	return &ProjectsLogic{di: di}
}

type ProjectsLogic struct {
	di di.IAppDeps
}

func (p ProjectsLogic) GetProjects(ctx context.Context) (dto.Projects, error) {
	// TODO implement me
	panic("implement me")
}

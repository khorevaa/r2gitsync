package db

import (
	"fmt"

	"github.com/elastic/go-ucfg"
	"github.com/khorevaa/r2gitsync/internal/services/db/repo"
	"gorm.io/gorm"
)

type Repository struct {
	Project repo.IProjectRepository
}

func New(cfg *ucfg.Config) (*Repository, error) {

	var orm *gorm.DB

	orm, err := connectDb(cfg)
	if err != nil {
		return nil, err
	}

	return &Repository{
		Project: repo.NewProjectRepository(orm),
	}, nil

}

func connectDb(_ *ucfg.Config) (*gorm.DB, error) {
	return nil, fmt.Errorf("error connect to db")
}

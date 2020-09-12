package manager

import (
	"github.com/khorevaa/r2gitsync/manager/flow"
	"github.com/v8platform/designer/repository"
	v8 "github.com/v8platform/v8"
)

type RepositoryAuthor interface {
	flow.RepositoryAuthor
}

type RepositoryVersion interface {
	flow.RepositoryVersion
}

type V8Endpoint interface {
	Infobase() *v8.Infobase
	Repository() *repository.Repository
	Extention() string
	Options() []interface{}
}

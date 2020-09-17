package types

import (
	"github.com/v8platform/designer/repository"
	v8 "github.com/v8platform/v8"
	"time"
)

type V8Endpoint interface {
	Infobase() *v8.Infobase
	Repository() *repository.Repository
	Extention() string
	Options() []interface{}
}

type RepositoryAuthor interface {
	Name() string
	Email() string
	Desc() string
}

type RepositoryVersion interface {
	Version() string
	Author() string
	Date() time.Time
	Comment() string
	Number() int64
}

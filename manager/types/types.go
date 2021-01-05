package types

import (
	v8 "github.com/v8platform/api"
	"github.com/v8platform/designer"
	"time"
)

type V8Endpoint interface {
	Infobase() *v8.Infobase
	Repository() *designer.Repository
	Extention() string
	Options() []interface{}
}

type RepositoryAuthor interface {
	Name() string
	Email() string
	Desc() string
}

type RepositoryAuthorsList map[string]RepositoryAuthor

type RepositoryVersionsList []RepositoryVersion

type RepositoryVersion interface {
	Version() string
	Author() string
	Date() time.Time
	Comment() string
	Number() int
}

package manager

import (
	"fmt"
	"github.com/khorevaa/r2gitsync/manager/flow"
	"github.com/v8platform/designer/repository"
	v8 "github.com/v8platform/v8"
	"time"
)

type RepositoryAuthor interface {
	Name() string
	Email() string
	Desc() string
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

func NewAuthor(name, email string) repositoryAuthor {

	return repositoryAuthor{
		name:  name,
		email: email,
	}

}

type repositoryVersion struct {
	version string
	author  string
	date    time.Time
	comment string
	number  int64
}

func (r repositoryVersion) Version() string {
	return r.version
}

func (r repositoryVersion) Author() string {
	return r.author
}

func (r repositoryVersion) Date() time.Time {
	return r.date
}

func (r repositoryVersion) Comment() string {
	return r.comment
}

func (r repositoryVersion) Number() int64 {
	return r.number
}

type repositoryAuthor struct {
	name  string
	email string
}

func (a repositoryAuthor) Name() string {
	return a.name
}

func (a repositoryAuthor) Email() string {
	return a.email
}

func (a repositoryAuthor) Desc() string {

	return fmt.Sprintf("%s <%s> ", a.name, a.email)
}

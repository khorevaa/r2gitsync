package manager

import (
	"github.com/v8platform/designer/repository"
	v8 "github.com/v8platform/v8"
)

type V8Endpoint interface {
	Infobase() *v8.Infobase
	Repository() *repository.Repository
	Extention() string
	Options() []interface{}
}

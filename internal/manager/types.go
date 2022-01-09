package manager

import (
	v8 "github.com/v8platform/api"
	"github.com/v8platform/designer"
)

type V8Endpoint interface {
	Infobase() *v8.Infobase
	Repository() *designer.Repository
	Extention() string
	Options() []interface{}
}

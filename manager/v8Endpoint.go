package manager

import (
	v8 "github.com/v8platform/api"
	"github.com/v8platform/designer"
)

type v8Endpoint struct {
	infobase   *v8.Infobase
	repository *designer.Repository
	options    []interface{}

	extention string
}

func (end *v8Endpoint) Infobase() *v8.Infobase {

	return end.infobase

}

func (end *v8Endpoint) Repository() *designer.Repository {

	return end.repository
}

func (end *v8Endpoint) Options() []interface{} {

	return end.options

}

func (end *v8Endpoint) Extention() string {
	return end.extention
}

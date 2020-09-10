package manager

import (
	"github.com/v8platform/designer/repository"
	"github.com/v8platform/runner"
	v8 "github.com/v8platform/v8"
)

type v8Endpoint struct {
	infobase   *v8.Infobase
	repository *repository.Repository
	options    []runner.Option

	extention string
}

func (end *v8Endpoint) Infobase() *v8.Infobase {

	return end.infobase

}

func (end *v8Endpoint) Repository() *repository.Repository {

	return end.repository

}

func (end *v8Endpoint) Options() []interface{} {

	var opts []interface{}

	for o := range end.options {

		opts = append(opts, o)

	}

	return opts

}

func (end *v8Endpoint) Extention() string {
	return end.extention
}

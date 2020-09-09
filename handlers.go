package main

import (
	"github.com/v8platform/designer/repository"
)

func UpdateCfgHandler(v8end V8Endpoint, workDir string, version int64) error {

	RepositoryUpdateCfgOptions := repository.RepositoryUpdateCfgOptions{
		Version:   version,
		Force:     true,
		Extension: v8end.Extention(),
	}.
		WithRepository(*v8end.Repository())

	err := Run(*v8end.Infobase(), RepositoryUpdateCfgOptions, v8end.Options()...)

	if err != nil {
		return err
	}
	return nil
}

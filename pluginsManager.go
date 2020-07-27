package main

import (
	"github.com/v8platform/designer/repository"
	v8 "github.com/v8platform/v8"
)

type PluginsManager struct {
}

func (m *PluginsManager) BeforeUpdateCfgHandler(workdir string, infobase v8.Infobase, repository repository.Repository,
	version int64, extention string) error {

	return nil

}

func (m *PluginsManager) WithUpdateCfgHandler(infobase v8.Infobase, repository repository.Repository,
	version int64, extention string, standartHandler *bool) error {

	return nil

}

func (m *PluginsManager) AfterUpdateCfgHandler(infobase v8.Infobase, repository repository.Repository,
	version int64, extention string) error {

	return nil

}

func (m *PluginsManager) WithDumpCfgToFilesHandler(infobase v8.Infobase, repository repository.Repository, extention string, dir *string, b *bool) error {

	return nil
}

func (m *PluginsManager) AfterDumpCfgToFilesHandler(infobase v8.Infobase, r repository.Repository, extention string, dir *string) error {
	return nil
}

func (m *PluginsManager) BeforeDumpConfigToFiles(dir string, dir2 string, infobase v8.Infobase, r repository.Repository, number int64, extention string) error {
	return nil
}

func (m *PluginsManager) WithClearWorkdirHandler(dir string, b *bool) error {
	return nil
}

func (m *PluginsManager) AfterClearWorkdirHandler(dir string) error {
	return nil
}

func (m *PluginsManager) WithMoveToWorkDirHandler(dir string, dir2 string, b *bool) error {
	return nil
}

func (m *PluginsManager) AfterMoveToWorkDirHandler(dir string, dir2 string) error {
	return nil
}

func (m *PluginsManager) BeforeClearWorkDir(dir string, number int64) error {

	return nil

}

func (m *PluginsManager) BeforeMoveToWorkDir(dir string, dir2 string, number int64) error {

	return nil

}

func (m *PluginsManager) FinishSyncVersionHandler(dir string, dir2 string, r repository.Repository, number int64, extention string, err error) {

}

func (m *PluginsManager) BeforeStartSyncVersionHandler(dir string, dir2 string, r repository.Repository, number int64, extention string) error {

	return nil

}

func (m *PluginsManager) AfterGetRepositoryHistoryHandler(dir string, r repository.Repository, versions *[]repositoryVersion) error {

	return nil

}

func (m *PluginsManager) WithGetRepositoryHistoryHandler(dir string, r repository.Repository, i *[]repositoryVersion, opts *SyncOptions, b *bool) error {

	return nil
}

func (m *PluginsManager) WithGetRepositoryAuthorsHandler(dir string, r repository.Repository, authors *AuthorsList, opts *SyncOptions, b *bool) error {

	return nil
}

func (m *PluginsManager) AfterGetGetRepositoryAuthorsHandler(dir string, r repository.Repository, authors *AuthorsList) error {
	return nil
}

func (m *PluginsManager) BeforeStartSyncVersions(i *[]repositoryVersion, currentVersion int64, nextVersion int64, maxVersion *int64) {

}

func (m *PluginsManager) BeforeSyncVersion(number int64, a string, comment string, b *bool, opts *SyncOptions) error {

	return nil

}

func (m *PluginsManager) AfterSyncVersion(number int64, a string, comment string, opts *SyncOptions) error {
	return nil
}

func (m *PluginsManager) BeforeStartSyncProcess(r repository.Repository, dir string) {

}

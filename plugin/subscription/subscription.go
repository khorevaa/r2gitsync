package subscription

import (
	"github.com/khorevaa/r2gitsync/context"
	. "github.com/khorevaa/r2gitsync/plugin/types"
	"sync"
)

type SubscribeManager struct {
	mu                    sync.Mutex
	UpdateCfg             UpdateCfgHandler
	DumpConfigToFiles     DumpConfigToFilesHandler
	GetRepositoryHistoryH GetRepositoryHistoryHandler

	subscribers []Plugin
	count       int
}

func (sm *SubscribeManager) Subscribe(p Plugin, ctx context.Context) error {

	count := sm.count

	sub := p.Subscribe(ctx)

	sm.subscribe(sub)

	if count == sm.count {
		return nil
	}

	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.subscribers = append(sm.subscribers, p)

	return nil

}

func (sm *SubscribeManager) subscribe(sub Subscriber) {

	handlers := sub.Handlers()

	for _, handler := range handlers {

		switch h := handler.(type) {
		case UpdateCfgSubscriber:
			sm.UpdateCfg.Subscribe(h)
		}
	}

	//sm.UpdateCfg.Subscribe(sub.UpdateCfg)

}

//
//
//
//
//func (m *PluginsManager) BeforeUpdateCfgHandler(workdir string, infobase v8.Infobase, repository repository.Repository,
//	version int64, extention string) error {
//
//	return m.sm.BeforeUpdateCfgHandler(workdir, infobase, repository, version, extention)
//
//}
//
//func (m *PluginsManager) WithUpdateCfgHandler(infobase v8.Infobase, repository repository.Repository,
//	version int64, extention string, standartHandler *bool) error {
//
//	return nil
//
//}
//
//func (m *PluginsManager) AfterUpdateCfgHandler(infobase v8.Infobase, repository repository.Repository,
//	version int64, extention string) error {
//
//	return nil
//
//}
//
//func (m *PluginsManager) WithDumpCfgToFilesHandler(infobase v8.Infobase, repository repository.Repository, extention string, dir *string, b *bool) error {
//
//	return nil
//}
//
//func (m *PluginsManager) AfterDumpCfgToFilesHandler(infobase v8.Infobase, r repository.Repository, extention string, dir *string) error {
//	return nil
//}
//
//func (m *PluginsManager) BeforeDumpConfigToFiles(dir string, dir2 string, infobase v8.Infobase, r repository.Repository, number int64, extention string) error {
//	return nil
//}
//
//func (m *PluginsManager) WithClearWorkdirHandler(dir string, b *bool) error {
//	return nil
//}
//
//func (m *PluginsManager) AfterClearWorkdirHandler(dir string) error {
//	return nil
//}
//
//func (m *PluginsManager) WithMoveToWorkDirHandler(dir string, dir2 string, b *bool) error {
//	return nil
//}
//
//func (m *PluginsManager) AfterMoveToWorkDirHandler(dir string, dir2 string) error {
//	return nil
//}
//
//func (m *PluginsManager) BeforeClearWorkDir(dir string, number int64) error {
//
//	return nil
//
//}
//
//func (m *PluginsManager) BeforeMoveToWorkDir(dir string, dir2 string, number int64) error {
//
//	return nil
//
//}
//
//func (m *PluginsManager) FinishSyncVersionHandler(dir string, dir2 string, r repository.Repository, number int64, extention string, err error) {
//
//}
//
//func (m *PluginsManager) BeforeStartSyncVersionHandler(dir string, dir2 string, r repository.Repository, number int64, extention string) error {
//
//	return nil
//
//}
//
//func (m *PluginsManager) AfterGetRepositoryHistoryHandler(dir string, r repository.Repository, versions *[]repositoryVersion) error {
//
//	return nil
//
//}
//
//func (m *PluginsManager) WithGetRepositoryHistoryHandler(dir string, r repository.Repository, i *[]repositoryVersion, opts *SyncOptions, b *bool) error {
//
//	return nil
//}
//
//func (m *PluginsManager) WithGetRepositoryAuthorsHandler(dir string, r repository.Repository, authors *AuthorsList, opts *SyncOptions, b *bool) error {
//
//	return nil
//}
//
//func (m *PluginsManager) AfterGetGetRepositoryAuthorsHandler(dir string, r repository.Repository, authors *AuthorsList) error {
//	return nil
//}
//
//func (m *PluginsManager) BeforeStartSyncVersions(i *[]repositoryVersion, currentVersion int64, nextVersion int64, maxVersion *int64) {
//
//}
//
//func (m *PluginsManager) BeforeSyncVersion(number int64, a string, comment string, b *bool, opts *SyncOptions) error {
//
//	return nil
//
//}
//
//func (m *PluginsManager) AfterSyncVersion(number int64, a string, comment string, opts *SyncOptions) error {
//	return nil
//}
//
//func (m *PluginsManager) StartSyncProcess(r repository.Repository, dir string) {
//
//}

package main

import (
	"github.com/hashicorp/go-multierror"
	"github.com/v8platform/designer"
	"github.com/v8platform/designer/repository"
	"github.com/v8platform/v8"
	"os"
	"strconv"
	"time"
)

type SyncInfobase struct {
	User             string
	Password         string
	ConnectionString string

	v8.Infobase
}

func syncInfobase(connString, user, password string) v8.Infobase {

	if len(connString) == 0 {
		return v8.NewTempIB()
	}
	// TODO Сделать получение базы для выполнения синхронизации
	return v8.NewTempIB()

}

type repositoryVersion struct {
	version int64
	user    string
	date    time.Time
	comment string
}

type SyncAction struct {
	do   func()
	next *SyncAction
}

type syncVersion struct {
	version repositoryVersion

	repository *SyncRepository
	oprions    *SyncOptions

	onError func(err error)
	action  *SyncAction

	err error

	dumpDir string
}

func (s *syncVersion) Do() error {

	var e error

	onError := func(err error) {
		e = err // TODO Написано ГУАНО временно
		s.err = err

		err2 := writeVersionFile(s.repository.workdir, strconv.FormatInt(s.repository.currentVersion, 10))
		if err2 != nil {
			e = multierror.Append(e, err2)
		}
	}

	s.onError = onError

	s.RepositoryUpdateCfg().
		DumpConfigToFiles().
		ClearWorkdir().
		MoveToWorkdir().
		WriteVersionFile()

	return e

}

func (s *syncVersion) RepositoryUpdateCfg() *syncVersion {

	if s.err != nil {
		return s
	}

	err := v8.Run(s.repository.infobase, repository.RepositoryUpdateCfgOptions{
		Repository: s.repository.Repository,
		Version:    s.version.version,
		Force:      true,
		Extension:  s.oprions.extention,
	})

	if err != nil {
		s.onError(err)
	}
	return s
}

func (s *syncVersion) DumpConfigToFiles() *syncVersion {

	if s.err != nil {
		return s
	}

	err := v8.Run(s.repository.infobase, designer.DumpConfigToFilesOptions{
		Dir:       s.dumpDir,
		Force:     true,
		Extension: s.oprions.extention,
	})

	if err != nil {
		s.onError(err)
	}
	return s
}

func (s *syncVersion) ClearWorkdir() *syncVersion {

	if s.err != nil {
		return s
	}

	err := os.RemoveAll(s.repository.workdir) // TODO Сделать копирование файлов или избранную очистку

	if err != nil {
		s.onError(err)
	}
	return s
}

func (s *syncVersion) MoveToWorkdir() *syncVersion {

	if s.err != nil {
		return s
	}

	var err error
	// TODO Сделать Перемещение в рабочий каталог

	if err != nil {
		s.onError(err)
	}
	return s
}

func (s *syncVersion) WriteVersionFile() *syncVersion {

	if s.err != nil {
		return s
	}

	err := writeVersionFile(s.repository.workdir, strconv.FormatInt(s.version, 10))
	if err != nil {
		s.onError(err)
	}
	return s
}

func (s *syncVersion) Commit() *syncVersion {

	if s.err != nil {
		return s
	}

	err := commitVersionFile(s.repository.workdir)
	if err != nil {
		s.onError(err)
	}
	return s
}

type SyncRepository struct {
	repository.Repository

	workdir            string
	infobase           v8.Infobase
	currentVersion     int64
	maxVersion         int64
	repositoryVersions []repositoryVersion
	authors            map[string]string
}

func NewSyncRepository(path string) *SyncRepository {

	return &SyncRepository{
		Repository: repository.Repository{
			Path: path,
		},
	}
}

type SyncOptions struct {
	v8Path    string
	v8version string
	extention string
	infobase  v8.Infobase
}

type SyncOption func(*SyncOptions)

func (r *SyncRepository) Auth(user, passowrd string) {

	r.User = user
	r.Password = passowrd

}

func WithInfobase(connString, user, password string) SyncOption {

	return func(o *SyncOptions) {

		if len(connString) == 0 {
			return
		}

		o.infobase = syncInfobase(connString, user, password)

	}

}

func (r *SyncRepository) readAuthors() error {

	return nil

}

func (r *SyncRepository) readCurrentVersion() error {

	return nil

}

func (r *SyncRepository) sync(workdir string, opts *SyncOptions) error {

	r.workdir = workdir
	err := r.prepare(opts)

	if err != nil {
		return err
	}

	for _, rVersion := range r.repositoryVersions {

		if r.maxVersion < rVersion.version {
			break
		}

		// TODO
		// 1. Загрузить конфигурацию в базу
		// 2. Выгрузить во временный каталог
		// 3. Переместить в рабочий каталог
		// 4. Выполнить коммит в гит

		sVersion := &syncVersion{
			repository: r,
			version:    rVersion,
			oprions:    opts,
		}

		doErr := sVersion.Do()

		if doErr != nil {
			return doErr
		}

	}

	return nil
}

func (r *SyncRepository) prepare(opts *SyncOptions) error {

	// TODO Получение последний синхронизированный версии
	// TODO Получение таблицы версий хранилища 1С
	// TODO Сделать определени максимальной версии для разбора
	r.currentVersion = 0
	r.maxVersion = 0

	err := v8.Run(opts.infobase, repository.RepositoryReportOptions{
		Repository: r.Repository,
		File:       "", // TODO Чтение временного файла
		Extension:  opts.extention,
		NBegin:     r.maxVersion,
		GroupBy:    repository.REPOSITORY_GROUP_BY_COMMENT,
	})

	if err != nil {
		return err
	}

	r.repositoryVersions = []repositoryVersion{}

	if r.currentVersion < r.maxVersion {

		createErr := v8.Run(opts.infobase, designer.CreateFileInfoBaseOptions{})

		if createErr != nil {
			return createErr
		}

	}

	return nil
}

func (r *SyncRepository) Sync(workdir string, opts ...SyncOption) error {

	options := &SyncOptions{}

	for _, opt := range opts {
		opt(options)
	}

	return r.sync(workdir, options)

}

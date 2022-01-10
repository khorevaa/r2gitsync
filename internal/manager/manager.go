package manager

import (
	"fmt"
	"github.com/lithammer/shortuuid/v3"
	v8 "github.com/v8platform/api"
	"github.com/v8platform/designer"
)

func Sync(r SyncRepository, options Options) error {

	return r.Sync(options)

}

func Init(r SyncRepository, options Options) error {

	return r.Init(options)
}

type SyncRepository struct {
	designer.Repository
	Name    string
	Workdir string
}

func (r *SyncRepository) Sync(options Options) error {

	// jobLogger := log2.Logger(log2.NullLogger)
	//
	// if options.Logger != nil {
	// 	options.Logger.Debug("use options logger")
	// 	jobLogger = options.Logger.Named("manager")
	// }

	ib, err := r.getInfobase(options)
	if err != nil {
		return err
	}

	if len(r.Name) == 0 {
		r.Name = shortuuid.New()
	}

	jobName := fmt.Sprintf("Sync job for repository: %s (%s)", r.Name, r.Path)

	job := syncJob{
		name:         jobName,
		workdir:      r.Workdir,
		tempDir:      options.TempDir,
		repo:         &r.Repository,
		infobase:     ib,
		increment:    !options.DisableIncrement,
		minVersion:   options.MinVersion,
		maxVersion:   options.MaxVersion,
		limitVersion: options.LimitVersions,
		flow:         Flow{SubscribeManager: options.Plugins},
		options:      options.Options(),
		// log:          jobLogger,
		domainEmail: options.DomainEmail,
		skipFiles:   []string{VERSION_FILE, AUTHORS_FILE, ".git"},
	}

	return job.Run()

}

func (r *SyncRepository) Init(options Options) error {

	// jobLogger := log2.Logger(log2.NullLogger)
	//
	// if options.Logger != nil {
	// 	options.Logger.Debug("use options logger")
	// 	jobLogger = options.Logger.Named("manager")
	// }

	ib, err := r.getInfobase(options)
	if err != nil {
		return err
	}

	if len(r.Name) == 0 {
		r.Name = shortuuid.New()
	}

	jobName := fmt.Sprintf("Init job for repository: %s (%s)", r.Name, r.Path)

	job := syncJob{
		name:         jobName,
		workdir:      r.Workdir,
		tempDir:      options.TempDir,
		repo:         &r.Repository,
		infobase:     ib,
		increment:    !options.DisableIncrement,
		minVersion:   options.MinVersion,
		maxVersion:   options.MaxVersion,
		limitVersion: options.LimitVersions,
		flow:         Flow{SubscribeManager: options.Plugins},
		options:      options.Options(),
		// log:          jobLogger,
		domainEmail: options.DomainEmail,
		skipFiles:   []string{VERSION_FILE, AUTHORS_FILE, ".git"},
	}

	return job.Run()
}

func (r *SyncRepository) getInfobase(opts Options) (*v8.Infobase, error) {

	if len(opts.InfobaseConnect) > 0 {
		return opts.Infobase()
	}

	// logger.Debug("Creating temp infobase")

	CreateFileInfobase := v8.CreateFileInfobase(v8.NewTempDir(opts.TempDir, "temp_ib"))

	ib, err := v8.CreateInfobase(CreateFileInfobase, opts.Options())

	if err != nil {
		return nil, err
	}

	if len(r.Extension) > 0 {

		tempExtension, err := restoreTempExtension()
		if err != nil {
			return nil, err
		}

		LoadExtensionCfg := v8.LoadExtensionCfg(tempExtension, r.Extension)
		err = v8.Run(ib, LoadExtensionCfg, opts.Options())

		if err != nil {
			return nil, err
		}
		// logger.Debug("Empty extension loaded into infobase")
	}

	return ib, nil

}

package db

import (
	"github.com/elastic/go-ucfg"
	"github.com/khorevaa/logos"
	"github.com/khorevaa/r2gitsync/internal/services/db/ent"
	"github.com/khorevaa/r2gitsync/internal/services/db/repo"
	_ "github.com/lib/pq"
)

var log = logos.New("github.com/khorevaa/r2gitsync/services/repo")

type Repository struct {
	Projects       repo.IProjectRepository
	Plugins        repo.IPluginRepository
	PluginVersions repo.IPluginVersionRepository
	Storages       repo.IStorageRepository
	StorageCommits repo.IStorageCommitRepository
	StoragePlugins repo.IStoragePluginRepository
	Assets         repo.IAssetRepository
}

func New(cfg *ucfg.Config) (*Repository, error) {

	orm, err := connectDb(cfg)
	if err != nil {
		return nil, err
	}

	return &Repository{
		// Projects:       repo.NewProjectRepository(orm),
		// Plugins:        repo.NewPluginRepository(orm),
		// PluginVersions: repo.NewPluginVersionRepository(orm),
		// Assets:         repo.NewAssetRepository(orm),
		// Storages:       repo.NewStorageRepository(orm),
		// StorageCommits: repo.NewStorageCommitRepository(orm),
		// StoragePlugins: repo.NewStoragePluginRepository(orm),
	}, nil

}

func connectDb(cfg *ucfg.Config) (*ent.Client, error) {

	// var dialector gorm.Dialector
	config := Config{}

	err := cfg.Unpack(config)
	if err != nil {
		return nil, err
	}

	// logLevel := logger.Warn
	// if config.TraceSQLCommands {
	// 	logLevel = logger.Info
	// }

	client, err := ent.Open("postgres", "host=<host> port=<port> user=<user> dbname=<database> password=<pass>")
	if err != nil {
		log.Fatal(err.Error())
	}

	return client, nil
}

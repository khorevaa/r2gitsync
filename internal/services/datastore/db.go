package datastore

import (
	"context"

	"github.com/elastic/go-ucfg"
	"github.com/khorevaa/logos"
	"github.com/khorevaa/r2gitsync/internal/services/datastore/repo"
	"github.com/khorevaa/r2gitsync/internal/services/db"
	"github.com/khorevaa/r2gitsync/internal/services/db/migrate"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var log = logos.New("github.com/khorevaa/r2gitsync/services/repo").Sugar()

type Repository struct {
	Projects       repo.IProjectRepository
	Plugins        repo.IPluginRepository
	PluginVersions repo.IPluginVersionRepository
	Storages       repo.IStorageRepository
	StorageCommits repo.IStorageCommitRepository
	StoragePlugins repo.IStoragePluginRepository
	Assets         repo.IAssetRepository
	Orm            *db.Client
}

func New(cfg *ucfg.Config) (*Repository, error) {

	orm, err := connectDb(cfg)
	if err != nil {
		return nil, err
	}

	return &Repository{
		Projects: repo.NewProjectRepository(orm),
		// Plugins:        repo.NewPluginRepository(orm),
		// PluginVersions: repo.NewPluginVersionRepository(orm),
		// Assets:         repo.NewAssetRepository(orm),
		// Storages:       repo.NewStorageRepository(orm),
		// StorageCommits: repo.NewStorageCommitRepository(orm),
		// StoragePlugins: repo.NewStoragePluginRepository(orm),
	}, nil

}

func connectDb(cfg *ucfg.Config) (*db.Client, error) {

	// var dialector gorm.Dialector
	// config := Config{}
	//
	// err := cfg.Unpack(config)
	// if err != nil {
	// 	return nil, err
	// }

	// logLevel := logger.Warn
	// if config.TraceSQLCommands {
	// 	logLevel = logger.Info
	// }

	// client, err := db.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	// if err != nil {
	// 	log.Fatalf("failed opening connection to sqlite: %v", err)
	// }
	client, err := db.Open("postgres", "host=localhost port=5432 user=postgres dbname=r2gitsync password=passw0rd sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}
	ctx := context.Background()
	// Run migration.
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client, nil
}

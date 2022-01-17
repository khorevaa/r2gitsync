package db

import (
	"time"

	"github.com/elastic/go-ucfg"
	"github.com/khorevaa/logos"
	"github.com/khorevaa/r2gitsync/internal/services/db/repo"
	"github.com/khorevaa/r2gitsync/internal/services/db/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	var orm *gorm.DB

	orm, err := connectDb(cfg)
	if err != nil {
		return nil, err
	}

	return &Repository{
		Projects:       repo.NewProjectRepository(orm),
		Plugins:        repo.NewPluginRepository(orm),
		PluginVersions: repo.NewPluginVersionRepository(orm),
		Assets:         repo.NewAssetRepository(orm),
		Storages:       repo.NewStorageRepository(orm),
		StorageCommits: repo.NewStorageCommitRepository(orm),
		StoragePlugins: repo.NewStoragePluginRepository(orm),
	}, nil

}

func connectDb(cfg *ucfg.Config) (*gorm.DB, error) {

	var dialector gorm.Dialector
	config := Config{}

	err := cfg.Unpack(config)
	if err != nil {
		return nil, err
	}

	logLevel := logger.Warn
	if config.TraceSQLCommands {
		logLevel = logger.Info
	}

	return gorm.Open(dialector, &gorm.Config{
		Logger: utils.NewLogger(log, logger.Config{
			// временной зазор определения медленных запросов SQL
			SlowThreshold: time.Duration(config.SQLSlowThreshold) * time.Second,
			LogLevel:      logLevel,
			Colorful:      false,
		}),
		AllowGlobalUpdate: true,
	})
}

// applyAutoMigrations - регистрация авто миграции схемы бд из моделей
func applyAutoMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&repo.Asset{},
		&repo.Plugin{},
		&repo.PluginVersion{},
		&repo.PluginProperty{},
		&repo.Project{},
		&repo.Storage{},
		&repo.StorageCommit{},
		&repo.StoragePlugin{},
		&repo.StoragePluginProperty{},
	)
}

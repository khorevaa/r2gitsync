package di

import "github.com/khorevaa/r2gitsync/internal/services/datastore"

// IAppDeps - dependency injection container
type IAppDeps interface {
	DB() *datastore.Repository
	// ElasticRepo() *elastic.Repository
	// MongoRepo() *mongo.MongoRepo
	// CacheRepo() *cache.CacheRepo
	// MQ() *message_queue.RabbitMq
	// Log() *zap.Logger
}

func New(db *datastore.Repository) IAppDeps {
	return &DI{db: db}
}

type DI struct {
	db *datastore.Repository
	// elasticRepo *elastic.Repository
	// mongoRepo   *mongo.MongoRepo
	// cacheRepo   *cache.CacheRepo
	// mq          *message_queue.RabbitMq
	// logger      *zap.Logger
}

func (d *DI) DB() *datastore.Repository {
	return d.db
}

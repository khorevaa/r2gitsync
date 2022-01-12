package di

// IAppDeps - dependency injection container
type IAppDeps interface {
	// DBRepo() *db.DBRepo
	// ElasticRepo() *elastic.Repository
	// MongoRepo() *mongo.MongoRepo
	// CacheRepo() *cache.CacheRepo
	// MQ() *message_queue.RabbitMq
	// Log() *zap.Logger
}

type DI struct {
	// dbRepo      *db.DBRepo
	// elasticRepo *elastic.Repository
	// mongoRepo   *mongo.MongoRepo
	// cacheRepo   *cache.CacheRepo
	// mq          *message_queue.RabbitMq
	// logger      *zap.Logger
}

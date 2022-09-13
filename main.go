package main

import (
	"gitee.com/phper95/pkg/cache"
	"gitee.com/phper95/pkg/es"
	"gitee.com/phper95/pkg/logger"
	"gitee.com/phper95/pkg/shutdown"
	"gitee.com/phper95/pkg/trace"
	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"
	"product-consumer/conf"
	"product-consumer/global"
	"product-consumer/internal/consumer"
)

func init() {
	global.LoadConfig()
	global.LOG = global.SetupLogger()
	initRedisClient()
	initMongoClient()
	initESClient()

}

func initRedisClient() {
	redisCfg := global.CONFIG.Redis
	opt := redis.Options{
		Addr:        redisCfg.Host,
		Password:    redisCfg.Password,
		IdleTimeout: redisCfg.IdleTimeout,
	}
	redisTrace := trace.Cache{
		Name:                  "redis",
		SlowLoggerMillisecond: 500,
		Logger:                logger.GetLogger(),
		AlwaysTrace:           global.CONFIG.App.RunMode == conf.RunModeDev,
	}
	err := cache.InitRedis(cache.DefaultRedisClient, &opt, &redisTrace)
	if err != nil {
		global.LOG.Error("redis init error", zap.Error(err), "client", cache.DefaultRedisClient)
		panic("initRedisClient error")
	}
}

//初始化ES
func initESClient() {
	err := es.InitClientWithOptions(es.DefaultClient, global.CONFIG.Elasticsearch.Hosts,
		global.CONFIG.Elasticsearch.Username,
		global.CONFIG.Elasticsearch.Password,
		es.WithScheme("https"))
	if err != nil {
		global.LOG.Error("InitClientWithOptions error", err, "client", es.DefaultClient)
		panic(err)
	}
	global.ES = es.GetClient(es.DefaultClient)
}

func initMongoClient() {
	// TO DO ...
}

func main() {
	//开启消费者
	consumer.StartConsumer()

	//优雅关闭
	shutdown.NewHook().Close(
		func() {
			//kafka consumer
			consumer.CloseProductConsumer()
		},
		func() {
			//es
			es.CloseAll()
		},
	)

}

package main

import (
	"gitee.com/phper95/pkg/aws_s3"
	"gitee.com/phper95/pkg/cache"
	"gitee.com/phper95/pkg/es"
	"gitee.com/phper95/pkg/logger"
	"gitee.com/phper95/pkg/shutdown"
	"gitee.com/phper95/pkg/trace"
	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"
	"mail-consumer/conf"
	"mail-consumer/global"
	"mail-consumer/internal/consumer"
)

func init() {
	global.LoadConfig()
	global.LOG = global.SetupLogger()
	initRedisClient()
	initMongoClient()
	initS3()
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

func initS3() {
	err := aws_s3.InitService(aws_s3.DefaultClientName, global.CONFIG.S3.SK, "", global.CONFIG.S3.Region, global.CONFIG.S3.Host)
	if err != nil {
		global.LOG.Error("s3 init error", err, "client", aws_s3.DefaultClientName)
		panic(err)
	}
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
			consumer.CloseConsumer()
		},
		func() {
			//es
			es.CloseAll()
		},
	)

}

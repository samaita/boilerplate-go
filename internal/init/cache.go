package init

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/samaita/boilerplate-go/config"
	rds "github.com/samaita/boilerplate-go/pkg/redis"
)

type CacheConnection struct {
	MainCache *redis.Client
	Timeout   time.Duration
}

// ConnectCache creating new cache connection(s) from config
func ConnectCache(config config.Config) (conn CacheConnection) {
	redisConf := config.Datastore.Cache.Redis
	redisConn := rds.Redis{
		DB:       redisConf.DB,
		Host:     redisConf.Host,
		Password: redisConf.Password,
		Port:     redisConf.Port,
		Timeout:  redisConf.Timeout,
	}
	conn.MainCache = redisConn.Connect()

	conn.Timeout = redisConf.Timeout
	if err := conn.testConnection(); err != nil {
		log.Fatalln("Cache Err:", err)
	}

	log.Println("Success Connect:", fmt.Sprintf("%s:%s DB: %d", redisConf.Host, redisConf.Port, redisConf.DB))
	return
}

// testConnection do ping within predefined timeout, better to be called once on connect
func (conn *CacheConnection) testConnection() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), conn.Timeout)
	defer cancel()

	if err = conn.MainCache.Ping(ctx).Err(); err != nil {
		return
	}

	return
}

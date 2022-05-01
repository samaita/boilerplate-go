package redis

import (
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	DB       int
	Host     string
	Password string
	Port     string
	Timeout  time.Duration
}

// Connect create a new SQLX connection
func (rds *Redis) Connect() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     rds.Host + ":" + rds.Port,
		Password: rds.Password,
		DB:       rds.DB,
	})
}

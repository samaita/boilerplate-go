package repositories

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

// HealthRepo implements models.HealthRepository
type HealthRepo struct {
	DB    *sqlx.DB
	Redis *redis.Client
}

// NewHealthRepo ..
func NewHealthRepo(db *sqlx.DB, redis *redis.Client) *HealthRepo {
	return &HealthRepo{
		DB:    db,
		Redis: redis,
	}
}

func (h *HealthRepo) PingDB(ctx context.Context) (latency time.Duration, err error) {
	t := time.Now()
	err = h.DB.PingContext(ctx)
	latency = time.Since(t)
	return
}

func (h *HealthRepo) PingRedis(ctx context.Context) (latency time.Duration, err error) {
	t := time.Now()
	_, err = h.Redis.Ping(ctx).Result()
	latency = time.Since(t)
	return
}

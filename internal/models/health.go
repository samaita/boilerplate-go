package models

import (
	"context"
	"time"
)

type Health struct {
	DB    string
	Redis string
}

type HealthRepository interface {
	PingDB(ctx context.Context) (time.Duration, error)
	PingRedis(ctx context.Context) (time.Duration, error)
}

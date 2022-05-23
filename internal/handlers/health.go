package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/samaita/boilerplate-go/internal/dto"
	"github.com/samaita/boilerplate-go/internal/models"
	"github.com/samaita/boilerplate-go/pkg/constants"
	"github.com/samaita/boilerplate-go/pkg/http"
)

// HealthHandler hold every handler require
type HealthHandler struct {
	HealthRepo models.HealthRepository
}

// NewHealthHandler create new HealthHandler
func NewHealthHandler(healthRepo models.HealthRepository) *HealthHandler {
	return &HealthHandler{
		HealthRepo: healthRepo,
	}
}

// RegisterEndpoint return all handler registered for this app
func (h *HealthHandler) RegisterEndpoint(router *echo.Group) {
	router.GET("/health", h.HealthCheck)
}

func (h *HealthHandler) HealthCheck(ctx echo.Context) (err error) {
	newCtx := ctx.Request().Context()

	data := dto.HealthCheckResponse{
		DBHealth:    constants.PingResponseOK,
		CacheHealth: constants.PingResponseOK,
	}

	latency, err := h.HealthRepo.PingDB(newCtx)
	if err != nil {
		data.DBHealth = constants.PingResponseNotOK
		data.DBError = err.Error()
	}
	data.DBLatency = latency.String()

	latency, err = h.HealthRepo.PingRedis(newCtx)
	if err != nil {
		data.CacheHealth = constants.PingResponseNotOK
		data.CacheError = err.Error()
	}
	data.CacheLatency = latency.String()

	http.SuccessResponse(ctx, data)
	return
}

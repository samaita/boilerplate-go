package http

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/samaita/boilerplate-go/pkg/constants"
)

type Response struct {
	Latency interface{} `json:"latency"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(ctx echo.Context, data interface{}) {
	response(ctx, http.StatusOK, data)
}

func BadRequestResponse(ctx echo.Context, data interface{}) {
	response(ctx, http.StatusBadRequest, data)
}

func UnauthorizedResponse(ctx echo.Context, data interface{}) {
	response(ctx, http.StatusUnauthorized, data)
}

func ForbiddenResponse(ctx echo.Context, data interface{}) {
	response(ctx, http.StatusForbidden, data)
}

func UnavailableResponse(ctx echo.Context, data interface{}) {
	response(ctx, http.StatusServiceUnavailable, data)
}

func response(ctx echo.Context, code int, data interface{}) {
	response := Response{
		Data:    data,
		Latency: getLatency(ctx),
	}
	ctx.JSON(code, response)
}

func getLatency(ctx echo.Context) string {
	val := ctx.Get(constants.CtxLatency)
	if t, ok := val.(time.Time); ok {
		return time.Since(t).String()
	}
	return "-1"
}

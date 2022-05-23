package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samaita/boilerplate-go/pkg/constants"
)

func Recover() echo.MiddlewareFunc {
	return middleware.Recover()
}

func Latency(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(constants.CtxLatency, time.Now())
		return next(c)
	}
}

package logger

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"time"
)

func MiddlewareLogger(log zerolog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			req := c.Request()
			res := c.Response()

			// Process request
			c.SetLogger(&EchoLogger{z: Log})
			err := next(c)

			stop := time.Now()

			// Log after request is processed
			log.Info().
				Str("method", req.Method).
				Str("uri", req.RequestURI).
				Int("status", res.Status).
				Str("remote_ip", c.RealIP()).
				Str("latency", stop.Sub(start).String()).
				Msg("HTTP Request")

			return err
		}
	}
}

package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		e.Logger.Info("Health check")
		e.Logger.Debug("Debug log")

		return c.JSON(http.StatusOK, map[string]string{
			"status": "OK",
		})
	})
}

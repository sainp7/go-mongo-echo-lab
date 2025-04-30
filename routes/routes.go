package routes

import (
	"github.com/sainp7/go-mongo-echo-lab/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterHealthRoute(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		e.Logger.Info("Health check")
		e.Logger.Debug("Debug log")

		return c.JSON(http.StatusOK, map[string]string{
			"status": "OK",
		})
	})
}

func RegisterBookRoutes(g *echo.Group) {
	books := g.Group("/books")
	books.GET("", handlers.GetAllBooks)
}

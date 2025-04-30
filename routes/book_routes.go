package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/sainp7/go-mongo-echo-lab/handlers"
)

func RegisterBookRoutes(g *echo.Group) {
	books := g.Group("/books")
	books.GET("", handlers.GetBooksPaginated)
	books.POST("", handlers.CreateBook)
	books.GET("/:id", handlers.GetBookByID)
	books.DELETE("/:id", handlers.DeleteBookByID)
	books.PUT("/:id", handlers.UpdateBookByID)
}

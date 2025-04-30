package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sainp7/go-mongo-echo-lab/services"
	"net/http"
)

func GetAllBooks(c echo.Context) error {
	books, err := services.GetAllBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Something went wrong"})
	}
	if len(books) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "No books found"})
	}
	return c.JSON(http.StatusOK, books)
}

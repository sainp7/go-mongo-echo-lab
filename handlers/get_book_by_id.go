package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sainp7/go-mongo-echo-lab/services"
	"net/http"
)

func GetBookByID(c echo.Context) error {
	id := c.Param("id")
	c.Logger().Debugf("ID: %s", id)
	book, err := services.GetBookByID(id)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Book not found"})
	}
	c.Logger().Debugf("Book: %+v", book)
	return c.JSON(http.StatusOK, book)
}

package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sainp7/go-mongo-echo-lab/services"
	"net/http"
)

func DeleteBookByID(c echo.Context) error {
	id := c.Param("id")
	c.Logger().Debugf("ID: %s", id)

	if err := services.DeleteBookByID(id); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Book not found"})
	}
	c.Logger().Debugf("Book with ID: %s deleted successfully", id)
	return c.JSON(http.StatusOK, echo.Map{"message": "Book deleted successfully", "id": id})
}

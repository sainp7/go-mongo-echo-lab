package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sainp7/go-mongo-echo-lab/models"
	"github.com/sainp7/go-mongo-echo-lab/services"
	"net/http"
)

func UpdateBookByID(c echo.Context) error {
	id := c.Param("id")
	c.Logger().Debugf("ID: %s", id)

	var book models.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}
	if book.Title == "" || len(book.Authors) == 0 || book.ISBN == "" || book.Price <= 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Missing or invalid fields"})
	}
	c.Logger().Debugf("Book: %+v", book)

	updatedBook, err := services.UpdateBookByID(id, book)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Book not found"})
	}
	c.Logger().Debugf("Book updated: %+v", updatedBook)
	return c.JSON(http.StatusOK, updatedBook)
}

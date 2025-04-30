package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sainp7/go-mongo-echo-lab/models"
	"github.com/sainp7/go-mongo-echo-lab/services"
	"net/http"
)

func CreateBook(c echo.Context) error {
	var book models.Book
	if err := c.Bind(&book); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}
	if book.Title == "" || len(book.Authors) == 0 || book.ISBN == "" || book.Price <= 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Missing or invalid fields"})
	}
	c.Logger().Debugf("Book: %+v", book)
	insertedId, err := services.CreateBook(&book)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	c.Logger().Debugf("Inserted ID: %s", insertedId)
	return c.JSON(http.StatusCreated, echo.Map{"id": insertedId})
}

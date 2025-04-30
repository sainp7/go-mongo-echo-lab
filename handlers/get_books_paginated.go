package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sainp7/go-mongo-echo-lab/services"
	"math"
	"net/http"
	"strconv"
)

const defaultPageLimit = 3
const minPageLimit = 1
const maxPageLimit = 10

func GetBooksPaginated(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 0 {
		page = 0
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < minPageLimit {
		limit = defaultPageLimit
	}
	limit = max(min(limit, maxPageLimit), minPageLimit)

	c.Logger().Debugf("Page: %d, Limit: %d", page, limit)
	books, totalDocuments, err := services.GetBooksPaginated(page, limit)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Something went wrong"})
	}
	if len(books) == 0 {
		c.Logger().Warn("No books found")
		return c.JSON(http.StatusNotFound, echo.Map{"error": "No books found"})
	}
	c.Logger().Debugf("Books: %+v", books)
	return c.JSON(http.StatusOK, echo.Map{
		"data":       books,
		"page":       page,
		"limit":      limit,
		"total":      totalDocuments,
		"totalPages": int(math.Ceil(float64(totalDocuments) / float64(limit))),
	})
}

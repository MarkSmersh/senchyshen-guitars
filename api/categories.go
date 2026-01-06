package api

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func (s Server) CategoriesGet(c *gin.Context) {
	categories, err := s.Product.GetCategories()

	if err != nil {
		slog.Error(err.Error())
		c.String(500, "Nie.")
		return
	}

	c.JSON(200, categories)
}

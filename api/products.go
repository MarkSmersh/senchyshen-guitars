package api

import (
	"errors"
	"log/slog"
	"strconv"

	"github.com/MarkSmersh/senchyshen-guitars/models"
	"github.com/gin-gonic/gin"
)

// am i cooked?

func (s Server) ProductsGet(c *gin.Context) {
	req := models.FindManyParams{}

	err := c.BindJSON(&req)

	if err != nil {
		c.String(400, "Nieprawidlowe żądanie.")
		return
	}

	if req.Limit <= 0 || req.Limit > 100 {
		c.String(400, "Parametr limit w żądaniu musi byc więcej 0 i nie więcej 100")
		return
	}

	products, err := s.Product.FindMany(req)

	if err != nil {
		var apierr models.ApiError

		if errors.As(err, &apierr) {
			c.String(apierr.Code(), apierr.Error())
			return
		}

		slog.Error(err.Error())
		apierr = models.InternalServerError()
		c.String(500, apierr.Error())
		return
	}

	c.JSON(200, products)
}

func (s Server) ProductsGetSlug(c *gin.Context) {
	pidString := c.Param("pid")
	pid, _ := strconv.Atoi(pidString)
	model := models.NewProduct(s.Conn)
	product, err := model.Find(pid)

	if err != nil {
		var apierr models.ApiError

		if errors.As(err, &apierr) {
			c.String(apierr.Code(), apierr.Error())
			return
		}

		slog.Error(err.Error())
		apierr = models.InternalServerError()
		c.String(500, apierr.Error())
		return
	}

	c.JSON(200, product)
}

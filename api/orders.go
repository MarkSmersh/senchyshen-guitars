package api

import (
	"errors"
	"log/slog"

	"github.com/MarkSmersh/senchyshen-guitars/api/utils"
	"github.com/MarkSmersh/senchyshen-guitars/models"
	"github.com/gin-gonic/gin"
)

func (s Server) OrdersGet(c *gin.Context) {
	uuid := c.Param("uuid")

	order, err := s.Order.Find(uuid)

	if err != nil {
		var apierr models.ApiError
		if errors.As(err, &apierr) {
			c.String(apierr.Code(), apierr.Error())
			return
		}

		slog.Error(err.Error())
		apierr = models.InternalServerError()
		c.String(apierr.Code(), apierr.Error())
		return
	}

	c.JSON(200, order)
}

func (s Server) OrdersPost(c *gin.Context) {
	cartUUID, _ := utils.GetCart(c, s.Conn)

	req := models.OrderCreate{CartUUID: cartUUID}

	if err := c.BindJSON(&req); err != nil {
		c.String(400, "Nieprawidlowe żądanie")
		return
	}

	if len(req.Tel) <= 0 {
		c.String(400, "Nie podany numer telefonu")
		return
	}

	model := models.NewOrder(s.Conn)

	uuid, err := model.Create(req)

	if err != nil {
		slog.Error(err.Error())
		c.String(500, "Wystarczy.")
		return
	}

	c.String(201, uuid)
}

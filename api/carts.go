package api

import (
	"log/slog"

	"github.com/MarkSmersh/senchyshen-guitars/api/utils"
	"github.com/MarkSmersh/senchyshen-guitars/models"
	"github.com/gin-gonic/gin"
)

func (s Server) CartsGet(c *gin.Context) {
	cartUuid, _ := utils.GetCart(c, s.Conn)

	products, _ := s.Cart.GetProducts(cartUuid)

	c.JSON(200, products)
}

func (s Server) CartsPost(c *gin.Context) {
	cartUuid, _ := utils.GetCart(c, s.Conn)

	req := models.CartAddProduct{UUID: cartUuid}

	if err := c.BindJSON(&req); err != nil {
		c.String(400, "Nieprawidlowe żądanie")
		return
	}

	if req.Count < 1 {
		c.String(400, "Ilość produktów nie może być mniej czym 1. Proszę nie bawić z API.")
		return
	}

	err := s.Cart.AddProduct(req)

	if err != nil {
		slog.Error(err.Error())
		c.String(400, err.Error())
		return
	}

	c.String(201, "Produkt jest dodano do koszyka")
}

func (s Server) CartsDelete(c *gin.Context) {
	cartUuid, _ := utils.GetCart(c, s.Conn)

	pid := c.Param("pid")

	// if err := c.BindJSON(&req); err != nil {
	// 	c.String(400, "Nieprawidlowe żądanie")
	// 	return
	// }

	err := s.Cart.RemoveProduct(cartUuid, pid)

	if err != nil {
		slog.Error(err.Error())
		c.String(400, err.Error())
		return
	}

	c.String(200, "Produkt jest usunięto")
}

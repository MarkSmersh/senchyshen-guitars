package api

import (
	"errors"

	"github.com/MarkSmersh/senchyshen-guitars/api/utils"
	"github.com/MarkSmersh/senchyshen-guitars/models"
	"github.com/gin-gonic/gin"
)

func (s Server) ConstructorsPost(c *gin.Context) {
	cartUUID, _ := utils.GetCart(c, s.Conn)

	params := models.ConstructorCreate{}

	if err := c.BindJSON(&params); err != nil {
		c.String(400, "Nieprawidlowe żądanie")
		return
	}

	if len(params.Title) <= 0 {
		c.String(400, "Brak nazwy guitary")
		return
	}

	if params.Color == "" {
		c.String(400, "Brak koloru guitary")
		return
	}

	if params.BodyshapeID <= 0 {
		c.String(400, "Brak kształtu guitary")
		return
	}

	if len(params.Pickups) < 1 {
		c.String(400, "Minimalna ilość przetwoników gitarowych nie może być mniej 1")
		return
	}

	if len(params.Pickups) > 3 {
		c.String(400, "Maksymalna ilość przetwoników gitarowych nie może być więcej 3")
		return
	}

	for i, p := range params.Pickups {
		if p.Position == "" {
			c.String(400, "Brak pozycji dla przetwornika gitarowego #%d", i+1)
			return
		}

		if p.PickupID <= 0 {
			c.String(400, "Brak identyfikora przetwornika gitarowego #%d", i+1)
			return
		}
	}

	pid, err := s.Product.CreateConstructor(params)

	if err != nil {
		var apierr models.ApiError

		if errors.As(err, &apierr) {
			c.String(apierr.Code(), apierr.Error())
			return
		}

		apierr = models.InternalServerError()
		c.String(apierr.Code(), apierr.Error())
		return
	}

	productParams := models.CartAddProduct{
		UUID:      cartUUID,
		ProductID: pid,
		Count:     1,
	}

	err = s.Cart.AddProduct(productParams)

	if err != nil {
		var apierr models.ApiError

		if errors.As(err, &apierr) {
			c.String(apierr.Code(), apierr.Error())
			return
		}

		apierr = models.InternalServerError()
		c.String(apierr.Code(), apierr.Error())
		return
	}

	c.String(201, "Stworzono guitarę oraz dodano ją do koszyka")
}

package api

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/MarkSmersh/senchyshen-guitars/models"
	"github.com/MarkSmersh/senchyshen-guitars/utils"
	"github.com/gin-gonic/gin"
)

func (s Server) CategoriesPost(c *gin.Context) {
	categories := []models.CreateCategoryReq{}

	if err := c.BindJSON(&categories); err != nil {
		slog.Error(err.Error())
		c.String(400, "Nieprawidlowe żądanie")
		return
	}

	for i, cat := range categories {
		if cat.Title == "" || cat.Description == "" || cat.Image == "" {
			c.String(400, fmt.Sprintf("Produkt #%d. Brak potzebnych parametrów: title, description, image", i))
			return
		}

		id, err := s.Product.CreateCategory(cat)

		if err != nil {
			var apierr models.ApiError

			if !errors.As(err, &apierr) {
				apierr = models.InternalServerError()
			}

			c.String(apierr.Code(), fmt.Sprintf("Produkt #%d. %s", i, apierr.Error()))
			return
		}

		uuid, ext, err := utils.SaveUrlImage(cat.Image, "assets")

		if err != nil {
			c.String(400, "Produkt #%d. %s", i, err.Error())
			return
		}

		err = s.Image.CategoryCreate(uuid, ext, fmt.Sprintf("%d", id))

		if err != nil {
			c.String(400, "Produkt #%d. %s", i, err.Error())
			return
		}
	}

	c.String(201, fmt.Sprintf("Stworzono %d kategorii", len(categories)))
}

func (s Server) ProductsPost(c *gin.Context) {
	products := []models.ProductModel{}

	if err := c.BindJSON(&products); err != nil {
		slog.Error(err.Error())
		c.String(400, "Nieprawidlowe żądanie")
		return
	}

	for i, p := range products {
		if p.Description == "" || p.Title == "" || p.Type == "" || p.Price <= 0 {
			c.String(400, "Brak potzebujących parametrów: title, description, type, price")
		}

		id, err := s.Product.Create(p)

		if err != nil {
			var apierr models.ApiError

			if !errors.As(err, &apierr) {
				apierr = models.InternalServerError()
			}

			c.String(apierr.Code(), fmt.Sprintf("Produkt #%d. %s", i, apierr.Error()))
			return
		}

		for j, img := range p.Images {
			if img.Url != "" {
				name, ext, err := utils.SaveUrlImage(img.Url, "assets")

				if err != nil {
					c.String(400, "Produkt #%d. Zdjęcie #%d. %s", i, j, err.Error())
					return
				}

				err = s.Image.ProductCreate(name, ext, fmt.Sprintf("%d", id))

				if err != nil {
					c.String(400, "Produkt #%d. Zdjęcie #%d. %s", i, j, err.Error())
					return
				}
			}
		}
	}

	c.String(201, fmt.Sprintf("Stworzono %d produktów", len(products)))
}

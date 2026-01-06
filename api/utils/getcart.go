package utils

import (
	"github.com/MarkSmersh/senchyshen-guitars/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetCart(c *gin.Context, conn *pgxpool.Pool) (string, error) {
	cartId, err := c.Cookie("cart_id")

	if err != nil {
		model := models.NewCart(conn)

		cartId, err = model.Create()

		if err != nil {
			return "", err
		}

		c.SetCookie(
			"cart_id", cartId,
			0, "/", "", true, true,
		)
	}

	return cartId, nil
}

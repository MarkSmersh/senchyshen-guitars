package api

import (
	"context"
	"log/slog"

	"github.com/MarkSmersh/senchyshen-guitars/models"
	"github.com/gin-gonic/gin"
)

type Preview struct {
	Bodyshapes int `json:"bodyshapes"`
	Pickups    int `json:"pickups"`
}

func (s Server) PreviewGet(c *gin.Context) {
	row := s.Conn.QueryRow(
		context.Background(),
		"select count(distinct b.id), count(distinct p.id) from bodyshapes b join pickups p on true",
	)

	body := Preview{}

	if err := row.Scan(&body.Bodyshapes, &body.Pickups); err != nil {
		slog.Error(err.Error())
		apierr := models.InternalServerError()
		c.String(apierr.Code(), apierr.Error())
		return
	}

	c.JSON(200, body)
}

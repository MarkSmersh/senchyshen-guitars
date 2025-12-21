package api

import "github.com/gin-gonic/gin"

func (s Server) HealthyGet(c *gin.Context) {
	c.String(200, "Server is healthy")
}

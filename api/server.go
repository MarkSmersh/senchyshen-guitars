package api

import (
	"fmt"

	"github.com/MarkSmersh/senchyshen-guitars/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Server struct {
	Engine  *gin.Engine
	R       *gin.RouterGroup
	Conn    *pgx.Conn
	Address string
	Port    int
	Cart    models.Cart
	Product models.Product
	Order   models.Order
}

func NewServer(conn *pgx.Conn, address string, port int) Server {
	engine := gin.Default()
	router := engine.Group("/api")

	return Server{
		Engine:  engine,
		R:       router,
		Conn:    conn,
		Address: address,
		Port:    port,
		Cart:    models.NewCart(conn),
		Product: models.NewProduct(conn),
		Order:   models.NewOrder(conn),
	}
}

func (s *Server) Start() error {
	if err := s.Engine.Run(
		fmt.Sprintf("%s:%d", s.Address, s.Port),
	); err != nil {
		return err
	}

	return nil
}

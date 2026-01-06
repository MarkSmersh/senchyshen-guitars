package api

import (
	"fmt"

	"github.com/MarkSmersh/senchyshen-guitars/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	Engine  *gin.Engine
	R       *gin.RouterGroup
	Conn    *pgxpool.Pool
	Address string
	Port    int
	Cart    models.Cart
	Product models.Product
	Order   models.Order
	Image   models.Image
}

func NewServer(conn *pgxpool.Pool, address string, port int) Server {
	engine := gin.Default()

	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOriginFunc = func(origin string) bool { return true }

	corsConfig.AllowCredentials = true

	engine.Use(cors.New(corsConfig))

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
		Image:   models.NewImage(conn),
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

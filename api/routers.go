package api

func (s *Server) SetRouters() {
	s.AssetsRouter()
	s.HealthyRouter()
	s.ProductsRouter()
	s.ImagesRouter()
	s.OrdersRouter()
	s.CartsRouter()
	s.ConstuctorsRouter()
	s.PreviewRouter()
}

func (s *Server) HealthyRouter() {
	g := s.R.Group("/healthy")

	g.GET("/", s.HealthyGet)
}

func (s *Server) ProductsRouter() {
	g := s.R.Group("/products")

	g.POST("/", s.ProductsGet)
	g.GET("/:pid", s.ProductsGetSlug)
}

func (s *Server) ImagesRouter() {
	g := s.R.Group("/images")

	g.POST("/products/:pid", s.ImagesProductsPost)
	g.POST("/categories/:cid", s.ImagesCategoriesPost)

	g.DELETE("/:id", s.ImagesDelete)
}

func (s *Server) OrdersRouter() {
	g := s.R.Group("/orders")

	g.GET("/:uuid", s.OrdersGet)
	g.POST("/", s.OrdersPost)
}

func (s *Server) CartsRouter() {
	g := s.R.Group("/carts")

	g.GET("/", s.CartsGet)
	g.POST("/", s.CartsPost)
	g.DELETE("/:pid", s.CartsDelete)
}

func (s *Server) ConstuctorsRouter() {
	g := s.R.Group("/constructors")

	g.POST("/", s.ConstructorsPost)
}

func (s *Server) PreviewRouter() {
	g := s.R.Group("/preview")

	g.GET("/", s.PreviewGet)
}

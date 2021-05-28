package controllers

func (s *Server) InitializeRoutes() {
	// Home Route
	s.Router.GET("/", s.Home)
	s.Router.POST("/users", s.CreateUser)
}

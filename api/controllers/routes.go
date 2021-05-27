package controllers

func (s *Server) InitializeRoutes() {
	// Home Route
	s.Router.GET("/", s.Home)
}

package service

func (s *Service) SetRouter() {
	s.router.GET("/", s.handler.Health)
	s.router.GET("/health", s.handler.Health)
	s.router.POST("/auth/register", s.handler.Register)
	s.router.POST("/auth/login", s.handler.Login)
	s.router.POST("/auth/refresh", s.handler.RefreshToken)
	s.router.GET("/auth/me", s.handler.AuthorizeJWT, s.handler.GetUser)
	user := s.router.Group("/user", s.handler.AuthorizeJWT)
	user.GET("/:id", s.handler.GetUser)
	user.DELETE("/delete/", s.handler.DeleteUser)
	users := s.router.Group("/users", s.handler.AuthorizeJWT)
	users.GET("/:rows/:page", s.handler.ListUsers)
}

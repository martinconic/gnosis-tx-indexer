package api

func (s *Service) InitializeRoutes() {
	s.router.GET("/api/status", Status)
}

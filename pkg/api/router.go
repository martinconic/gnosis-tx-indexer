package api

func (s *Service) InitializeRoutes() {
	s.router.GET("/api/status", Status)
	s.router.GET("/api/transactions/:address", GetTransactions)
}

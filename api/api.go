package api

import "github.com/gin-gonic/gin"

type ServiceApi struct {
	Engine *gin.Engine
}

func New() *ServiceApi {
	return &ServiceApi{
		Engine: gin.Default(),
	}
}

func (s *ServiceApi) Start(addr ...string) {
	s.GetRoutes()
	s.Engine.Run(addr...)
}

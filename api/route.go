package api

import (
	infra_controller "github.com/dyhalmeida/golang-students-api/api/controller/infra"
)

func (s *ServiceApi) GetRoutes() {
	s.Engine.GET("/heart", infra_controller.HeartHandler)
}

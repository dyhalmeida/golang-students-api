package api

import (
	infra_controller "github.com/dyhalmeida/golang-students-api/api/controller/infra"
	student_controller "github.com/dyhalmeida/golang-students-api/api/controller/student"
)

func (s *ServiceApi) GetRoutes() {
	s.Engine.GET("/heart", infra_controller.HeartHandler)

	studentGroupV1 := s.Engine.Group("/api/v1")
	{
		studentGroupV1.GET("/student", student_controller.List)
		studentGroupV1.POST("/student", student_controller.Create)
	}
}

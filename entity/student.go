package entity

import "github.com/dyhalmeida/golang-students-api/entity/shared"

type Student struct {
	ID       string `json:"id"`
	Fullname string `json:"fullname"`
	Age      int64  `json:"age"`
}

var Students = []Student{
	{ID: shared.GetID().String(), Fullname: "Diego", Age: 31},
	{ID: shared.GetID().String(), Fullname: "Mayara", Age: 27},
}

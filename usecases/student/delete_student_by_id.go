package student

import (
	"fmt"

	"github.com/dyhalmeida/golang-students-api/entity"
)

func DeleteStudentByID(id string) (*entity.Student, error) {
	indexStudentFound := -1

	var student *entity.Student

	for i, s := range entity.Students {
		if s.ID == id {
			student = &s
			indexStudentFound = i
			break
		}
	}

	if indexStudentFound == -1 {
		return nil, fmt.Errorf("student with id: %s not found", id)
	}

	entity.Students = append(entity.Students[:indexStudentFound], entity.Students[indexStudentFound+1:]...)

	return student, nil
}

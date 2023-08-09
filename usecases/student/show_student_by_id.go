package student

import (
	"fmt"

	"github.com/dyhalmeida/golang-students-api/entity"
)

func ShowStudentByID(id string) (*entity.Student, error) {
	var student *entity.Student

	for _, s := range entity.Students {
		if s.ID == id {
			student = &s
			break
		}
	}

	if student == nil {
		return nil, fmt.Errorf("student with id: %s not found", id)
	}

	return student, nil
}

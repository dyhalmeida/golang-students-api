package shared

import (
	"github.com/google/uuid"
)

func GetID() uuid.UUID {
	return uuid.New()
}

func GetIDByString(s string) (uuid.UUID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return id, err
	}
	return id, nil
}

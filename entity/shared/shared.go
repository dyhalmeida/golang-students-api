package shared

import (
	"github.com/google/uuid"
)

func GetID() uuid.UUID {
	return uuid.New()
}

func GetIDByString(s string) (string, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return id.String(), err
	}
	return id.String(), nil
}

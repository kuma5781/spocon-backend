//go:generate mockgen -source=$GOFILE -destination=mocks/mock_$GOFILE -package=mocks
package util

import "github.com/google/uuid"

type UUIDv7Generator interface {
	Generate() (string, error)
}

type UUIDv7GeneratorImpl struct {
}

func (*UUIDv7GeneratorImpl) Generate() (string, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

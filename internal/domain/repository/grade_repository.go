//go:generate mockgen -source=$GOFILE -destination=mocks/mock_$GOFILE -package=mocks

package repository

import (
	g "spocon-backend/internal/domain/model/grade"
)

type GradeRepository interface {
	FetchAll() ([]g.Grade, error)
}

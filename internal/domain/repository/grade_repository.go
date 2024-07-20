package repository

import (
	g "spocon-backend/internal/domain/model/grade"
)

type GradeRepository interface {
	FetchAll() ([]g.GradeEntity, error)
}

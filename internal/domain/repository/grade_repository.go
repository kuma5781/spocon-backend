package repository

import (
	"database/sql"
	g "spocon-backend/internal/domain/model/grade"
	"spocon-backend/internal/infra"
)

type GradeRepository interface {
	FetchAll() ([]g.GradeEntity, error)
}

func NewGradeRepository(db *sql.DB) GradeRepository {
	return &infra.GradeRepositoryImpl{DB: db}
}

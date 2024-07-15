package domain

import (
	"database/sql"
	domain "spocon-backend/internal/domain/model"
	"spocon-backend/internal/infra"
)

type GradeRepository interface {
	FetchAll() ([]domain.GradeEntity, error)
}

func NewGradeRepository(db *sql.DB) GradeRepository {
	return &infra.GradeRepositoryImpl{DB: db}
}

package app

import (
	"database/sql"
	r "spocon-backend/internal/domain/repository"
)

func NewRepositories(db *sql.DB) *r.Repositories {
	return &r.Repositories{
		GradeRepository: r.NewGradeRepository(db),
	}
}

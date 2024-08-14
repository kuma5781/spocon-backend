package infra

import (
	"database/sql"
	r "spocon-backend/internal/domain/repository"
)

type RepositoryImpls struct {
}

func NewRepositories(db *sql.DB) *r.Repositories {
	return &r.Repositories{
		GradeRepository: NewGradeRepositoryImpl(db),
	}
}

package repository

import "database/sql"

type Repositories struct {
	GradeRepository GradeRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		GradeRepository: NewGradeRepository(db),
	}
}

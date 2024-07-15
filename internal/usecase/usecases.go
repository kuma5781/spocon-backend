package usecase

import (
	r "spocon-backend/internal/domain/repository"
)

type Usecases struct {
	GradeUsecase GradeUsecase
}

func NewUsecases(r *r.Repositories) *Usecases {
	return &Usecases{
		GradeUsecase: NewGradeUsecase(r.GradeRepository),
	}
}

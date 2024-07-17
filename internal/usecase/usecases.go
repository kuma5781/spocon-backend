package usecase

import (
	r "spocon-backend/internal/domain/repository"
)

type Usecases struct {
	GradeUsecase GradeUsecase
	TeamUsecase  TeamUsecase
}

func NewUsecases(r *r.Repositories) *Usecases {
	return &Usecases{
		GradeUsecase: NewGradeUsecase(r.GradeRepository),
	}
}

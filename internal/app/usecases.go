package app

import (
	r "spocon-backend/internal/domain/repository"
	"spocon-backend/internal/usecase"
)

func NewUsecases(r *r.Repositories) *usecase.Usecases {
	return &usecase.Usecases{
		GradeUsecase: usecase.NewGradeUsecase(r.GradeRepository),
	}

}

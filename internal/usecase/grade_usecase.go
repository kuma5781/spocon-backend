package usecase

import (
	g "spocon-backend/internal/domain/model/grade"
	r "spocon-backend/internal/domain/repository"
)

type GradeUsecase struct {
	GradeRepository r.GradeRepository
}

func NewGradeUsecase(r r.GradeRepository) GradeUsecase {
	return GradeUsecase{
		GradeRepository: r,
	}
}

func (u *GradeUsecase) GetGrades() ([]g.Grade, error) {
	grades, err := u.GradeRepository.FetchAll()
	if err != nil {
		return nil, err
	}

	return grades, nil
}

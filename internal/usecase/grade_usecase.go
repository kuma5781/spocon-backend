package usecase

import (
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

func (u *GradeUsecase) GetGrades() ([]GradeDto, error) {
	entities, err := u.GradeRepository.FetchAll()
	if err != nil {
		return nil, err
	}
	grades := make([]GradeDto, len(entities))
	for i, entity := range entities {
		grades[i] = from(&entity)
	}

	return grades, nil
}

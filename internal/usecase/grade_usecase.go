package usecase

import (
	r "spocon-backend/internal/domain/repository"
	"spocon-backend/internal/usecase/dto"
)

type GradeUsecase struct {
	GradeRepository r.GradeRepository
}

func NewGradeUsecase(r r.GradeRepository) GradeUsecase {
	return GradeUsecase{
		GradeRepository: r,
	}
}

func (u *GradeUsecase) GetGrades() ([]dto.GradeOutput, error) {
	entities, err := u.GradeRepository.FetchAll()
	if err != nil {
		return nil, err
	}
	grades := make([]dto.GradeOutput, len(entities))
	for i, entity := range entities {
		grades[i] = dto.From(&entity)
	}

	return grades, nil
}

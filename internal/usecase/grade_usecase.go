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
	grades, err := u.GradeRepository.FetchAll()
	if err != nil {
		return nil, err
	}
	outputs := make([]dto.GradeOutput, len(grades))
	for i, grade := range grades {
		outputs[i] = dto.From(&grade)
	}

	return outputs, nil
}

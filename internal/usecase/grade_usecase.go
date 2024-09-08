package usecase

import (
	g "spocon-backend/internal/domain/model/grade"
	r "spocon-backend/internal/domain/repository"

	"github.com/pkg/errors"
)

type GradeUsecase interface {
	GetGrades() ([]g.Grade, error)
}

func NewGradeUsecase(r r.GradeRepository) GradeUsecase {
	return &GradeUsecaseImpl{
		GradeRepository: r,
	}
}

type GradeUsecaseImpl struct {
	GradeRepository r.GradeRepository
}

func (u *GradeUsecaseImpl) GetGrades() ([]g.Grade, error) {
	grades, err := u.GradeRepository.FetchAll()
	if err != nil {
		return nil, errors.Wrap(err, "gradeのFetchAll()に失敗しました。")
	}

	return grades, nil
}

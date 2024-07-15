package usecase

import m "spocon-backend/internal/domain/model"

type GradeDto struct {
	Id   string
	Name string
}

func from(e *m.GradeEntity) GradeDto {
	return GradeDto{
		Id:   e.Id,
		Name: e.Name,
	}
}

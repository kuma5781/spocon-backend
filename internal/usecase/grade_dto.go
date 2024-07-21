package usecase

import g "spocon-backend/internal/domain/model/grade"

type GradeDto struct {
	Id   string
	Name string
}

func from(e *g.Grade) GradeDto {
	return GradeDto{
		Id:   e.Id,
		Name: e.Name,
	}
}

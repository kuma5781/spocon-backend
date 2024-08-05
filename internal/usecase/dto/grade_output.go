package dto

import g "spocon-backend/internal/domain/model/grade"

type GradeOutput struct {
	Id   g.GradeId
	Name string
}

func From(g *g.Grade) GradeOutput {
	return GradeOutput{
		Id:   g.Id,
		Name: g.Name,
	}
}

package adapter

import (
	g "spocon-backend/internal/domain/model/grade"
	"spocon-backend/internal/openapi"

	"github.com/samber/lo"
)

func ToGetGradesResponse(list []g.Grade) *openapi.GetGradesResponse {
	grades := lo.Map(list, func(g g.Grade, _ int) openapi.Grade {
		return openapi.Grade{
			Id:   g.Id.Value(),
			Name: g.Name,
		}
	})
	return &openapi.GetGradesResponse{
		Grades: grades,
	}
}

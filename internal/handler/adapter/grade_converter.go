package adapter

import (
	"spocon-backend/internal/openapi"
	"spocon-backend/internal/usecase/dto"

	"github.com/samber/lo"
)

func ToGetGradesResponse(gradeDtoList []dto.GradeOutput) *openapi.GetGradesResponse {
	grades := lo.Map(gradeDtoList, func(g dto.GradeOutput, _ int) openapi.Grade {
		return openapi.Grade{
			Id:   g.Id,
			Name: g.Name,
		}
	})
	return &openapi.GetGradesResponse{
		Grades: grades,
	}
}

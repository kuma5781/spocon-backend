package handler

import (
	"net/http"
	"spocon-backend/internal/openapi"
	"spocon-backend/internal/usecase"

	"github.com/labstack/echo/v4"

	"github.com/samber/lo"
)

type GradeHandler struct {
	GradeUsecase usecase.GradeUsecase
}

func (h *Handlers) GetGrades(c echo.Context) error {
	grades, err := h.GradeHandler.GradeUsecase.GetGrades()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	res := toGetGradesResponse(grades)
	return c.JSON(http.StatusOK, res)
}

func toGetGradesResponse(gradeDtoList []usecase.GradeDto) *openapi.GetGradesResponse {
	grades := lo.Map(gradeDtoList, func(gradeDto usecase.GradeDto, _ int) openapi.Grade {
		return openapi.Grade{
			Id:   gradeDto.Id,
			Name: gradeDto.Name,
		}
	})
	return &openapi.GetGradesResponse{
		Grades: grades,
	}
}

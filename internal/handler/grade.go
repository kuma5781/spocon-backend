package handler

import (
	"net/http"
	"spocon-backend/internal/openapi"
	"spocon-backend/internal/usecase"

	"github.com/labstack/echo/v4"
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

func toGetGradesResponse(grades []usecase.GradeDto) *openapi.GetGradesResponse {
	var res openapi.GetGradesResponse
	for _, grade := range grades {
		res.Grades = append(res.Grades, openapi.Grade{
			Id:   grade.Id,
			Name: grade.Name,
		},
		)
	}
	return &res
}

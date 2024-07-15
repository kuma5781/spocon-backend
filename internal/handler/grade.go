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

func (h *Handlers) GetGrades(ctx echo.Context) error {
	// ctx := c.Request().Context()
	grades, err := h.GradeHandler.GradeUsecase.GetGrades()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	res := toGetGradesResponse(grades)
	return ctx.JSON(http.StatusOK, res)
}

func toGetGradesResponse(grades []usecase.GradeDto) *openapi.GetGradesResponse {
	var res openapi.GetGradesResponse
	for _, grade := range grades {
		res.Grades = append(res.Grades, struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		}{
			Id:   grade.Id,
			Name: grade.Name,
		},
		)
	}
	return &res
}

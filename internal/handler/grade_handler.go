package handler

import (
	"net/http"
	a "spocon-backend/internal/handler/adapter"
	"spocon-backend/internal/usecase"

	"github.com/labstack/echo/v4"
)

type GradeHandler struct {
	GradeUsecase usecase.GradeUsecase
}

func NewGradeHandler(u usecase.GradeUsecase) GradeHandler {
	return GradeHandler{
		GradeUsecase: u,
	}
}

func (h *Handlers) GetGrades(c echo.Context) error {
	grades, err := h.GradeHandler.GradeUsecase.GetGrades()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	res := a.ToGetGradesResponse(grades)
	return c.JSON(http.StatusOK, res)
}

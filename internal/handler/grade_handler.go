package handler

import (
	"net/http"
	conv "spocon-backend/internal/handler/converter"
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
	res := conv.ToGetGradesResponse(grades)
	return c.JSON(http.StatusOK, res)
}

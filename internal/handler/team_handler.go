package handler

import (
	"net/http"
	a "spocon-backend/internal/handler/adapter"
	"spocon-backend/internal/openapi"
	"spocon-backend/internal/usecase"

	"github.com/labstack/echo/v4"
)

type TeamHandler struct {
	TeamUsecase usecase.TeamUsecase
}

func (h *Handlers) CreateTeam(c echo.Context) error {
	var body openapi.CreateTeamRequest
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	input := a.ToCreateTeamInput(&body)
	if err := h.TeamHandler.TeamUsecase.CreateTeam(input); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

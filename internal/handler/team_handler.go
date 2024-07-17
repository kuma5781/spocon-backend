package handler

import (
	"net/http"
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
	input := toCreateTeamInput(&body)
	if err := h.TeamHandler.TeamUsecase.CreateTeam(input); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

func toCreateTeamInput(req *openapi.CreateTeamRequest) *usecase.TeamCreateInput {
	return &usecase.TeamCreateInput{
		TeamName:     req.Name,
		SportId:      req.SportId,
		GradeId:      req.GradeId,
		IconPath:     req.IconPath,
		Description:  req.Description,
		AddressCity:  req.AddressCity,
		AddressState: req.AddressState,
	}
}

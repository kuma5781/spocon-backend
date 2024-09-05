package handler

import (
	"net/http"
	e "spocon-backend/internal/errors"
	a "spocon-backend/internal/handler/adapter"
	"spocon-backend/internal/openapi"
	"spocon-backend/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type TeamHandler struct {
	TeamUsecase usecase.TeamUsecase
}

func NewTeamHandler(u usecase.TeamUsecase) TeamHandler {
	return TeamHandler{
		TeamUsecase: u,
	}
}

func (h *Handlers) CreateTeam(c echo.Context) error {
	var body openapi.CreateTeamRequest
	if err := c.Bind(&body); err != nil {
		return errors.Wrap(e.InvalidRequestParamsError, err.Error())
	}
	input := a.ToCreateTeamInput(&body)
	result, err := h.TeamHandler.TeamUsecase.CreateTeam(input)
	if err != nil {
		return errors.Wrap(err, "チーム登録に失敗しました。")
	}
	return c.JSON(http.StatusCreated, a.ToCreateTeamOutput(result))

}

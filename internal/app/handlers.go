package app

import (
	"spocon-backend/internal/handler"
	"spocon-backend/internal/usecase"
)

func NewHandlers(u *usecase.Usecases) *handler.Handlers {
	return &handler.Handlers{
		HealthHandler: handler.HealthHandler{},
		UserHandler:   handler.UserHandler{},
		GradeHandler:  handler.GradeHandler{GradeUsecase: u.GradeUsecase},
	}
}

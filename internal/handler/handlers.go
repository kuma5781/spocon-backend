package handler

import "spocon-backend/internal/usecase"

type Handlers struct {
	HealthHandler HealthHandler
	GradeHandler  GradeHandler
	TeamHandler   TeamHandler
}

func NewHandlers(u *usecase.Usecases) *Handlers {
	return &Handlers{
		HealthHandler: HealthHandler{},
		GradeHandler:  GradeHandler{GradeUsecase: u.GradeUsecase},
		TeamHandler:   TeamHandler{TeamUsecase: u.TeamUsecase},
	}
}

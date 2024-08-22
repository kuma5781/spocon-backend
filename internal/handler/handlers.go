package handler

import "spocon-backend/internal/usecase"

type Handlers struct {
	HealthHandler HealthHandler
	GradeHandler  GradeHandler
	TeamHandler   TeamHandler
}

func NewHandlers(u *usecase.Usecases) *Handlers {
	return &Handlers{
		HealthHandler: NewHealthHandler(),
		GradeHandler:  NewGradeHandler(u.GradeUsecase),
		TeamHandler:   NewTeamHandler(u.TeamUsecase),
	}
}

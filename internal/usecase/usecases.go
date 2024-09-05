package usecase

import (
	r "spocon-backend/internal/domain/repository"
	s "spocon-backend/internal/domain/service"
)

type Usecases struct {
	GradeUsecase GradeUsecase
	TeamUsecase  TeamUsecase
}

func NewUsecases(r *r.Repositories, s *s.Services) *Usecases {
	return &Usecases{
		GradeUsecase: NewGradeUsecase(r.GradeRepository),
		TeamUsecase:  NewTeamUsecase(s.TeamService, r.TeamRepository),
	}
}

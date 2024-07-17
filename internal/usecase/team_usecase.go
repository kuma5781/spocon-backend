package usecase

import (
	s "spocon-backend/internal/domain/service"
)

type TeamUsecase struct {
	TeamService s.TeamService
}

func (u *TeamUsecase) CreateTeam(input *TeamCreateInput) error {
	// team := u.TeamService.Create(
	// 	"id1",
	// 	"uuid1",
	// 	input.TeamName,
	// 	input.SportId,
	// 	input.GradeId,
	// 	input.IconPath,
	// 	input.Description,
	// 	input.AddressState,
	// 	input.AddressCity,
	// )

	return nil
}

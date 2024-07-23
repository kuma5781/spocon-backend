package usecase

import (
	s "spocon-backend/internal/domain/service"
	"spocon-backend/internal/usecase/dto"
)

type TeamUsecase struct {
	TeamService s.TeamService
}

func (u *TeamUsecase) CreateTeam(input *dto.TeamCreateInput) error {
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

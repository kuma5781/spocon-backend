package usecase

import (
	r "spocon-backend/internal/domain/repository"
	s "spocon-backend/internal/domain/service"
	"spocon-backend/internal/usecase/dto"

	"github.com/pkg/errors"
)

type TeamUsecase struct {
	TeamService    s.TeamService
	TeamRepository r.TeamRepository
}

func (u *TeamUsecase) CreateTeam(input *dto.TeamCreateInput) error {
	team, err := u.TeamService.Create(
		input.TeamName,
		input.SportId,
		input.GradeId,
		input.IconPath,
		input.Description,
		input.AddressState,
		input.AddressCity,
	)
	if err != nil {
		return errors.Wrap(err, "チームの生成に失敗しました。")
	}

	if err := u.TeamRepository.Create(team); err != nil {
		return errors.Wrap(err, "チームの登録に失敗しました。")
	}

	return nil
}

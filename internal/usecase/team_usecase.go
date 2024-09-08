package usecase

import (
	t "spocon-backend/internal/domain/model/team"
	r "spocon-backend/internal/domain/repository"
	s "spocon-backend/internal/domain/service"
	"spocon-backend/internal/usecase/dto"

	"github.com/pkg/errors"
)

type TeamUsecase interface {
	CreateTeam(input *dto.TeamCreateInput) (*t.Team, error)
}

func NewTeamUsecase(ts s.TeamService, tr r.TeamRepository) TeamUsecase {
	return &TeamUsecaseImpl{
		TeamService:    ts,
		TeamRepository: tr,
	}
}

type TeamUsecaseImpl struct {
	TeamService    s.TeamService
	TeamRepository r.TeamRepository
}

func (u *TeamUsecaseImpl) CreateTeam(input *dto.TeamCreateInput) (*t.Team, error) {
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
		return nil, errors.Wrap(err, "チームのモデル生成に失敗しました。")
	}

	result, err := u.TeamRepository.Create(team)
	if err != nil {
		return nil, errors.Wrap(err, "チームの登録に失敗しました。")
	}

	return result, nil
}

package service

import (
	g "spocon-backend/internal/domain/model/grade"
	s "spocon-backend/internal/domain/model/sport"
	t "spocon-backend/internal/domain/model/team"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type TeamService struct {
}

func (s *TeamService) Create(
	name string,
	sportId s.SportId,
	gradeId g.GradeId,
	iconPath string,
	description string,
	addressState string,
	addressCity string,
) (*t.Team, error) {
	teamUuid, err := uuid.NewV7()
	if err != nil {
		return nil, errors.Wrap(err, "UUIDの生成に失敗しました。")
	}
	uuid, err := uuid.NewV7()
	if err != nil {
		return nil, errors.Wrap(err, "UUIDの生成に失敗しました。")
	}
	return &t.Team{
		Id:           t.NewTeamId(teamUuid.String()),
		Uuid:         uuid.String(), // このUUIDは何に使うのか確認
		Name:         name,
		SportId:      sportId,
		GradeId:      gradeId,
		IconPath:     iconPath,
		Description:  description,
		AddressState: addressState,
		AddressCity:  addressCity,
	}, nil
}

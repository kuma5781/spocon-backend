package service

import (
	g "spocon-backend/internal/domain/model/grade"
	s "spocon-backend/internal/domain/model/sport"
	t "spocon-backend/internal/domain/model/team"
	"spocon-backend/internal/util"

	"github.com/pkg/errors"
)

type TeamService struct {
	UUIDv7Generator util.UUIDv7Generator
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
	teamUuid, err := s.UUIDv7Generator.Generate()
	if err != nil {
		return nil, errors.Wrap(err, "UUIDの生成に失敗しました。")
	}
	uuid, err := s.UUIDv7Generator.Generate()
	if err != nil {
		return nil, errors.Wrap(err, "UUIDの生成に失敗しました。")
	}
	return &t.Team{
		Id:           t.NewTeamId(teamUuid),
		Uuid:         uuid, // このUUIDは何に使うのか確認
		Name:         name,
		SportId:      sportId,
		GradeId:      gradeId,
		IconPath:     iconPath,
		Description:  description,
		AddressState: addressState,
		AddressCity:  addressCity,
	}, nil
}

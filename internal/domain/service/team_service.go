package service

import (
	g "spocon-backend/internal/domain/model/grade"
	t "spocon-backend/internal/domain/model/team"
)

type TeamService struct {
}

func (s *TeamService) Create(
	id t.TeamId,
	uuid string,
	name string,
	sportId string,
	gradeId g.GradeId,
	iconPath string,
	description string,
	addressState string,
	addressCity string,
) *t.Team {
	return &t.Team{
		Id:           id,
		Uuid:         uuid,
		Name:         name,
		SportId:      sportId,
		GradeId:      gradeId,
		IconPath:     iconPath,
		Description:  description,
		AddressState: addressState,
		AddressCity:  addressCity,
	}
}

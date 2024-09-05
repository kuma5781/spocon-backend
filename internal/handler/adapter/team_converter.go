package adapter

import (
	g "spocon-backend/internal/domain/model/grade"
	s "spocon-backend/internal/domain/model/sport"
	t "spocon-backend/internal/domain/model/team"
	"spocon-backend/internal/openapi"
	"spocon-backend/internal/usecase/dto"
)

func ToCreateTeamInput(req *openapi.CreateTeamRequest) *dto.TeamCreateInput {
	return &dto.TeamCreateInput{
		TeamName:     req.Name,
		SportId:      s.NewSportId(req.SportId),
		GradeId:      g.NewGradeId(req.GradeId),
		IconPath:     req.IconPath,
		Description:  req.Description,
		AddressCity:  req.AddressCity,
		AddressState: req.AddressState,
	}
}

func ToCreateTeamOutput(team *t.Team) *openapi.CreateTeamResponse {
	return &openapi.CreateTeamResponse{
		Id:           team.Id.Value(),
		Uuid:         team.Uuid,
		Name:         team.Name,
		SportId:      team.SportId.Value(),
		GradeId:      team.GradeId.Value(),
		IconPath:     team.IconPath,
		Description:  team.Description,
		AddressState: team.AddressState,
		AddressCity:  team.AddressCity,
	}
}

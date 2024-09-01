package adapter

import (
	g "spocon-backend/internal/domain/model/grade"
	s "spocon-backend/internal/domain/model/sport"
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

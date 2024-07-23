package converter

import (
	"spocon-backend/internal/openapi"
	"spocon-backend/internal/usecase/dto"
)

func ToCreateTeamInput(req *openapi.CreateTeamRequest) *dto.TeamCreateInput {
	return &dto.TeamCreateInput{
		TeamName:     req.Name,
		SportId:      req.SportId,
		GradeId:      req.GradeId,
		IconPath:     req.IconPath,
		Description:  req.Description,
		AddressCity:  req.AddressCity,
		AddressState: req.AddressState,
	}
}

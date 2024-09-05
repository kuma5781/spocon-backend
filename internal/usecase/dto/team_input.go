package dto

import (
	g "spocon-backend/internal/domain/model/grade"
	s "spocon-backend/internal/domain/model/sport"
)

type TeamCreateInput struct {
	TeamName     string
	SportId      s.SportId
	GradeId      g.GradeId
	IconPath     string
	Description  string
	AddressCity  string
	AddressState string
}

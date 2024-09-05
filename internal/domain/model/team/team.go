package model

import (
	g "spocon-backend/internal/domain/model/grade"
	s "spocon-backend/internal/domain/model/sport"
)

type Team struct {
	Id           TeamId
	Uuid         string
	Name         string
	SportId      s.SportId
	GradeId      g.GradeId
	IconPath     string
	Description  string
	AddressState string
	AddressCity  string
}

package team

import (
	g "spocon-backend/internal/domain/model/grade"
)

type TeamEntity struct {
	Id           TeamId
	Uuid         string
	Name         string
	SportId      string
	GradeId      g.GradeId
	IconPath     string
	Description  string
	AddressState string
	AddressCity  string
	CreateAt     string
	UpdateAt     string
}

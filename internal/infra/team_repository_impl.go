package infra

import (
	"database/sql"
	t "spocon-backend/internal/domain/model/team"
)

type TeamRepositoryImpl struct {
	DB *sql.DB
}

func NewTeamRepositoryImpl(db *sql.DB) *TeamRepositoryImpl {
	return &TeamRepositoryImpl{DB: db}
}

func (ri *TeamRepositoryImpl) Create(team *t.Team) error {
	return nil
}

package infra

import (
	"database/sql"
	g "spocon-backend/internal/domain/model/grade"
	s "spocon-backend/internal/domain/model/sport"
	t "spocon-backend/internal/domain/model/team"

	"github.com/pkg/errors"
)

type TeamRepositoryImpl struct {
	DB *sql.DB
}

func NewTeamRepositoryImpl(db *sql.DB) *TeamRepositoryImpl {
	return &TeamRepositoryImpl{DB: db}
}

func (ri *TeamRepositoryImpl) Create(team *t.Team) (*t.Team, error) {
	_, err := ri.DB.Exec(
		"INSERT INTO teams (id, uuid, name, sport_id, grade_id, icon_path, description, address_state, address_city) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		team.Id.Value(),
		team.Uuid,
		team.Name,
		team.SportId.Value(),
		team.GradeId.Value(),
		team.IconPath,
		team.Description,
		team.AddressState,
		team.AddressCity,
	)
	if err != nil {
		return nil, errors.Wrap(err, "teamsテーブルへのインサートに失敗しました。")
	}

	var lastInsertTeam t.Team
	var teamId string
	var sportId string
	var gradeId string

	err = ri.DB.QueryRow("SELECT id, uuid, name, sport_id, grade_id, icon_path, description, address_state, address_city FROM teams WHERE id = ?", team.Id.Value()).
		Scan(&teamId, &lastInsertTeam.Uuid, &lastInsertTeam.Name, &sportId, &gradeId, &lastInsertTeam.IconPath, &lastInsertTeam.Description, &lastInsertTeam.AddressState, &lastInsertTeam.AddressCity)
	if err != nil {
		return nil, errors.Wrap(err, "teamsテーブルにインサートしたレコード取得に失敗しました。")
	}

	lastInsertTeam.Id = t.NewTeamId(teamId)
	lastInsertTeam.SportId = s.NewSportId(sportId)
	lastInsertTeam.GradeId = g.NewGradeId(gradeId)

	return &lastInsertTeam, nil
}

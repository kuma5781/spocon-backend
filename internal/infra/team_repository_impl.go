package infra

import (
	"database/sql"
	t "spocon-backend/internal/domain/model/team"

	"github.com/pkg/errors"
)

type TeamRepositoryImpl struct {
	DB *sql.DB
}

func NewTeamRepositoryImpl(db *sql.DB) *TeamRepositoryImpl {
	return &TeamRepositoryImpl{DB: db}
}

func (ri *TeamRepositoryImpl) Create(team *t.Team) error {
	_, err := ri.DB.Exec("INSERT INTO teams (id, uuid, name, sport_id, grade_id, icon_path, description, address_state, address_city) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", team.Id.Value(), team.Uuid, team.Name, team.SportId.Value(), team.GradeId.Value(), team.IconPath, team.Description, team.AddressState, team.AddressCity)
	if err != nil {
		return errors.Wrap(err, "teamテーブルのレコード作成に失敗しました。")
	}
	return nil
}

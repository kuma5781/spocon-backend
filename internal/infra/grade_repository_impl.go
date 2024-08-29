package infra

import (
	"database/sql"
	g "spocon-backend/internal/domain/model/grade"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type GradeRepositoryImpl struct {
	DB *sql.DB
}

func NewGradeRepositoryImpl(db *sql.DB) *GradeRepositoryImpl {
	return &GradeRepositoryImpl{DB: db}
}

func (ri *GradeRepositoryImpl) FetchAll() ([]g.Grade, error) {
	rows, err := ri.DB.Query("SELECT id, name FROM grades")
	if err != nil {
		return nil, errors.Wrap(err, "gradesのレコード取得に失敗しました。")
	}
	defer rows.Close()

	var grades []g.Grade
	var id string
	for rows.Next() {
		var grade g.Grade
		if err := rows.Scan(&id, &grade.Name); err != nil {
			return nil, errors.Wrap(err, "gradesレコードの変換処理に失敗しました。")
		}
		grade.Id = g.NewGradeId(id)
		grades = append(grades, grade)
	}

	return grades, nil
}

package infra

import (
	"database/sql"
	g "spocon-backend/internal/domain/model/grade"

	_ "github.com/go-sql-driver/mysql"
)

type GradeRepositoryImpl struct {
	DB *sql.DB
}

func (ri *GradeRepositoryImpl) FetchAll() ([]g.Grade, error) {
	rows, err := ri.DB.Query("SELECT id, name FROM grades")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var grades []g.Grade
	for rows.Next() {
		var grade g.Grade
		if err := rows.Scan(&grade.Id, &grade.Name); err != nil {
			return nil, err
		}
		grades = append(grades, grade)
	}

	return grades, nil
}

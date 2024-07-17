package infra

import (
	"database/sql"
	g "spocon-backend/internal/domain/model/grade"

	_ "github.com/go-sql-driver/mysql"
)

type GradeRepositoryImpl struct {
	DB *sql.DB
}

func (ri *GradeRepositoryImpl) FetchAll() ([]g.GradeEntity, error) {
	rows, err := ri.DB.Query("SELECT * FROM grades")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var grades []g.GradeEntity
	for rows.Next() {
		var grade g.GradeEntity
		if err := rows.Scan(&grade.Id, &grade.Name, &grade.CreatedAt, &grade.UpdatedAt); err != nil {
			return nil, err
		}
		grades = append(grades, grade)
	}

	return grades, nil
}

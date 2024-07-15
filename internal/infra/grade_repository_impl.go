package infra

import (
	"database/sql"
	m "spocon-backend/internal/domain/model"

	_ "github.com/go-sql-driver/mysql"
)

type GradeRepositoryImpl struct {
	DB *sql.DB
}

func (g *GradeRepositoryImpl) FetchAll() ([]m.GradeEntity, error) {
	rows, err := g.DB.Query("SELECT * FROM grades")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var grades []m.GradeEntity
	for rows.Next() {
		var grade m.GradeEntity
		if err := rows.Scan(&grade.Id, &grade.Name, &grade.CreatedAt, &grade.UpdatedAt); err != nil {
			return nil, err
		}
		grades = append(grades, grade)
	}

	return grades, nil
}

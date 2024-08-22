package infra_test

import (
	"sort"
	g "spocon-backend/internal/domain/model/grade"
	model "spocon-backend/internal/domain/model/grade"
	"spocon-backend/internal/infra"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchAll(t *testing.T) {

	tests := []struct {
		name       string
		assertions func(t *testing.T, actual []g.Grade)
	}{
		{
			name: "Test FetchAll",
			assertions: func(t *testing.T, actual []g.Grade) {
				sort.Slice(actual, func(i, j int) bool {
					return actual[i].Id.Value() < actual[j].Id.Value()
				})
				expected := []g.Grade{
					{Id: model.NewGradeId("1"), Name: "小学生"},
					{Id: model.NewGradeId("2"), Name: "中学生"},
					{Id: model.NewGradeId("3"), Name: "高校生"},
					{Id: model.NewGradeId("4"), Name: "大学生"},
					{Id: model.NewGradeId("5"), Name: "社会人"},
				}
				assert.Equal(t, expected, actual)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := infra.NewGradeRepositoryImpl(testDB.DB)
			actual, err := sut.FetchAll()
			if err != nil {
				t.Fatal(err)
			}
			tt.assertions(t, actual)
		})
	}
}

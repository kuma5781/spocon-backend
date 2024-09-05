package infra_test

import (
	g "spocon-backend/internal/domain/model/grade"
	sp "spocon-backend/internal/domain/model/sport"
	te "spocon-backend/internal/domain/model/team"
	"spocon-backend/internal/infra"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {

	sut := infra.NewTeamRepositoryImpl(testDB.DB)

	tests := []struct {
		name       string
		input      *te.Team
		assertions func(t *testing.T, actual *te.Team, expected *te.Team)
	}{
		{
			name: "チーム登録に成功すること",
			input: &te.Team{
				Id:           te.NewTeamId("0191b7d3-a2af-702a-9180-1fff5454a6c9"),
				Uuid:         "0191b7d3-a2af-702a-9180-1fff5454a6c8",
				Name:         "team1",
				SportId:      sp.NewSportId("1"),
				GradeId:      g.NewGradeId("1"),
				IconPath:     "icon1",
				Description:  "description1",
				AddressState: "state1",
				AddressCity:  "city1",
			},
			assertions: func(t *testing.T, actual *te.Team, expected *te.Team) {
				assert.Equal(t, actual, expected)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := testDB.DB.Begin()
			if err != nil {
				t.Fatal(err)
			}
			actual, err := sut.Create(tt.input)
			if err != nil {
				t.Fatal(err) // TODO: 異常系のテストをできる作りにしてテストする
			}
			tt.assertions(t, actual, tt.input)
			tx.Rollback()
		})
	}
}

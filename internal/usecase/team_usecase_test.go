package usecase_test

import (
	"errors"
	g "spocon-backend/internal/domain/model/grade"
	sp "spocon-backend/internal/domain/model/sport"
	te "spocon-backend/internal/domain/model/team"
	rm "spocon-backend/internal/domain/repository/mocks"
	s "spocon-backend/internal/domain/service"
	"spocon-backend/internal/usecase"
	"spocon-backend/internal/usecase/dto"
	um "spocon-backend/internal/util/mocks"

	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateTeam(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockTeamRepository := rm.NewMockTeamRepository(ctrl)
	mockUUIDv7Generator := um.NewMockUUIDv7Generator(ctrl)
	teamService := s.NewTeamService(mockUUIDv7Generator)
	sut := usecase.NewTeamUsecase(teamService, mockTeamRepository)

	tests := []struct {
		name         string
		input        *dto.TeamCreateInput
		expectedCall func()
		assertions   func(t *testing.T, actual *te.Team, err error)
	}{
		{
			name: "チームの登録に成功すること",
			input: &dto.TeamCreateInput{
				TeamName:     "team1",
				SportId:      sp.NewSportId("1"),
				GradeId:      g.NewGradeId("1"),
				IconPath:     "icon1",
				Description:  "description1",
				AddressState: "state1",
				AddressCity:  "city1",
			},
			expectedCall: func() {
				mockUUIDv7Generator.EXPECT().Generate().Return("0191b7d3-a2af-702a-9180-1fff5454a6c9", nil).Times(2)
				expectedTeam := &te.Team{
					Id:           te.NewTeamId("0191b7d3-a2af-702a-9180-1fff5454a6c9"),
					Uuid:         "0191b7d3-a2af-702a-9180-1fff5454a6c9",
					Name:         "team1",
					SportId:      sp.NewSportId("1"),
					GradeId:      g.NewGradeId("1"),
					IconPath:     "icon1",
					Description:  "description1",
					AddressState: "state1",
					AddressCity:  "city1",
				}
				mockTeamRepository.EXPECT().Create(expectedTeam).Return(expectedTeam, nil)
			},
			assertions: func(t *testing.T, actual *te.Team, err error) {
				expected := &te.Team{
					Id:           te.NewTeamId("0191b7d3-a2af-702a-9180-1fff5454a6c9"),
					Uuid:         "0191b7d3-a2af-702a-9180-1fff5454a6c9",
					Name:         "team1",
					SportId:      sp.NewSportId("1"),
					GradeId:      g.NewGradeId("1"),
					IconPath:     "icon1",
					Description:  "description1",
					AddressState: "state1",
					AddressCity:  "city1",
				}
				assert.Equal(t, expected, actual)
				assert.Nil(t, err)

			},
		},
		{
			name: "モデルの生成に失敗したとき、エラーが返ること",
			input: &dto.TeamCreateInput{
				TeamName:     "team1",
				SportId:      sp.NewSportId("1"),
				GradeId:      g.NewGradeId("1"),
				IconPath:     "icon1",
				Description:  "description1",
				AddressState: "state1",
				AddressCity:  "city1",
			},
			expectedCall: func() {
				mockUUIDv7Generator.EXPECT().Generate().Return("", errors.New("error occured.")).Times(1)
			},
			assertions: func(t *testing.T, actual *te.Team, err error) {
				assert.Nil(t, actual)
				assert.Equal(t, "チームのモデル生成に失敗しました。: UUIDの生成に失敗しました。: error occured.", err.Error())
			},
		},
		{
			name: "チームの登録に失敗したとき、エラーが返ること",
			input: &dto.TeamCreateInput{
				TeamName:     "team1",
				SportId:      sp.NewSportId("1"),
				GradeId:      g.NewGradeId("1"),
				IconPath:     "icon1",
				Description:  "description1",
				AddressState: "state1",
				AddressCity:  "city1",
			},
			expectedCall: func() {
				mockUUIDv7Generator.EXPECT().Generate().Return("0191b7d3-a2af-702a-9180-1fff5454a6c9", nil).Times(2)
				expectedTeam := &te.Team{
					Id:           te.NewTeamId("0191b7d3-a2af-702a-9180-1fff5454a6c9"),
					Uuid:         "0191b7d3-a2af-702a-9180-1fff5454a6c9",
					Name:         "team1",
					SportId:      sp.NewSportId("1"),
					GradeId:      g.NewGradeId("1"),
					IconPath:     "icon1",
					Description:  "description1",
					AddressState: "state1",
					AddressCity:  "city1",
				}
				mockTeamRepository.EXPECT().Create(expectedTeam).Return(nil, errors.New("error occured."))
			},
			assertions: func(t *testing.T, actual *te.Team, err error) {
				assert.Nil(t, actual)
				assert.Equal(t, "チームの登録に失敗しました。: error occured.", err.Error())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectedCall()
			actual, err := sut.CreateTeam(tt.input)
			tt.assertions(t, actual, err)
		})
	}
}

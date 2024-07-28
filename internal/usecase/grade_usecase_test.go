package usecase_test

import (
	g "spocon-backend/internal/domain/model/grade"
	"spocon-backend/internal/domain/repository/mocks"
	"spocon-backend/internal/usecase"
	"spocon-backend/internal/usecase/dto"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetGrades(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockGradeRepository := mocks.NewMockGradeRepository(ctrl)
	sut := usecase.NewGradeUsecase(mockGradeRepository)

	tests := []struct {
		name         string
		expectedCall func()
		assertions   func(t *testing.T, actual []dto.GradeOutput)
	}{
		{
			name: "グレードの一覧を正常に取得することができること",
			expectedCall: func() {
				mockGradeRepository.EXPECT().FetchAll().Return([]g.Grade{
					{Id: "1", Name: "grade1"},
					{Id: "2", Name: "grade2"},
				}, nil)
			},
			assertions: func(t *testing.T, actual []dto.GradeOutput) {
				expected := []dto.GradeOutput{
					{Id: "1", Name: "grade1"},
					{Id: "2", Name: "grade2"},
				}
				assert.Equal(t, expected, actual)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectedCall()
			actual, err := sut.GetGrades()
			if err != nil {
				t.Fatal(err)
			}
			tt.assertions(t, actual)
		})
	}
}

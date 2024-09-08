package usecase_test

import (
	g "spocon-backend/internal/domain/model/grade"
	model "spocon-backend/internal/domain/model/grade"
	"spocon-backend/internal/domain/repository/mocks"
	"spocon-backend/internal/usecase"
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
		assertions   func(t *testing.T, actual []g.Grade, err error)
	}{
		{
			name: "グレードの一覧を正常に取得することができること",
			expectedCall: func() {
				mockGradeRepository.EXPECT().FetchAll().Return([]g.Grade{
					{Id: model.NewGradeId("1"), Name: "grade1"},
					{Id: model.NewGradeId("2"), Name: "grade2"},
				}, nil)
			},
			assertions: func(t *testing.T, actual []g.Grade, err error) {
				expected := []g.Grade{
					{Id: model.NewGradeId("1"), Name: "grade1"},
					{Id: model.NewGradeId("2"), Name: "grade2"},
				}
				assert.Nil(t, err)
				assert.Equal(t, expected, actual)
			},
		},
		{
			name: "グレードの一覧を取得できなかった場合、エラーが返却されること",
			expectedCall: func() {
				mockGradeRepository.EXPECT().FetchAll().Return(nil, assert.AnError)
			},
			assertions: func(t *testing.T, actual []g.Grade, err error) {
				assert.Nil(t, actual)
				assert.Equal(t, "gradeのFetchAll()に失敗しました。: "+assert.AnError.Error(), err.Error())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectedCall()
			actual, err := sut.GetGrades()
			tt.assertions(t, actual, err)
		})
	}
}

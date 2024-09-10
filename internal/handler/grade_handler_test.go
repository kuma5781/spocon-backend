package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	g "spocon-backend/internal/domain/model/grade"
	model "spocon-backend/internal/domain/model/grade"
	"spocon-backend/internal/handler"
	"spocon-backend/internal/openapi"
	"spocon-backend/internal/usecase/mocks"
	"spocon-backend/testutil"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetGrade(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGradeUsecase := mocks.NewMockGradeUsecase(ctrl)

	handlers := handler.Handlers{
		GradeHandler: handler.GradeHandler{
			GradeUsecase: mockGradeUsecase,
		},
	}
	e := testutil.EchoInit(&handlers)

	tests := []struct {
		name         string
		expectedCall func()
		assertions   func(t *testing.T, result *http.Response)
	}{
		{
			name: "グレードの一覧を正常に取得することができること",
			expectedCall: func() {
				mockGradeUsecase.EXPECT().GetGrades().Return([]g.Grade{
					{Id: model.NewGradeId("1"), Name: "grade1"},
					{Id: model.NewGradeId("2"), Name: "grade2"},
				}, nil)
			},
			assertions: func(t *testing.T, result *http.Response) {
				assert.Equal(t, http.StatusOK, result.StatusCode)
				defer result.Body.Close()
				var response openapi.GetGradesResponse
				if err := json.NewDecoder(result.Body).Decode(&response); err != nil {
					t.Fatal(err)
				}
				expected := openapi.GetGradesResponse{
					Grades: []openapi.Grade{
						{Id: "1", Name: "grade1"},
						{Id: "2", Name: "grade2"},
					},
				}
				assert.Equal(t, expected, response)
			},
		},
		{
			name: "グレードの一覧を取得できなかった場合、500のステータスコードでレスポンスが返却されること",
			expectedCall: func() {
				mockGradeUsecase.EXPECT().GetGrades().Return(nil, assert.AnError)
			},
			assertions: func(t *testing.T, result *http.Response) {
				assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectedCall()

			req := httptest.NewRequest(http.MethodGet, "/grades", nil)
			res := httptest.NewRecorder()
			e.ServeHTTP(res, req)
			result := res.Result()

			tt.assertions(t, result)
		})
	}

}

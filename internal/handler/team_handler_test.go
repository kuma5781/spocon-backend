package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	g "spocon-backend/internal/domain/model/grade"
	s "spocon-backend/internal/domain/model/sport"
	te "spocon-backend/internal/domain/model/team"

	"spocon-backend/internal/handler"
	"spocon-backend/internal/openapi"
	"spocon-backend/internal/usecase/mocks"
	"spocon-backend/testutil"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateTeam(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTeamUsecase := mocks.NewMockTeamUsecase(ctrl)
	handlers := handler.Handlers{
		TeamHandler: handler.TeamHandler{
			TeamUsecase: mockTeamUsecase,
		},
	}
	e := testutil.EchoInit(&handlers)

	tests := []struct {
		name         string
		body         string
		expectedCall func()
		assertions   func(t *testing.T, result *http.Response)
	}{
		{
			name: "チームの登録に成功すること",
			body: `{"name":"name_test", "sport_id":"1", "grade_id": "1", "icon_path": "path_test", "description": "description_test", "address_state": "address_state_test", "address_city": "address_city_test"}`,
			expectedCall: func() {
				mockTeamUsecase.EXPECT().CreateTeam(gomock.Any()).Return(&te.Team{
					Id:           te.NewTeamId("1"),
					Uuid:         "uuid_test",
					Name:         "name_test",
					SportId:      s.NewSportId("1"),
					GradeId:      g.NewGradeId("1"),
					IconPath:     "path_test",
					Description:  "description_test",
					AddressState: "address_state_test",
					AddressCity:  "address_city_test",
				}, nil)
			},
			assertions: func(t *testing.T, result *http.Response) {
				defer result.Body.Close()
				var response openapi.CreateTeamResponse
				if err := json.NewDecoder(result.Body).Decode(&response); err != nil {
					t.Fatal(err)
				}
				expected := openapi.CreateTeamResponse{
					Id:           "1",
					Uuid:         "uuid_test",
					Name:         "name_test",
					SportId:      "1",
					GradeId:      "1",
					IconPath:     "path_test",
					Description:  "description_test",
					AddressState: "address_state_test",
					AddressCity:  "address_city_test",
				}

				assert.Equal(t, http.StatusCreated, result.StatusCode)
				assert.Equal(t, expected, response)
			},
		},
		{
			name: "チーム登録に失敗した場合、500のステータスコードでレスポンスが返却されること",
			body: `{"name":"name_test", "sport_id":"1", "grade_id": "1", "icon_path": "path_test", "description": "description_test", "address_state": "address_state_test", "address_city": "address_city_test"}`,
			expectedCall: func() {
				mockTeamUsecase.EXPECT().CreateTeam(gomock.Any()).Return(nil, assert.AnError)
			},
			assertions: func(t *testing.T, result *http.Response) {
				assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectedCall != nil {
				tt.expectedCall()
			}

			req := httptest.NewRequest(http.MethodPost, "/teams", strings.NewReader(tt.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()
			e.ServeHTTP(res, req)
			result := res.Result()

			tt.assertions(t, result)
		})
	}
}

// Code generated by MockGen. DO NOT EDIT.
// Source: team_usecase.go
//
// Generated by this command:
//
//	mockgen -source=team_usecase.go -destination=mocks/mock_team_usecase.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	model "spocon-backend/internal/domain/model/team"
	dto "spocon-backend/internal/usecase/dto"

	gomock "go.uber.org/mock/gomock"
)

// MockTeamUsecase is a mock of TeamUsecase interface.
type MockTeamUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockTeamUsecaseMockRecorder
}

// MockTeamUsecaseMockRecorder is the mock recorder for MockTeamUsecase.
type MockTeamUsecaseMockRecorder struct {
	mock *MockTeamUsecase
}

// NewMockTeamUsecase creates a new mock instance.
func NewMockTeamUsecase(ctrl *gomock.Controller) *MockTeamUsecase {
	mock := &MockTeamUsecase{ctrl: ctrl}
	mock.recorder = &MockTeamUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamUsecase) EXPECT() *MockTeamUsecaseMockRecorder {
	return m.recorder
}

// CreateTeam mocks base method.
func (m *MockTeamUsecase) CreateTeam(input *dto.TeamCreateInput) (*model.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTeam", input)
	ret0, _ := ret[0].(*model.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTeam indicates an expected call of CreateTeam.
func (mr *MockTeamUsecaseMockRecorder) CreateTeam(input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTeam", reflect.TypeOf((*MockTeamUsecase)(nil).CreateTeam), input)
}

// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package openapi

// CreateTeamRequest defines model for CreateTeamRequest.
type CreateTeamRequest struct {
	AddressCity  string `json:"address_city"`
	AddressState string `json:"address_state"`
	Description  string `json:"description"`
	GradeId      string `json:"grade_id"`
	IconPath     string `json:"icon_path"`
	Name         string `json:"name"`
	SportId      string `json:"sport_id"`
}

// GetGradesResponse defines model for GetGradesResponse.
type GetGradesResponse struct {
	Grades []Grade `json:"grades"`
}

// Grade defines model for Grade.
type Grade struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// CreateTeamJSONRequestBody defines body for CreateTeam for application/json ContentType.
type CreateTeamJSONRequestBody = CreateTeamRequest

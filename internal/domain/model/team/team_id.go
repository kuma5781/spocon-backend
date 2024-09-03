package model

type TeamId struct {
	id string
}

func (t TeamId) Value() string {
	return t.id
}

func NewTeamId(id string) TeamId {
	return TeamId{id: id}
}

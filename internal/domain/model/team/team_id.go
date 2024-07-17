package team

type TeamId struct {
	id string
}

func NewTeamId(id string) TeamId {
	return TeamId{id: id}
}

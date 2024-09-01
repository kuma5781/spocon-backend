package model

type SportId struct {
	id string
}

func NewSportId(id string) SportId {
	return SportId{id: id}
}

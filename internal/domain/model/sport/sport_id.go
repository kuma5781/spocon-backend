package model

type SportId struct {
	id string
}

func (s SportId) Value() string {
	return s.id
}

func NewSportId(id string) SportId {
	return SportId{id: id}
}

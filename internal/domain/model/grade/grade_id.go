package model

type GradeId struct {
	id string
}

func (gid GradeId) Value() string {
	return gid.id
}

func NewGradeId(id string) GradeId {
	return GradeId{id: id}
}

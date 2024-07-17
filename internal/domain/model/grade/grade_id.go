package grade

type GradeId struct {
	id string
}

func NewGradeId(id string) GradeId {
	return GradeId{id: id}
}

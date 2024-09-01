package repository

import (
	t "spocon-backend/internal/domain/model/team"
)

type TeamRepository interface {
	Create(*t.Team) error
}

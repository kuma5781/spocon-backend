//go:generate mockgen -source=$GOFILE -destination=mocks/mock_$GOFILE -package=mocks
package repository

import (
	t "spocon-backend/internal/domain/model/team"
)

type TeamRepository interface {
	Create(*t.Team) (*t.Team, error)
}

package service

import "spocon-backend/internal/util"

type Services struct {
	TeamService TeamService
}

// utilパッケージから他に依存するものが増えれば、utilsのような構造体を作成してutilsを引数に渡すようにする
func NewServices(uuidv7Generator util.UUIDv7Generator) *Services {
	return &Services{
		TeamService: NewTeamService(uuidv7Generator),
	}
}

package app

import (
	"errors"

	e "github.com/pkg/errors"
)

var (
	InvalidParamsError  = errors.New("リクエストに不正な値が含まれています。")
	NotFoundEntityError = errors.New("指定されたリソースが見つかりません。")

	BadRequestErrors = errors.Join(
		InvalidParamsError,
	)

	NotFoundErrors = errors.Join(
		NotFoundEntityError,
	)
)

func getStatusCode(err error) int {
	code := 500
	cause := e.Cause(err)

	switch cause {
	case BadRequestErrors:
		code = 400
	case NotFoundErrors:
		code = 404
	}
	return code
}

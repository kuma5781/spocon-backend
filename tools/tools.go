//go:build tools
// +build tools

package tools

import (
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
	_ "github.com/segmentio/golines"
	_ "go.uber.org/mock/mockgen"
)

package testutil

import (
	"spocon-backend/internal/app"
	"spocon-backend/internal/handler"
	"spocon-backend/internal/openapi"

	"github.com/labstack/echo/v4"
)

func EchoInit(h *handler.Handlers) *echo.Echo {
	e := echo.New()
	app.InitMiddleware(e)
	openapi.RegisterHandlers(e, h)
	return e
}

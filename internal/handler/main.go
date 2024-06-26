package handler

import (
	"net/http"
	"sample/internal/openapi"

	"github.com/labstack/echo/v4"
)

type Server struct{}

func (s *Server) HealthCheck(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

func (s *Server) CreateUser(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

func main() {
	e := echo.New()

	s := &Server{}
	openapi.RegisterHandlers(e, s)

	e.Logger.Fatal(e.Start(":1323"))
}

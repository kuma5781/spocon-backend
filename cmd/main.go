package main

import (
	"sample/internal/handler"
	"sample/internal/openapi"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	s := &handler.Handlers{}

	openapi.RegisterHandlers(e, s)

	e.Logger.Fatal(e.Start(":1323"))
}

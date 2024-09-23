package main

import (
	"os"
	"spocon-backend/internal/app"
	"spocon-backend/internal/domain/service"
	"spocon-backend/internal/handler"
	"spocon-backend/internal/infra"
	"spocon-backend/internal/openapi"
	"spocon-backend/internal/usecase"
	"spocon-backend/internal/util"

	"github.com/labstack/echo/v4"
)

func main() {
	logger := app.InitLogger(os.Stdout, os.Stderr)
	util.InitAppLogger(logger)

	e := echo.New()
	app.InitMiddleware(e)

	db, err := app.InitDB()

	if err != nil {
		e.Logger.Fatal(err)
	}

	r := infra.NewRepositories(db)
	s := service.NewServices(util.NewUUIDv7Generator())
	u := usecase.NewUsecases(r, s)
	h := handler.NewHandlers(u)
	openapi.RegisterHandlers(e, h)

	e.Logger.Fatal(e.Start(":1323"))
}

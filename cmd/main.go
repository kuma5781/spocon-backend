package main

import (
	"spocon-backend/internal/app"
	"spocon-backend/internal/domain/repository"
	"spocon-backend/internal/handler"
	"spocon-backend/internal/openapi"
	"spocon-backend/internal/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	app.InitMiddleware(e)

	db, err := app.InitDB()

	if err != nil {
		e.Logger.Fatal(err)
	}

	r := repository.NewRepositories(db)
	u := usecase.NewUsecases(r)
	h := handler.NewHandlers(u)
	openapi.RegisterHandlers(e, h)

	e.Logger.Fatal(e.Start(":1323"))
}

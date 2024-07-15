package main

import (
	"spocon-backend/internal/app"
	"spocon-backend/internal/openapi"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := app.InitDB()

	if err != nil {
		e.Logger.Fatal(err) // TODO: エラー時のハンドリング
	}

	r := app.NewRepositories(db)
	u := app.NewUsecases(r)
	h := app.NewHandlers(u)
	openapi.RegisterHandlers(e, h)

	e.Logger.Fatal(e.Start(":1323"))
}

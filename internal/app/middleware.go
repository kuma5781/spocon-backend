package app

import (
	"fmt"

	e "spocon-backend/internal/errors"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitMiddleware(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = CustomHttpErrorHandler
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	})) // TODO: 動作確認
	e.Use(middleware.Logger())
}

func CustomHttpErrorHandler(err error, c echo.Context) {
	fmt.Printf("STACK TRACE:\n %+v\n", err)
	if c.Response().Committed {
		return
	}
	c.JSON(e.GetStatusCode(err), map[string]string{"error": err.Error()})
}

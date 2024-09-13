package app

import (
	"fmt"

	e "spocon-backend/internal/errors"
	"spocon-backend/internal/openapi"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	oapiMiddleware "github.com/oapi-codegen/echo-middleware"
)

func InitMiddleware(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = CustomHttpErrorHandler
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	})) // TODO: 動作確認
	e.Use(middleware.Logger())

	// openapi.yamlから作成したバリデーションをミドルウェアに設定
	swager, _ := openapi.GetSwagger()
	opts := &oapiMiddleware.Options{
		ErrorHandler:          OapiErrorHandler,
		SilenceServersWarning: true,
	}
	e.Use(oapiMiddleware.OapiRequestValidatorWithOptions(swager, opts))
}

func CustomHttpErrorHandler(err error, c echo.Context) {
	fmt.Printf("STACK TRACE:\n %+v\n", err)
	if c.Response().Committed {
		return
	}
	err = c.JSON(e.GetStatusCode(err), map[string]string{"error": err.Error()})
	if err != nil {
		c.Logger().Error(err) // TODO: ログ出力はログ出力基盤作成時に再考する
	}
}

func OapiErrorHandler(c echo.Context, err *echo.HTTPError) error {
	fmt.Printf("STACK TRACE:\n %+v\n", err)
	return c.JSON(err.Code, map[string]string{"error": fmt.Sprint(err.Message)})
}

package action

import (
	"github.com/labstack/echo/v4"
	"github.com/ytake/draft/header"
	"github.com/ytake/draft/payload"
	"net/http"
)

// HTTPVndErrorResponse HTTPエラーカスタム
func HTTPVndErrorResponse(c echo.Context) error {
	c.Response().Header().Add(echo.HeaderContentType, header.MIMEApplicationHal)
	code := http.StatusNotFound
	pe := payload.NewError()
	return c.JSON(code, pe.VngErrorFactory(
		"not found.",
		payload.VndRequestURI(c.Request().RequestURI),
		payload.VndAboutURI(c.Request().RequestURI)),
	)
}

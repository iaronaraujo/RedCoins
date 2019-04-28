package routers

import (
	"net/http"

	"github.com/labstack/echo"
)

var App *echo.Echo

func init() {
	App = echo.New()

	App.GET("/", home)
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "RED COINS!")
}

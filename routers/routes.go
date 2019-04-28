package routers

import (
	"net/http"

	"github.com/iaronaraujo/RedCoins/controllers"
	"github.com/labstack/echo"
)

var App *echo.Echo

func init() {
	App = echo.New()

	App.GET("/", home)

	api := App.Group("/api")
	api.POST("/signup", controllers.CreateUser)
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "RED COINS!")
}

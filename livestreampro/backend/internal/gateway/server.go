package gateway

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// NewRouter sets up the Echo router with all routes.
func NewRouter() *echo.Echo {
	e := echo.New()
	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	e.GET("/channels", func(c echo.Context) error {
		channels := []string{"general", "random", "gaming"}
		return c.JSON(http.StatusOK, channels)
	})
	return e
}

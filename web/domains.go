package web

import (
	"github.com/SzymanskiFilip/uptime-monitoring-go/storage"
	"github.com/labstack/echo/v4"
)

func GetRegisteredDomains(c echo.Context) error {
	domains := storage.GetDomains()
	return c.JSON(200, domains)
}

func RegisterDomain(c echo.Context) error {

	return c.JSON(200, nil)
}
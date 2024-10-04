package web

import (
	"fmt"

	"github.com/SzymanskiFilip/uptime-monitoring-go/storage"
	"github.com/SzymanskiFilip/uptime-monitoring-go/worker"
	"github.com/labstack/echo/v4"
)

func GetRegisteredDomains(c echo.Context) error {
	domains := storage.GetDomains()
	return c.JSON(200, domains)
}

type DomainPost struct {
	Url string `json:"url"`
}

func RegisterDomain(c echo.Context) error {

	newDomain := DomainPost{}

	err := c.Bind(&newDomain); if err != nil {
		fmt.Println(err)
		return c.JSON(400, nil)
	}

	status := storage.SaveDomain(newDomain.Url)
	worker.UpdateAddresses()

	if status == 1 {
		return c.JSON(200, nil)
	} else {
		return c.JSON(409, nil)
	}
}


type DomainDelete struct {
	ID string `query:"id"`
}
func DeleteDomain(c echo.Context) error{

	d := DomainDelete{}

	err := c.Bind(&d); if err != nil {
		fmt.Println(err)
		return c.JSON(400, nil)
	}

	status := storage.DeleteDomain(d.ID)
	worker.UpdateAddresses()

	if status {
		return c.JSON(200, nil)
	}
	return c.JSON(400, nil)
}
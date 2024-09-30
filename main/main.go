package main

import (
	"fmt"
	"net/http"

	"github.com/SzymanskiFilip/uptime-monitoring-go/storage"
	"github.com/SzymanskiFilip/uptime-monitoring-go/worker"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	fmt.Println("Program started...")
	
	storage.InitializeDatabase()

	worker.StartPinging()

	e.Logger.Fatal(e.Start(":1323"))
}
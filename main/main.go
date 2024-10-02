package main

import (
	"fmt"

	"github.com/SzymanskiFilip/uptime-monitoring-go/storage"
	"github.com/SzymanskiFilip/uptime-monitoring-go/web"
	"github.com/SzymanskiFilip/uptime-monitoring-go/worker"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	web.RegisterEndpoints(e)

	fmt.Println("Program started...")
	
	storage.InitializeDatabase()

	go worker.StartPinging()

	e.Logger.Fatal(e.Start(":1323"))
}
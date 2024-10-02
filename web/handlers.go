package web

import "github.com/labstack/echo/v4"


func RegisterEndpoints(e *echo.Echo){
	e.GET("/domains", GetRegisteredDomains)
	e.POST("/domains", RegisterDomain)
}
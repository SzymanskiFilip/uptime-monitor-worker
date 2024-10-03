package web

import (
	"github.com/SzymanskiFilip/uptime-monitoring-go/storage"
	"github.com/labstack/echo/v4"
)

//funkcja ma zwracać response time średni z dnia i dzień na stronę główną


type QueryId struct {
	ID string `query:"id"`
}

type ResponseTimeAverageResponse struct {
	ID string `json:"id"`
	Data []storage.ResponseTimeRow `json:"data"`
}

func GetDailyResponseTimeAverage(c echo.Context) error {
	id := QueryId{}

	err := c.Bind(&id); if err != nil {
		return c.JSON(404, nil)
	}

	data, err := storage.GetDailyResponseTimeAverage(id.ID); if err != nil {
		return c.JSON(400, nil)
	}

	response := ResponseTimeAverageResponse{
		ID: id.ID,
		Data: data,
	}

	return c.JSON(200, response)
}
package web

import (
	"log"

	"github.com/SzymanskiFilip/uptime-monitoring-go/storage"
	"github.com/SzymanskiFilip/uptime-monitoring-go/types"
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

func GetDetailedStatistics(c echo.Context) error{
	id := QueryId{}
	err := c.Bind(&id); if err != nil {
		c.JSON(500, nil)
	}

	//domena i id
	urlId, urlValue := storage.GetDomainById(id.ID)

	//responsetime average time last month
	row, err := storage.GetDailyResponseTimeAverage(id.ID)
	if err != nil {
		log.Fatal(err)
	}

	//response time max and minimum
	min, max := storage.GetMaxAndMinRespTime(id.ID)
	
	//response time last 7 and 7-14 days
	prev, prev2 := storage.GetPrevWeeks(id.ID)

	outages := storage.GetOutages(id.ID)

	allStats := storage.GetStatistics(id.ID)

	response := DetailedStatistics{
		ID: urlId,
		URL: urlValue,
		ResponseTimeRowMonth: row,
		MinimumStat: min,
		MaximumStat: max,
		AllStats: allStats,
		ResponseTimePrev7: prev,
		ResponseTimePrev14: prev2,
		Outages: outages,
	}

	return c.JSON(200, response)
}

type DetailedStatistics struct {
	URL string `json:"url"`
	ID string `json:"id"`
	ResponseTimeRowMonth []storage.ResponseTimeRow `json:"response_times"`
	MinimumStat types.StatisticStored `json:"minimum"`
	MaximumStat types.StatisticStored `json:"maximum"`
	AllStats []types.StatisticStored `json:"all"`
	ResponseTimePrev7 []storage.ResponseTimeRow `json:"response_times_7"`
	ResponseTimePrev14 []storage.ResponseTimeRow `json:"response_times_14"`
	Outages []types.StatisticStored `json:"outages"`
}
package types

import (
	"time"

	"github.com/google/uuid"
)

type Statistic struct {
	Id uuid.UUID `json:"id"`
	URL string `json:"url"`
	Headers string `json:"headers"`
	Success bool `json:"success"`
	ResponseTime int64 `json:"response_time"`
	SavedAt time.Time `json:"saved_at"`
}

type StatisticStored struct {
	Id uuid.UUID `json:"id"`
	URL_ID string `json:"url_id"`
	Headers string `json:"headers"`
	Success bool `json:"success"`
	ResponseTime int64 `json:"response_time"`
	SavedAt time.Time `json:"saved_at"`
}


type URLStored struct {
	Id string `json:"id"`
	Domain string `json:"url"`
}
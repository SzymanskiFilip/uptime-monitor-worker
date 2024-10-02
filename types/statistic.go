package types

import "github.com/google/uuid"

type Statistic struct {
	Id uuid.UUID `json:"id"`
	URL string `json:"url"`
	Headers string `json:"headers"`
	Status int16 `json:"status"`
	Success bool `json:"success"`
	ResponseTime int64 `json:"response_time"`
}


type URLStored struct {
	Id string `json:"id"`
	Domain string `json:"url"`
}
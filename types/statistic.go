package types

import "github.com/google/uuid"

type Statistic struct {
	id uuid.UUID `json:"id"`
	url string `json:"url"`
}

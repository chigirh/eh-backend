package entities

type (
	AggregateSchedule struct {
		Date   string `json:"date"`
		Period int    `json:"period"`
		Count  int    `json:"count"`
	}
)

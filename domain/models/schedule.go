package models

import (
	"time"
)

type (
	DailyAggregate struct {
		Date    time.Time
		Periods []Period
	}

	Period struct {
		Period int
		Count  int
	}

	Schedule struct {
		Date   time.Time
		Period int
	}

	PeriodDetail struct {
		Period     int
		HourFrom   int
		MinuteFrom int
		HourTo     int
		MinuteTo   int
	}
)

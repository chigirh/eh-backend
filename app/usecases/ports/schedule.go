package ports

import (
	"context"
	"eh-backend-api/domain/models"
	"time"
)

type ScheduleInputPort interface {
	AddSchedule(ctx context.Context, userName models.UserName, schedules []models.Schedule) error
	Aggregate(ctx context.Context, from time.Time, to time.Time) ([]*models.DailyAggregate, error)
}

type ScheduleRepository interface {
	Add(ctx context.Context, userName models.UserName, schedules []models.Schedule) error
	FetchByDays(ctx context.Context, from time.Time, to time.Time) ([]*models.DailyAggregate, error)
}

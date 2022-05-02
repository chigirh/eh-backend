package ports

import (
	"context"
	"eh-backend-api/domain/models"
	"time"
)

type ScheduleInputPort interface {
	AddSchedule(ctx context.Context, userName models.UserName, schedules []models.Schedule) error
	Aggregate(ctx context.Context, from time.Time, to time.Time) ([]*models.DailyAggregate, error)
	GetByPeriod(ctx context.Context, date time.Time, period int) ([]*models.User, error)
	GetPeriodDetail(ctx context.Context) ([]*models.PeriodDetail, error)
}

type ScheduleRepository interface {
	Add(ctx context.Context, userName models.UserName, schedules []models.Schedule) error
	FetchByDays(ctx context.Context, from time.Time, to time.Time) ([]*models.DailyAggregate, error)
	FetchByPeriod(ctx context.Context, date time.Time, period int) ([]*models.User, error)
	FetchPeriodDetail(ctx context.Context) ([]*models.PeriodDetail, error)
}

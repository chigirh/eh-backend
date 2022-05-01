package interactors

import (
	"context"
	"eh-backend-api/app/usecases/ports"
	"eh-backend-api/domain/models"
	"time"
)

type ScheduleInteractor struct {
	ScheduleRepository ports.ScheduleRepository
}

func (it *ScheduleInteractor) AddSchedule(
	ctx context.Context,
	userName models.UserName,
	schedules []models.Schedule,
) error {
	if err := it.ScheduleRepository.Add(ctx, userName, schedules); err != nil {
		return err
	}

	return nil
}

func (it *ScheduleInteractor) Aggregate(
	ctx context.Context,
	from time.Time,
	to time.Time,
) ([]*models.DailyAggregate, error) {
	aggregates, err := it.ScheduleRepository.FetchByDays(ctx, from, to)
	if err != nil {
		return nil, err
	}

	return aggregates, nil
}

// di
func NewScheduleInputPort(shceduleRepository ports.ScheduleRepository) ports.ScheduleInputPort {
	return &ScheduleInteractor{
		ScheduleRepository: shceduleRepository,
	}
}

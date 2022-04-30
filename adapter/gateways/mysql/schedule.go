package mysql

import (
	"context"
	"eh-backend-api/adapter/gateways/entities"
	"eh-backend-api/app/usecases/ports"
	"eh-backend-api/domain/models"
	"time"

	"github.com/google/uuid"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	gormbulk "github.com/t-tiger/gorm-bulk-insert/v2"
)

type ScheduleGateway struct {
}

func (it *ScheduleGateway) Add(
	ctx context.Context,
	userName models.UserName,
	schedules []models.Schedule,
) error {

	db, err := NewDbConnection()
	if err != nil {
		return err
	}

	tx := db.Begin()

	ids := []string{}
	var scEntities []interface{}

	for i := 0; i < len(schedules); i++ {

		sc := schedules[i]
		rt := []*entities.Schedule{}
		if err := tx.Where(
			"user_id = ? AND date = ? AND period = ?",
			userName, toSqlString(sc.Date),
			sc.Period,
		).Find(&rt).Error; err != nil {
			return err
		}

		if 0 < len(rt) {
			continue
		}

		uuid, _ := uuid.NewRandom()
		scId := uuid.String()
		scEntities = append(
			scEntities,
			entities.Schedule{
				ScheduleId: string(scId),
				UserId:     string(userName),
				Date:       toSqlString(sc.Date),
				Period:     sc.Period,
			},
		)
		ids = append(ids, scId)
	}

	if err := gormbulk.BulkInsert(tx, scEntities, 100); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (it *ScheduleGateway) FetchByDays(
	ctx context.Context,
	from time.Time,
	to time.Time,
) ([]*models.DailyAggregate, error) {

	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}

	result := []entities.AggregateSchedule{}
	err = db.Debug().Model(&entities.Schedule{}).
		Select("`date`, `period`, count(`user_id`) AS `count`").
		Table("schedules").
		Where("? <= `date` AND `date` <= ?", toSqlString(from), toSqlString(to)).
		Group("`date`, `period`").Find(&result).Error

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil
	}

	m := map[string]*models.DailyAggregate{}
	model := []*models.DailyAggregate{}

	for i := 0; i < len(result); i++ {
		r := result[i]
		d := r.Date
		_, has := m[d]
		if !has {
			agg := models.DailyAggregate{Date: ToDate(d)}
			model = append(model, &agg)
			m[d] = &agg
		}

		agg, _ := m[d]
		agg.Periods = append(agg.Periods, models.Period{Period: r.Period, Count: r.Count})
	}
	return model, nil
}

// di
func NewScheduleRepository() ports.ScheduleRepository {
	return &ScheduleGateway{}
}

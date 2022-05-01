package schedule

import (
	"context"
	"net/http"

	"github.com/labstack/echo"

	"eh-backend-api/adapter/controllers"
	"eh-backend-api/app/usecases/ports"
	"eh-backend-api/domain/models"
)

type ScheduleApi interface {
	Aggregate(ctx context.Context) func(c echo.Context) error
	Post(ctx context.Context) func(c echo.Context) error
}

type ScheduleController struct {
	requestMapper controllers.RequestMapper
	authPort      ports.AuthInputPort
	inputPort     ports.ScheduleInputPort
}

func (it *ScheduleController) Aggregate(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		req := new(AggregateGetRequest)
		if error := it.requestMapper.Parse(c, req); error != nil {
			return error
		}

		// If session token is set, have general.
		token, err := it.requestMapper.GetSessionToken(c)
		if err != nil {
			return err
		}
		usrl, err := it.authPort.GetUserRole(ctx, token)
		if err != nil {
			return controllers.ErrorHandle(c, err)
		}
		if !usrl.HaveAdmin() && !usrl.HaveCorp() {
			return c.JSON(http.StatusForbidden, controllers.ErrorResponse{Message: "Requires administrator or corporation"})
		}

		// usecase
		var aggregates []*models.DailyAggregate
		aggregates, err = it.inputPort.Aggregate(ctx, controllers.ToDate(req.From), controllers.ToDate(req.To))

		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		resAggs := []*DailyAggregateDto{}
		for i := 0; i < len(aggregates); i++ {
			aggregate := aggregates[i]
			resAgg := &DailyAggregateDto{Date: controllers.ToString(aggregate.Date)}
			resAggs = append(resAggs, resAgg)
			for j := 0; j < len(aggregate.Periods); j++ {
				period := aggregate.Periods[j]
				resAgg.Periods = append(resAgg.Periods, &PeriodDto{Period: period.Period, Count: period.Count})
			}
		}

		res := AggregateGetResponse{Aggregates: resAggs}

		return c.JSON(http.StatusOK, res)
	}
}

func (it *ScheduleController) Post(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		req := new(PostRequest)
		if error := it.requestMapper.Parse(c, req); error != nil {
			return error
		}

		// If session token is set, have general.
		token, err := it.requestMapper.GetSessionToken(c)
		if err != nil {
			return err
		}
		usrl, err := it.authPort.GetUserRole(ctx, token)
		if err != nil {
			return controllers.ErrorHandle(c, err)
		}
		if !usrl.HaveGene() {
			return c.JSON(http.StatusForbidden, controllers.ErrorResponse{Message: "Requires general"})
		}

		schedules := []models.Schedule{}

		for i := 0; i < len(req.Schedules); i++ {
			reqSc := req.Schedules[i]
			for j := 0; j < len(reqSc.Periods); j++ {
				schedules = append(schedules, models.Schedule{
					Date:   controllers.ToDate(reqSc.Date),
					Period: reqSc.Periods[j],
				})
			}
		}

		if err = it.inputPort.AddSchedule(ctx, usrl.UserName, schedules); err != nil {
			return controllers.ErrorHandle(c, err)
		}

		return c.JSON(http.StatusOK, controllers.DefaultResponse)
	}
}

// dto
type (
	AggregateGetRequest struct {
		From string `query:"from" validate:"required,date"`
		To   string `query:"to" validate:"required,date"`
	}

	AggregateGetResponse struct {
		Aggregates []*DailyAggregateDto `json:"aggregates"`
	}

	DailyAggregateDto struct {
		Date    string       `json:"date"`
		Periods []*PeriodDto `json:"periods"`
	}

	PeriodDto struct {
		Period int `json:"period"`
		Count  int `json:"count"`
	}

	PostRequest struct {
		Schedules []DaylySchedule `json:"schedules" validate:"min=1,dive"`
	}

	DaylySchedule struct {
		Date    string `json:"date" validate:"required,date"`
		Periods []int  `json:"periods" validate:"min=1,max=48,unique,dive,min=1,max=48"`
	}
)

// di
func NewScheduleController(
	requestMapper controllers.RequestMapper,
	authPort ports.AuthInputPort,
	inputPort ports.ScheduleInputPort) ScheduleApi {
	return &ScheduleController{
		requestMapper: requestMapper,
		authPort:      authPort,
		inputPort:     inputPort,
	}
}

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
	AggregateGet(ctx context.Context) func(c echo.Context) error
	DetailsGet(ctx context.Context) func(c echo.Context) error
	PeriodsGet(ctx context.Context) func(c echo.Context) error
	Post(ctx context.Context) func(c echo.Context) error
}

type ScheduleController struct {
	requestMapper controllers.RequestMapper
	authPort      ports.AuthInputPort
	inputPort     ports.ScheduleInputPort
}

func (it *ScheduleController) AggregateGet(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		req := new(AggregateGetRequest)
		if error := it.requestMapper.Parse(c, req); error != nil {
			return error
		}

		// If session token is set, have corporation administrator or corporation.
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

func (it *ScheduleController) DetailsGet(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		req := new(DetailsGetRequest)
		if error := it.requestMapper.Parse(c, req); error != nil {
			return error
		}

		// If session token is set, have general.have corporation administrator or corporation.
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
		var users []*models.User
		users, err = it.inputPort.GetByPeriod(ctx, controllers.ToDate(req.Date), req.Period)

		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		details := []*UserDto{}
		for i := 0; i < len(users); i++ {
			u := users[i]
			details = append(details, &UserDto{
				UserName:   string(u.UserId),
				FirstName:  u.Firstname,
				FamilyName: u.FamilyName,
			})
		}

		res := DetailsGetResponse{Details: details}

		return c.JSON(http.StatusOK, res)
	}
}

func (it *ScheduleController) PeriodsGet(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		// If session token is set, have any.
		token, err := it.requestMapper.GetSessionToken(c)
		if err != nil {
			return err
		}
		_, err = it.authPort.GetUserRole(ctx, token)
		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		// usecase
		var peripds []*models.PeriodDetail
		peripds, err = it.inputPort.GetPeriodDetail(ctx)

		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		resPeriods := []*PeriodsDto{}
		for i := 0; i < len(peripds); i++ {
			p := peripds[i]
			resPeriods = append(resPeriods, &PeriodsDto{
				Period:     p.Period,
				HourFrom:   p.HourFrom,
				MinuteFrom: p.MinuteFrom,
				HourTo:     p.HourTo,
				MinuteTo:   p.MinuteTo,
			})
		}

		res := PeriodsGetResponse{Periods: resPeriods}

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

	DetailsGetRequest struct {
		Date   string `query:"date" validate:"required,date"`
		Period int    `query:"period" validate:"required,min=1,max=48"`
	}

	DetailsGetResponse struct {
		Details []*UserDto `json:"details"`
	}

	UserDto struct {
		UserName   string `json:"user_name"`
		FirstName  string `json:"first_name"`
		FamilyName string `json:"family_name"`
	}

	PeriodsGetResponse struct {
		Periods []*PeriodsDto `json:"periods"`
	}

	PeriodsDto struct {
		Period     int `json:"period"`
		HourFrom   int `json:"hour_from"`
		MinuteFrom int `json:"minute_from"`
		HourTo     int `json:"hour_to"`
		MinuteTo   int `json:"minute_to"`
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

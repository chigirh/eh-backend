package entities

type (
	User struct {
		UserId     string `json:"user_id"`
		FirstName  string `json:"first_name"`
		FamilyName string `json:"family_name"`
	}

	Password struct {
		UserId   string `json:"user_id"`
		Password string `json:"password"`
	}

	Role struct {
		UserId string `json:"user_id"`
		Role   string `json:"role"`
	}

	Schedule struct {
		ScheduleId string `json:"schedule_id"`
		UserId     string `json:"user_id"`
		Date       string `json:"date"`
		Period     int    `json:"period"`
	}

	MasterSchedule struct {
		Period     int `json:"period"`
		HourFrom   int `json:"hour_from"`
		MinuteFrom int `json:"minute_from"`
		HourTo     int `json:"hour_to"`
		MinuteTo   int `json:"minute_to"`
	}
)

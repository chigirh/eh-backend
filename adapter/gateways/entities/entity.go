package entities

type User struct {
	UserId     string `json:"user_id"`
	FirstName  string `json:"first_name"`
	FamilyName string `json:"family_name"`
}

type Password struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

type Role struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`
}

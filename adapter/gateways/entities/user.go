package entities

type User struct {
	UserId     string `json:"user_id"`
	FirstName  string `json:"first_name"`
	FamilyName string `json:"family_name"`
}
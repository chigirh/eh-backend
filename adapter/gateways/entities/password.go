package entities

type Password struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

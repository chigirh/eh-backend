package models

type User struct {
	UserId     string
	Firstname  string
	FamilyName string
}

func (p *User) Set(
	userId string,
	firstname string,
	familyName string,
) {
	p.UserId = userId
	p.Firstname = firstname
	p.FamilyName = familyName
}

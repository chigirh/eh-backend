package models

type User struct {
	UserId     UserName
	Firstname  string
	FamilyName string
	Password   Password
	Roles      []Role
}

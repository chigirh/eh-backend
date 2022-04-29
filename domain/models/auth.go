package models

type UserName string

type Password string

type SessionToken string

type Role string

var RoleAadmin = Role("ADMIN")
var RoleCorp = Role("CORP")
var RoleGene = Role("GENE")

type UserRole struct {
	UserName UserName
	Roles    []Role
}

func (r *UserRole) HaveAdmin() bool {
	return r.have(RoleAadmin)
}

func (r *UserRole) HaveCorp() bool {
	return r.have(RoleCorp)
}

func (r *UserRole) HaveGene() bool {
	return r.have(RoleGene)
}

func (r *UserRole) have(role Role) bool {
	for i := 0; i < len(r.Roles); i++ {
		if r.Roles[i] == role {
			return true
		}
	}
	return false
}

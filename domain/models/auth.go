package models

type (
	UserName     string
	Password     string
	SessionToken string
	Role         string
)

const (
	RoleAadmin = Role("ADMIN")
	RoleCorp   = Role("CORP")
	RoleGene   = Role("GENE")
)

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

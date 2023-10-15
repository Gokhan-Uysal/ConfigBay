package valueobject

var (
	Admin Role = NewRole(
		"Admin",
		ReadProject,
		ManageGroups,
		ManageUsers,
		ReadSecrets,
		WriteSecrets,
		DeleteSecrets,
	)

	TeamLead Role = NewRole(
		"Team Leader",
		ManageGroups,
		ManageUsers,
	)

	ProductOwner Role = NewRole(
		"Product Owner",
		ReadSecrets,
		WriteSecrets,
		DeleteSecrets,
	)

	Developer Role = NewRole(
		"Developer",
		ReadProject,
		ReadSecrets,
	)
)

type (
	Role interface {
		Name() string
		Permissions() []Permission
		AddPermissions(...Permission)
	}

	role struct {
		name        string
		permissions []Permission
	}
)

func NewRole(
	name string,
	permissions ...Permission,
) Role {
	return &role{
		name:        name,
		permissions: permissions,
	}
}

func (r *role) Name() string {
	return r.name
}

func (r *role) Permissions() []Permission {
	return r.permissions
}

func (r *role) AddPermissions(p ...Permission) {
	r.permissions = append(r.permissions, p...)
}

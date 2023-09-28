package aggregate

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/model"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
	"slices"
	"time"
)

type (
	GroupBuilder interface {
		Role(role valueobject.Role) GroupBuilder
		Users(...valueobject.UserID) GroupBuilder
		CreatedAt(time.Time) GroupBuilder
		UpdatedAt(time.Time) GroupBuilder
		model.Builder[Group]
	}

	Group interface {
		Id() valueobject.GroupID
		Title() string
		RolePermissions() []valueobject.Permission
		RoleName() string
		Users() []valueobject.UserID
		ProjectId() valueobject.ProjectID
		CreatedAt() time.Time
		UpdatedAt() time.Time
		SetRole(role valueobject.Role)
		AddUsers(...valueobject.UserID)
		HasPermission(valueobject.Permission) bool
	}

	groupBuilder struct {
		group
	}

	group struct {
		id        valueobject.GroupID
		title     string
		role      valueobject.Role
		users     []valueobject.UserID
		projectId valueobject.ProjectID
		createdAt time.Time
		updatedAt time.Time
	}
)

func NewGroupBuilder(
	id valueobject.GroupID,
	title string,
	projectId valueobject.ProjectID,
) GroupBuilder {
	return &groupBuilder{group{id: id, title: title, projectId: projectId}}
}

func (gb *groupBuilder) Users(users ...valueobject.UserID) GroupBuilder {
	gb.users = users
	return gb
}

func (gb *groupBuilder) Role(role valueobject.Role) GroupBuilder {
	gb.role = role
	return gb
}

func (gb *groupBuilder) CreatedAt(time time.Time) GroupBuilder {
	gb.createdAt = time
	return gb
}

func (gb *groupBuilder) UpdatedAt(time time.Time) GroupBuilder {
	gb.updatedAt = time
	return gb
}

func (gb *groupBuilder) Build() Group {
	return &group{
		id:        gb.id,
		title:     gb.title,
		role:      gb.role,
		users:     gb.users,
		projectId: gb.projectId,
		createdAt: gb.createdAt,
		updatedAt: gb.updatedAt,
	}
}

func (g *group) Id() valueobject.GroupID {
	return g.id
}

func (g *group) Title() string {
	return g.title
}

func (g *group) Users() []valueobject.UserID {
	return g.users
}

func (g *group) RoleName() string {
	return g.role.Name()
}

func (g *group) RolePermissions() []valueobject.Permission {
	return g.role.Permissions()
}

func (g *group) ProjectId() valueobject.ProjectID {
	return g.projectId
}

func (g *group) SetRole(role valueobject.Role) {
	g.role = role
}

func (g *group) AddUsers(users ...valueobject.UserID) {
	g.users = append(g.users, users...)
}

func (g *group) HasPermission(p valueobject.Permission) bool {
	return slices.Contains(g.role.Permissions(), p)
}

func (g *group) CreatedAt() time.Time {
	return g.createdAt
}

func (g *group) UpdatedAt() time.Time {
	return g.updatedAt
}

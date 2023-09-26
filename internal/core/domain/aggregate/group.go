package aggregate

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/model"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/entity"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
	"slices"
	"time"
)

type (
	GroupBuilder interface {
		Users(...valueobject.UserID) GroupBuilder
		Roles(...entity.Role) GroupBuilder
		CreatedAt(time.Time) GroupBuilder
		UpdatedAt(time.Time) GroupBuilder
		model.Builder[Group]
	}

	Group interface {
		Id() valueobject.GroupID
		Title() string
		Users() []valueobject.UserID
		Roles() []entity.Role
		AddRoles(...entity.Role)
		AddUsers(...valueobject.UserID)
		HasRole(entity.Role) bool
		CreatedAt() time.Time
		UpdatedAt() time.Time
	}

	groupBuilder struct {
		group
	}

	group struct {
		id        valueobject.GroupID
		title     string
		users     []valueobject.UserID
		roles     []entity.Role
		createdAt time.Time
		updatedAt time.Time
	}
)

func NewGroupBuilder(id valueobject.GroupID, title string) GroupBuilder {
	return &groupBuilder{group{id: id, title: title}}
}

func (gb *groupBuilder) Users(users ...valueobject.UserID) GroupBuilder {
	gb.users = users
	return gb
}

func (gb *groupBuilder) Roles(roles ...entity.Role) GroupBuilder {
	gb.roles = roles
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
		users:     gb.users,
		roles:     gb.roles,
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

func (g *group) Roles() []entity.Role {
	return g.roles
}

func (g *group) AddUsers(users ...valueobject.UserID) {
	g.users = append(g.users, users...)
}

func (g *group) AddRoles(roles ...entity.Role) {
	g.roles = append(g.roles, roles...)
}

func (g *group) HasRole(role entity.Role) bool {
	return slices.Contains(g.roles, role)
}

func (g *group) CreatedAt() time.Time {
	return g.createdAt
}

func (g *group) UpdatedAt() time.Time {
	return g.updatedAt
}

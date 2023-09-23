package aggregate

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/entity"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
	"time"
)

type (
	GroupBuilder interface {
		Users(...valueobject.ID) GroupBuilder
		Roles(...entity.Role) GroupBuilder
		CreatedAt(time.Time) GroupBuilder
		UpdatedAt(time.Time) GroupBuilder
		domain.Builder[Group]
	}

	Group interface {
		BaseAggregate
		Title() string
		Users() []valueobject.ID
		Roles() []entity.Role
		AddUser(user valueobject.ID)
	}

	groupBuilder struct {
		group
	}

	group struct {
		*baseAggregate
		title string
		users []valueobject.ID
		roles []entity.Role
	}
)

func NewGroupBuilder(id valueobject.ID, title string) GroupBuilder {
	base := newBaseAggregate(id)
	return &groupBuilder{group{baseAggregate: base, title: title}}
}

func (gb *groupBuilder) Users(users ...valueobject.ID) GroupBuilder {
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
		baseAggregate: gb.baseAggregate,
		title:         gb.title,
		users:         gb.users,
		roles:         gb.roles,
	}
}

func (g *group) Title() string {
	return g.title
}

func (g *group) Users() []valueobject.ID {
	return g.users
}

func (g *group) Roles() []entity.Role {
	return g.roles
}

func (g *group) AddUser(user valueobject.ID) {
	g.users = append(g.users, user)
}

func (g *group) AddRole(role entity.Role) {
	g.roles = append(g.roles, role)
}

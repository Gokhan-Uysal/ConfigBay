package aggregate

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/model"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/entity"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
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
		model.BaseAggregate
		Title() string
		Users() []valueobject.UserID
		Roles() []entity.Role
		AddUser(user valueobject.UserID)
	}

	groupBuilder struct {
		group
	}

	group struct {
		*baseAggregate
		title string
		users []valueobject.UserID
		roles []entity.Role
	}
)

func NewGroupBuilder(id valueobject.GroupID, title string) GroupBuilder {
	base := newBaseAggregate(id)
	return &groupBuilder{group{baseAggregate: base, title: title}}
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
		baseAggregate: gb.baseAggregate,
		title:         gb.title,
		users:         gb.users,
		roles:         gb.roles,
	}
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

func (g *group) AddUser(user valueobject.UserID) {
	g.users = append(g.users, user)
}

func (g *group) AddRole(role entity.Role) {
	g.roles = append(g.roles, role)
}

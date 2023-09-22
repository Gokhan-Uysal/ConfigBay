package domain

import "time"

type (
	GroupBuilder interface {
		Users(...User) GroupBuilder
		Roles(...Role) GroupBuilder
		CreatedAt(time.Time) GroupBuilder
		UpdatedAt(time.Time) GroupBuilder
		Builder[Group]
	}

	Group interface {
		Id() ID
		Title() string
		Users() []User
		Roles() []Role
		AddUser(user User)
		Timestamp
	}

	groupBuilder struct {
		group
	}

	group struct {
		id        ID
		title     string
		users     []User
		roles     []Role
		createdAt time.Time
		updatedAt time.Time
	}
)

func NewGroupBuilder(id ID, title string) GroupBuilder {
	return &groupBuilder{group{id: id, title: title}}
}

func (gb *groupBuilder) Users(users ...User) GroupBuilder {
	gb.users = users
	return gb
}

func (gb *groupBuilder) Roles(roles ...Role) GroupBuilder {
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
		roles:     gb.roles,
		createdAt: gb.createdAt,
		updatedAt: gb.updatedAt,
	}
}

func (g *group) Id() ID {
	return g.id
}

func (g *group) Title() string {
	return g.title
}

func (g *group) Users() []User {
	return g.users
}

func (g *group) Roles() []Role {
	return g.roles
}

func (g *group) CreatedAt() time.Time {
	return g.createdAt
}

func (g *group) UpdatedAt() time.Time {
	return g.updatedAt
}

func (g *group) AddUser(user User) {
	g.users = append(g.users, user)
}

func (g *group) AddRole(role Role) {
	g.roles = append(g.roles, role)
}

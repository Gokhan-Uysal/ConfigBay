package domain

import "time"

type (
	// UserBuilder constructs a User entity
	UserBuilder interface {
		Active(bool) UserBuilder
		CreatedAt(time.Time) UserBuilder
		UpdatedAt(time.Time) UserBuilder
		Builder[User]
	}

	// User represents a user entity
	User interface {
		Id() ID
		Username() string
		Email() Email
		Active() bool
		Timestamp
	}

	userBuilder struct {
		user
	}

	user struct {
		id        ID
		username  string
		email     Email
		active    bool
		createdAt time.Time
		updatedAt time.Time
	}
)

func NewUserBuilder(id ID, username string, email Email) UserBuilder {
	return &userBuilder{user{id: id, username: username, email: email}}
}

func (ub *userBuilder) Active(active bool) UserBuilder {
	ub.active = active
	return ub
}

func (ub *userBuilder) CreatedAt(time time.Time) UserBuilder {
	ub.createdAt = time
	return ub
}

func (ub *userBuilder) UpdatedAt(time time.Time) UserBuilder {
	ub.updatedAt = time
	return ub
}

func (ub *userBuilder) Build() User {
	return &user{
		id:        ub.id,
		username:  ub.username,
		email:     ub.email,
		active:    ub.active,
		createdAt: ub.createdAt,
		updatedAt: ub.updatedAt,
	}
}

func (u *user) Id() ID {
	return u.id
}

func (u *user) Username() string {
	return u.username
}

func (u *user) Email() Email {
	return u.email
}

func (u *user) Active() bool {
	return u.active
}

func (u *user) CreatedAt() time.Time {
	return u.createdAt
}

func (u *user) UpdatedAt() time.Time {
	return u.updatedAt
}

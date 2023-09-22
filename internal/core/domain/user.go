package domain

import "time"

type (
	// UserBuilder constructs a User entity
	UserBuilder interface {
		Id(int) UserBuilder
		Username(string) UserBuilder
		Email(Email) UserBuilder
		CreatedAt(time.Time) UserBuilder
		UpdatedAt(time.Time) UserBuilder
		Builder[User]
	}

	// User represents a user entity
	User interface {
		Id() int
		Username() string
		Email() Email
		Timestamp
	}

	userBuilder struct {
		user
	}

	user struct {
		id        int
		username  string
		email     Email
		createdAt time.Time
		updatedAt time.Time
	}
)

func NewUserBuilder(username string, email Email) UserBuilder {
	return &userBuilder{user{id: -1, username: username, email: email}}
}

func (ub *userBuilder) Id(id int) UserBuilder {
	ub.id = id
	return ub
}

func (ub *userBuilder) Username(username string) UserBuilder {
	ub.username = username
	return ub
}

func (ub *userBuilder) Email(email Email) UserBuilder {
	ub.email = email
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
		createdAt: ub.createdAt,
		updatedAt: ub.updatedAt,
	}
}

func (u *user) Id() int {
	return u.id
}

func (u *user) Username() string {
	return u.username
}

func (u *user) Email() Email {
	return u.email
}

func (u *user) CreatedAt() time.Time {
	return u.createdAt
}

func (u *user) UpdatedAt() time.Time {
	return u.updatedAt
}

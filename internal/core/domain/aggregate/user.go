package aggregate

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/model"
	valueobject2 "github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
	"time"
)

type (
	// UserBuilder constructs a User entity
	UserBuilder interface {
		Active(bool) UserBuilder
		CreatedAt(time.Time) UserBuilder
		UpdatedAt(time.Time) UserBuilder
		model.Builder[User]
	}

	// User represents a user entity
	User interface {
		model.BaseAggregate
		Username() string
		Email() valueobject2.Email
		Active() bool
	}

	userBuilder struct {
		user
	}

	user struct {
		*baseAggregate
		username string
		email    valueobject2.Email
		active   bool
	}
)

func NewUserBuilder(id model.ID, username string, email valueobject2.Email) UserBuilder {
	base := newBaseAggregate(id)
	return &userBuilder{user{baseAggregate: base, username: username, email: email}}
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
		baseAggregate: ub.baseAggregate,
		username:      ub.username,
		email:         ub.email,
		active:        ub.active,
	}
}

func (u *user) Username() string {
	return u.username
}

func (u *user) Email() valueobject2.Email {
	return u.email
}

func (u *user) Active() bool {
	return u.active
}

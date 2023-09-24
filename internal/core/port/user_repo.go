package port

import (
	"database/sql"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
)

type (
	UserRepo interface {
		Save(user aggregate.User) (sql.Result, error)
		Find(id valueobject.UserID) (aggregate.User, error)
	}
)

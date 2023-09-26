package port

import (
	"database/sql"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
)

type (
	UserRepo interface {
		Save(aggregate.User) (sql.Result, error)
		Find(valueobject.UserID) (aggregate.User, error)
	}
)

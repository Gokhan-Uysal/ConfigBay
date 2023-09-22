package port

import (
	"database/sql"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain"
)

type (
	UserRepo interface {
		Create(user domain.User) (sql.Result, error)
		GetById(id domain.ID) (domain.User, error)
	}
)

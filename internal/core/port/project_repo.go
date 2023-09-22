package port

import (
	"database/sql"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain"
)

type (
	ProjectRepo interface {
		Save(domain.Project) (sql.Result, error)
	}
)

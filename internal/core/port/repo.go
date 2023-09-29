package port

import (
	"database/sql"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
)

type (
	BaseRepo interface {
		Exec(*sql.Tx, string, ...any) (sql.Result, error)
		Query(*sql.Tx, string, ...any) (*sql.Rows, error)
		QueryRow(
			*sql.Tx,
			string,
			...any,
		) *sql.Row
		CloseRows(rows *sql.Rows)
		Begin() (*sql.Tx, error)
		CommitOrRollback(*sql.Tx, error) error
	}

	ProjectRepo interface {
		Save(aggregate.Project) error
		Find(valueobject.ProjectID) (aggregate.Project, error)
	}

	GroupRepo interface {
		Save(group aggregate.Group) error
		Find(valueobject.GroupID) (aggregate.Group, error)
	}

	UserRepo interface {
		Save(aggregate.User) (sql.Result, error)
		Find(valueobject.UserID) (aggregate.User, error)
	}
)

package port

import (
	"database/sql"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain"
)

type (
	ProjectRepo interface {
		Save(project domain.Project) error
		AssignGroupToProject(
			tx *sql.Tx,
			projectId domain.ID,
			group domain.Group,
		) (sql.Result, error)
		AssignUserToGroup(
			tx *sql.Tx,
			groupId domain.ID,
			userId domain.ID,
		) (sql.Result, error)
		AssignRoleToGroup(
			tx *sql.Tx,
			groupId domain.ID,
			role domain.Role,
		) (sql.Result, error)
	}
)

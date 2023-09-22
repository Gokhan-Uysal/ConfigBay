package repo

import (
	"database/sql"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
)

type (
	projectRepo struct {
		db *sql.DB
	}
)

func NewProjectRepo(db *sql.DB) port.ProjectRepo {
	if db == nil {

	}
	return &projectRepo{db: db}
}

func (p projectRepo) Save(project domain.Project) (sql.Result, error) {
	//TODO implement me
	panic("implement me")
}

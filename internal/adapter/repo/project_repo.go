package repo

import (
	"database/sql"

	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
)

type (
	projectRepo struct {
		baseRepo port.BaseRepo
	}
)

func NewProjectRepo(db *sql.DB) (port.ProjectRepo, error) {
	base, err := newBaseRepo(db)
	if err != nil {
		return nil, err
	}
	return &projectRepo{baseRepo: base}, nil
}

func (pr *projectRepo) Save(project domain.Project) error {
	var (
		tx  *sql.Tx
		err error
	)

	tx, err = pr.baseRepo.Begin()
	if err != nil {
		logger.ERR.Println(err)
		return pr.baseRepo.CommitOrRollback(tx, err)
	}

	_, err = pr.SaveProject(tx, project)
	if err != nil {
		logger.ERR.Println(err)
		return pr.baseRepo.CommitOrRollback(tx, err)
	}

	for _, group := range project.Groups() {
		_, err = pr.AssignGroupToProject(tx, project.Id(), group)
		if err != nil {
			logger.ERR.Println(err)
			return pr.baseRepo.CommitOrRollback(tx, err)
		}

		for _, role := range group.Roles() {
			_, err = pr.AssignRoleToGroup(tx, group.Id(), role)
			if err != nil {
				logger.ERR.Println(err)
				return pr.baseRepo.CommitOrRollback(tx, err)
			}
		}

		for _, user := range group.Users() {
			_, err = pr.AssignUserToGroup(tx, group.Id(), user.Id())
			if err != nil {
				logger.ERR.Println(err)
				return pr.baseRepo.CommitOrRollback(tx, err)
			}
		}
	}

	return nil
}

func (pr *projectRepo) SaveProject(tx *sql.Tx, project domain.Project) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = pr.baseRepo.Exec(
		tx,
		"INSERT INTO projects (id, title) VALUES ($1, $2)",
		project.Id(), project.Title(),
	)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	return result, nil
}

func (pr *projectRepo) AssignGroupToProject(
	tx *sql.Tx,
	projectId domain.ID,
	group domain.Group,
) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = pr.baseRepo.Exec(
		tx,
		"INSERT INTO groups (id, title, project_id) VALUES ($1, $2, $3)",
		group.Id(), group.Title(), projectId,
	)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	return result, nil
}

func (pr *projectRepo) AssignUserToGroup(
	tx *sql.Tx,
	groupId domain.ID,
	userId domain.ID,
) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = pr.baseRepo.Exec(
		tx,
		"INSERT INTO group_users (group_id, user_id) VALUES ($1, $2)",
		groupId, userId,
	)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	return result, nil
}

func (pr *projectRepo) AssignRoleToGroup(
	tx *sql.Tx,
	groupId domain.ID,
	role domain.Role,
) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = pr.baseRepo.Exec(
		tx,
		"INSERT INTO group_roles (group_id, role) VALUES ($1, $2)",
		groupId, role,
	)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	return result, nil
}

package repo

import (
	"database/sql"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/entity"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
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

func (pr projectRepo) Save(project aggregate.Project) error {
	var (
		tx  *sql.Tx
		err error
	)
	logger.DEBUG.Println("Starting to save project.")

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

	for _, groupId := range project.Groups() {
		_, err = pr.AssignGroup(tx, project.Id(), groupId)
		if err != nil {
			logger.ERR.Println(err)
			return pr.baseRepo.CommitOrRollback(tx, err)
		}

	}

	for _, secret := range project.Secrets() {
		_, err = pr.AddSecret(tx, project.Id(), secret)
		if err != nil {
			logger.ERR.Println(err)
			return pr.baseRepo.CommitOrRollback(tx, err)
		}

	}

	logger.DEBUG.Println("Project saved.")
	return pr.baseRepo.CommitOrRollback(tx, err)
}

func (pr projectRepo) SaveProject(tx *sql.Tx, project aggregate.Project) (sql.Result, error) {
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

func (pr projectRepo) AssignGroup(
	tx *sql.Tx,
	projectId valueobject.ID,
	groupId valueobject.ID,
) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = pr.baseRepo.Exec(
		tx,
		"INSERT INTO project_groups (project_id, group_id) VALUES ($1, $2)",
		projectId, groupId,
	)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	return result, nil
}

func (pr projectRepo) AddSecret(
	tx *sql.Tx,
	projectId valueobject.ID,
	secret entity.Secret,
) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = pr.baseRepo.Exec(
		tx,
		"INSERT INTO secrets (id, key, value, project_id) VALUES ($1, $2, $3, $4)",
		secret.Id(), secret.Key(), secret.Value(), projectId,
	)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	return result, nil
}

func (pr projectRepo) UpdateSecretValue(
	tx *sql.Tx, projectId valueobject.ID, secret entity.Secret,
) (sql.
	Result,
	error) {
	var (
		result sql.Result
		err    error
	)

	result, err = pr.baseRepo.Exec(
		tx,
		"UPDATE secrets SET value=$1 WHERE id=$2 AND project_id=$3",
		secret.Value(), secret.Id(), projectId,
	)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	return result, nil
}

func (pr projectRepo) DeleteSecret(
	tx *sql.Tx,
	projectId valueobject.ID,
	secretId valueobject.ID,
) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = pr.baseRepo.Exec(
		tx,
		"DELETE FROM secrets WHERE id=$1 AND project_id=$2",
		secretId, projectId,
	)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	return result, nil
}

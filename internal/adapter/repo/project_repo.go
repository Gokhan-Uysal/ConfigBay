package repo

import (
	"database/sql"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/entity"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
	"github.com/google/uuid"
	"time"
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

func (pr projectRepo) Find(projectId valueobject.ProjectID) (aggregate.Project, error) {
	var (
		project  aggregate.Project
		secrets  []entity.Secret
		groupIds []valueobject.GroupID
		err      error
	)

	project, err = pr.FindProject(projectId)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	secrets, err = pr.FindSecrets(projectId)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	groupIds, err = pr.FindGroups(projectId)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	project.AddGroups(groupIds...)
	project.AddSecrets(secrets...)

	return project, nil
}

func (pr projectRepo) Update(project aggregate.Project) error {
	var (
		tx *sql.Tx
		_  error
	)

	return pr.baseRepo.CommitOrRollback(tx, nil)
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
		return nil, err
	}

	return result, nil
}

func (pr projectRepo) FindProject(projectId valueobject.ProjectID) (aggregate.Project, error) {
	var (
		project aggregate.Project
		row     *sql.Row
		err     error
	)

	row = pr.baseRepo.QueryRow(nil, "SELECT * FROM projects WHERE id=$1", projectId)

	project, err = pr.MapProject(row)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (pr projectRepo) FindGroups(projectId valueobject.ProjectID) ([]valueobject.GroupID, error) {
	var (
		groupIds []valueobject.GroupID
		rows     *sql.Rows
		err      error
	)

	rows, err = pr.baseRepo.Query(
		nil, "SELECT id FROM groups WHERE project_id=$1",
		projectId,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id uuid.UUID
		)

		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		groupIds = append(groupIds, id)
	}

	defer pr.baseRepo.CloseRows(rows)
	return groupIds, nil
}

func (pr projectRepo) AddSecret(
	tx *sql.Tx,
	projectId valueobject.ProjectID,
	secret entity.Secret,
) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = pr.baseRepo.Exec(
		tx,
		"INSERT INTO secrets (project_id, key, value) VALUES ($1, $2, $3)",
		projectId, secret.Key(), secret.Value(),
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (pr projectRepo) FindSecrets(projectId valueobject.ProjectID) ([]entity.Secret, error) {
	var (
		rows    *sql.Rows
		secrets []entity.Secret
		err     error
	)

	rows, err = pr.baseRepo.Query(
		nil,
		"SELECT key, value, version, created_at, updated_at FROM secrets WHERE project_id=$1",
		projectId,
	)

	for rows.Next() {
		var (
			secret entity.Secret
		)

		secret, err = pr.MapSecret(rows)
		if err != nil {
			return nil, err
		}

		secrets = append(secrets, secret)
	}

	defer pr.baseRepo.CloseRows(rows)
	return secrets, nil
}

func (pr projectRepo) UpdateSecretValue(
	tx *sql.Tx,
	projectId valueobject.ProjectID,
	secret entity.Secret,
) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = pr.baseRepo.Exec(
		tx,
		"UPDATE secrets SET value=$1 WHERE project_id=$2 AND key=$3",
		secret.Value(), projectId, secret.Key(),
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (pr projectRepo) DeleteSecret(
	tx *sql.Tx,
	projectId valueobject.ProjectID,
	key string,
) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = pr.baseRepo.Exec(
		tx,
		"DELETE FROM secrets WHERE project_id=$1 AND key=$2",
		projectId, key,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (pr projectRepo) MapProject(s Scanner) (aggregate.Project, error) {
	var (
		project   aggregate.Project
		id        uuid.UUID
		title     string
		createdAt time.Time
		updatedAt time.Time
		err       error
	)

	err = s.Scan(&id, &title, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	project = aggregate.NewProjectBuilder(id, title).
		CreatedAt(createdAt).
		UpdatedAt(updatedAt).
		Build()

	return project, err
}

func (pr projectRepo) MapSecret(s Scanner) (entity.Secret, error) {
	var (
		secret    entity.Secret
		key       string
		value     string
		version   int
		createdAt time.Time
		updatedAt time.Time
		err       error
	)

	err = s.Scan(&key, &value, &version, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	secret = entity.NewSecretBuilder(key, value).
		Version(version).
		CreatedAt(createdAt).
		UpdatedAt(updatedAt).
		Build()

	return secret, nil
}

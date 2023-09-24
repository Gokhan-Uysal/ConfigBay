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

func (pr projectRepo) Find(projectId valueobject.ID) (aggregate.Project, error) {
	var (
		project  aggregate.Project
		secrets  []entity.Secret
		groupIds []valueobject.ID
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

func (pr projectRepo) FindProject(projectId valueobject.ID) (aggregate.Project, error) {
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
		return nil, err
	}

	return result, nil
}

func (pr projectRepo) FindGroups(projectId valueobject.ID) ([]valueobject.ID, error) {
	var (
		projectIds []valueobject.ID
		rows       *sql.Rows
		err        error
	)

	rows, err = pr.baseRepo.Query(
		nil, "SELECT group_id FROM project_groups WHERE project_id=$1",
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

		projectIds = append(projectIds, id)
	}

	defer pr.baseRepo.CloseRows(rows)
	return projectIds, nil
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
		return nil, err
	}

	return result, nil
}

func (pr projectRepo) FindSecrets(projectId valueobject.ID) ([]entity.Secret, error) {
	var (
		rows    *sql.Rows
		secrets []entity.Secret
		err     error
	)

	rows, err = pr.baseRepo.Query(
		nil,
		"SELECT id, key, value, version, created_at, updated_at FROM secrets WHERE project_id=$1",
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
		id        uuid.UUID
		key       string
		value     string
		version   int
		createdAt time.Time
		updatedAt time.Time
		err       error
	)

	err = s.Scan(&id, &key, &value, &version, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	secret = entity.NewSecretBuilder(id).
		Key(key).
		Value(value).
		Version(version).
		CreatedAt(createdAt).
		UpdatedAt(updatedAt).
		Build()

	return secret, nil
}

package repo

import (
	"database/sql"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
	"github.com/google/uuid"
	"time"
)

type (
	groupRepo struct {
		*baseRepo
	}
)

func NewGroupRepo(db *sql.DB) (port.GroupRepo, error) {
	base, err := newBaseRepo(db)
	if err != nil {
		return nil, err
	}

	return &groupRepo{baseRepo: base}, nil
}

func (gr groupRepo) Save(group aggregate.Group) error {
	var (
		tx  *sql.Tx
		err error
	)
	logger.DEBUG.Println("Starting to save group.")

	tx, err = gr.Begin()
	if err != nil {
		logger.ERR.Println(err)
		return gr.CommitOrRollback(tx, err)
	}

	_, err = gr.SaveGroup(tx, group)
	if err != nil {
		logger.ERR.Println(err)
		return gr.CommitOrRollback(tx, err)
	}

	for _, userId := range group.Users() {
		_, err = gr.AssignUser(tx, group.Id(), userId)
		if err != nil {
			logger.ERR.Println(err)
			return gr.CommitOrRollback(tx, err)
		}
	}

	logger.DEBUG.Println("Group saved.")
	return gr.CommitOrRollback(tx, err)
}

func (gr groupRepo) Find(groupId valueobject.GroupID) (aggregate.Group, error) {
	var (
		group   aggregate.Group
		role    valueobject.Role
		userIds []valueobject.UserID
		err     error
	)

	group, err = gr.FindGroup(groupId)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	role, err = gr.FindRole(groupId)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	userIds, err = gr.FindUsers(groupId)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	group.SetRole(role)
	group.AddUsers(userIds...)

	return group, nil
}

func (gr groupRepo) SaveGroup(
	tx *sql.Tx,
	group aggregate.Group,
) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = gr.Exec(
		tx,
		"INSERT INTO groups (id, title, project_id, role) VALUES ($1, $2, $3, $4)",
		group.Id(), group.Title(), group.ProjectId(), group.RoleName(),
	)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	return result, nil
}

func (gr groupRepo) FindGroup(groupId valueobject.GroupID) (aggregate.Group, error) {
	var (
		group aggregate.Group
		row   *sql.Row
		err   error
	)

	row = gr.QueryRow(
		nil,
		"SELECT id, title, project_id ,created_at, updated_at FROM groups WHERE id=$1",
		groupId,
	)

	group, err = gr.MapGroup(row)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	return group, nil
}

func (gr groupRepo) AssignUser(tx *sql.Tx, groupId, userId valueobject.UserID) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = gr.Exec(
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

func (gr groupRepo) FindUsers(groupId valueobject.GroupID) ([]valueobject.UserID, error) {
	var (
		userIds []valueobject.UserID
		rows    *sql.Rows
		err     error
	)

	rows, err = gr.Query(
		nil,
		"SELECT user_id FROM group_users WHERE group_id=$1",
		groupId,
	)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	for rows.Next() {
		var (
			userId uuid.UUID
		)

		err = rows.Scan(&userId)
		if err != nil {
			logger.ERR.Println(err)
			return nil, err
		}

		userIds = append(userIds, userId)
	}

	return userIds, nil
}

func (gr groupRepo) DropUser(tx *sql.Tx, groupId, userId valueobject.UserID) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = gr.Exec(
		tx,
		"DELETE FROM group_users WHERE group_id=$1 AND user_id=$2",
		groupId, userId,
	)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	return result, nil
}

func (gr groupRepo) FindRole(groupId valueobject.GroupID) (valueobject.Role, error) {
	var (
		role        valueobject.Role
		permissions []valueobject.Permission
		rows        *sql.Rows
		row         *sql.Row
		err         error
	)

	row = gr.QueryRow(
		nil,
		"SELECT r.name FROM roles r INNER JOIN groups g ON g.role=r.name WHERE g.id=$1",
		groupId,
	)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	role, err = gr.MapRole(row)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	rows, err = gr.Query(
		nil,
		"SELECT p.name FROM role_permissions rp INNER JOIN roles r ON r.name=rp.name INNER JOIN permissions p ON p.name=rp.name WHERE r.name=$1",
		role.Name(),
	)

	for rows.Next() {
		var (
			permission valueobject.Permission
		)

		permission, err = gr.MapPermission(rows)
		if err != nil {
			logger.ERR.Println(err)
			return nil, err
		}

		permissions = append(permissions, permission)
	}

	role.AddPermissions(permissions...)

	defer gr.CloseRows(rows)
	return role, nil
}

func (gr groupRepo) SaveRole(role valueobject.Role) error {
	var (
		tx  *sql.Tx
		err error
	)

	tx, err = gr.Begin()
	if err != nil {
		logger.ERR.Println(err)
		return err
	}

	_, err = gr.Exec(
		tx,
		"INSERT INTO roles (name) VALUES ($1)",
		role.Name(),
	)
	if err != nil {
		logger.ERR.Println(err)
		return gr.CommitOrRollback(tx, err)
	}

	for _, permission := range role.Permissions() {
		_, err = gr.SavePermission(tx, permission)
		if err != nil {
			logger.ERR.Println(err)
			return gr.CommitOrRollback(tx, err)
		}

		_, err := gr.AssignPermission(tx, role, permission)
		if err != nil {
			logger.ERR.Println(err)
			return gr.CommitOrRollback(tx, err)
		}
	}

	return nil
}

func (gr groupRepo) AssignPermission(
	tx *sql.Tx,
	role valueobject.Role,
	permission valueobject.Permission,
) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = gr.Exec(
		tx,
		"INSERT INTO role_permissions (role, permission) VALUES ($1, $2)",
		role.Name(), valueobject.ToString(permission),
	)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	return result, nil
}

func (gr groupRepo) SavePermission(
	tx *sql.Tx,
	permission valueobject.Permission,
) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = gr.Exec(
		tx,
		"INSERT INTO permissions (name) VALUES ($1)",
		valueobject.ToString(permission),
	)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	return result, nil
}

func (gr groupRepo) DropPermission(
	tx *sql.Tx,
	role valueobject.Role,
	permission valueobject.Permission,
) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = gr.Exec(
		tx,
		"DELETE FROM role_permissions WHERE role=$1 AND permission=$2",
		role.Name(), valueobject.ToString(permission),
	)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	return result, nil
}

func (gr groupRepo) MapGroup(s Scanner) (aggregate.Group, error) {
	var (
		id        uuid.UUID
		title     string
		projectId uuid.UUID
		createdAt time.Time
		updatedAt time.Time
		group     aggregate.Group
		err       error
	)

	err = s.Scan(&id, &title, &projectId, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	group = aggregate.NewGroupBuilder(id, title, projectId).
		CreatedAt(createdAt).
		UpdatedAt(updatedAt).
		Build()

	return group, nil
}

func (gr groupRepo) MapRole(s Scanner) (valueobject.Role, error) {
	var (
		name string
		err  error
	)

	err = s.Scan(&err, &name)
	if err != nil {
		return nil, err
	}

	return valueobject.NewRole(name), nil
}

func (gr groupRepo) MapPermission(s Scanner) (valueobject.Permission, error) {
	var (
		name       string
		permission valueobject.Permission
		err        error
	)

	err = s.Scan(&name)
	if err != nil {
		return "", err
	}

	permission, err = valueobject.ToPermission(name)
	if err != nil {
		return "", err
	}

	return permission, nil
}

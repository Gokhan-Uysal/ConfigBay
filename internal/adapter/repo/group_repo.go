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
		baseRepo port.BaseRepo
	}
)

func NewGroupRepo(db *sql.DB) (port.GroupRepo, error) {
	base, err := newBaseRepo(db)
	if err != nil {
		return nil, err
	}

	return &groupRepo{baseRepo: base}, nil
}

func (gr groupRepo) Save(group aggregate.Group, projectId valueobject.ProjectID) error {
	var (
		tx  *sql.Tx
		err error
	)
	logger.DEBUG.Println("Starting to save group.")

	tx, err = gr.baseRepo.Begin()
	if err != nil {
		logger.ERR.Println(err)
		return gr.baseRepo.CommitOrRollback(tx, err)
	}

	_, err = gr.SaveGroup(tx, group, projectId)
	if err != nil {
		logger.ERR.Println(err)
		return gr.baseRepo.CommitOrRollback(tx, err)
	}

	for _, role := range group.Roles() {
		_, err = gr.AssignRole(tx, group.Id(), role)
		if err != nil {
			logger.ERR.Println(err)
			return gr.baseRepo.CommitOrRollback(tx, err)
		}
	}

	for _, userId := range group.Users() {
		_, err = gr.AssignUser(tx, group.Id(), userId)
		if err != nil {
			logger.ERR.Println(err)
			return gr.baseRepo.CommitOrRollback(tx, err)
		}

	}

	logger.DEBUG.Println("Group saved.")
	return gr.baseRepo.CommitOrRollback(tx, err)
}

func (gr groupRepo) Find(groupId valueobject.GroupID) (aggregate.Group, error) {
	var (
		group   aggregate.Group
		roles   []valueobject.Role
		userIds []valueobject.UserID
		err     error
	)

	group, err = gr.FindGroup(groupId)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	roles, err = gr.FindRolesByGroup(groupId)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	userIds, err = gr.FindUsers(groupId)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	group.AddRoles(roles...)
	group.AddUsers(userIds...)

	return group, nil
}

func (gr groupRepo) SaveGroup(
	tx *sql.Tx,
	group aggregate.Group,
	projectId valueobject.ProjectID,
) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = gr.baseRepo.Exec(
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

func (gr groupRepo) FindGroup(groupId valueobject.GroupID) (aggregate.Group, error) {
	var (
		group aggregate.Group
		row   *sql.Row
		err   error
	)

	row = gr.baseRepo.QueryRow(
		nil,
		"SELECT id, title, created_at, updated_at FROM groups WHERE id=$1",
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

	result, err = gr.baseRepo.Exec(
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

	rows, err = gr.baseRepo.Query(
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

	result, err = gr.baseRepo.Exec(
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

func (gr groupRepo) AssignRole(
	tx *sql.Tx, groupId valueobject.GroupID, role valueobject.Role,
) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = gr.baseRepo.Exec(
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

func (gr groupRepo) FinRoleByName(name valueobject.Role) (valueobject.Role, error) {
	var (
		row  *sql.Row
		role valueobject.Role
		err  error
	)

	row = gr.baseRepo.QueryRow(nil, "SELECT name FROM roles WHERE name=$1", name)
	role, err = gr.MapRole(row)
	if err != nil {
		logger.ERR.Println(err)
		return "", err
	}

	return role, nil
}

func (gr groupRepo) FindRolesByGroup(groupId valueobject.GroupID) ([]valueobject.Role, error) {
	var (
		roles []valueobject.Role
		rows  *sql.Rows
		err   error
	)

	rows, err = gr.baseRepo.Query(
		nil,
		"SELECT r.name FROM group_roles gr INNER JOIN roles r ON r.name=gr.role WHERE gr.group_id=$1",
		groupId,
	)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	for rows.Next() {
		var (
			role valueobject.Role
		)

		role, err = gr.MapRole(rows)
		if err != nil {
			logger.ERR.Println(err)
			return nil, err
		}

		roles = append(roles, role)
	}

	defer gr.baseRepo.CloseRows(rows)
	return roles, nil
}

func (gr groupRepo) DropRole(
	tx *sql.Tx,
	groupId valueobject.GroupID,
	role valueobject.Role,
) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = gr.baseRepo.Exec(
		tx,
		"DELETE FROM group_roles WHERE group_id=$1 AND role=$2",
		groupId, role,
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
		createdAt time.Time
		updatedAt time.Time
		group     aggregate.Group
		err       error
	)

	err = s.Scan(&id, &title, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	group = aggregate.NewGroupBuilder(id, title).
		CreatedAt(createdAt).
		UpdatedAt(updatedAt).
		Build()

	return group, nil
}

func (gr groupRepo) MapRole(s Scanner) (valueobject.Role, error) {
	var (
		name string
		role valueobject.Role
		err  error
	)

	err = s.Scan(&name)
	if err != nil {
		return "", err
	}

	role, err = valueobject.ToRoleName(name)
	if err != nil {
		return "", err
	}

	return role, nil
}

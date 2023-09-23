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

func (gr groupRepo) Save(group aggregate.Group) error {
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

	_, err = gr.SaveGroup(tx, group)
	if err != nil {
		logger.ERR.Println(err)
		return gr.baseRepo.CommitOrRollback(tx, err)
	}

	for _, roleId := range group.Roles() {
		_, err = gr.AssignRole(tx, group.Id(), roleId)
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

func (gr groupRepo) SaveGroup(tx *sql.Tx, group aggregate.Group) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = gr.baseRepo.Exec(
		tx,
		"INSERT INTO groups (id, title) VALUES ($1, $2)",
		group.Id(), group.Title(),
	)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	return result, nil
}

func (gr groupRepo) AssignUser(tx *sql.Tx, groupId, userId valueobject.ID) (sql.Result, error) {
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

func (gr groupRepo) DropUser(tx *sql.Tx, groupId, userId valueobject.ID) (sql.Result, error) {
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
	tx *sql.Tx, groupId valueobject.ID, role entity.Role,
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

func (pr projectRepo) DropRole(
	tx *sql.Tx,
	groupId valueobject.ID,
	role entity.Role,
) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = pr.baseRepo.Exec(
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

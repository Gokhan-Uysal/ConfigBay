package repo

import (
	"database/sql"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
	"github.com/google/uuid"
	"time"

	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
)

type (
	userRepo struct {
		baseRepo port.BaseRepo
	}
)

func NewUserRepo(db *sql.DB) (port.UserRepo, error) {
	base, err := newBaseRepo(db)
	if err != nil {
		return nil, err
	}

	return &userRepo{baseRepo: base}, nil
}

func (ur *userRepo) Create(user domain.User) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	result, err = ur.baseRepo.Exec(
		nil,
		"INSERT INTO users (id, username, email) VALUES ($1, $2, $3)", user.Id(),
		user.Username(), user.Email().String(),
	)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}
	return result, nil
}

func (ur *userRepo) GetById(id domain.ID) (domain.User, error) {
	var (
		rows  *sql.Rows
		users []domain.User
		err   error
	)

	rows, err = ur.baseRepo.Query(
		nil,
		"SELECT id, username, email, active, created_at, updated_at FROM users WHERE id=$1",
		id.String(),
	)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	users, err = ur.mapUser(rows)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}
	if len(users) == 0 {
		err := domain.ItemNotFoundErr{Item: "user"}
		logger.ERR.Println(err)
		return nil, err
	}

	return users[0], nil
}

func (ur *userRepo) mapUser(rows *sql.Rows) ([]domain.User, error) {
	var (
		users []domain.User
	)

	defer ur.baseRepo.CloseRows(rows)

	for rows.Next() {
		var (
			id        uuid.UUID
			username  string
			email     string
			active    bool
			createdAt time.Time
			updatedAt time.Time
			user      domain.User
			err       error
		)
		err = rows.Scan(&id, &username, &email, &active, &createdAt, &updatedAt)
		if err != nil {
			logger.ERR.Println(err)
			return nil, err
		}

		user = domain.NewUserBuilder(id, username, domain.NewEmail(email)).
			Active(active).
			CreatedAt(createdAt).
			UpdatedAt(updatedAt).
			Build()

		users = append(users, user)
	}

	return users, nil
}

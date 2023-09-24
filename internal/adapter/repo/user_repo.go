package repo

import (
	"database/sql"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
	valueobject2 "github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
	"github.com/google/uuid"
	"time"

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

func (ur *userRepo) Save(user aggregate.User) (sql.Result, error) {
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

func (ur *userRepo) Find(userId valueobject.UserID) (aggregate.User, error) {
	var (
		row  *sql.Row
		user aggregate.User
		err  error
	)

	row = ur.baseRepo.QueryRow(
		nil,
		"SELECT id, username, email, active, created_at, updated_at FROM users WHERE id=$1",
		userId.String(),
	)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	user, err = ur.mapUser(row)
	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}

	return user, nil
}

func (ur *userRepo) mapUser(s Scanner) (aggregate.User, error) {
	var (
		id        uuid.UUID
		username  string
		email     string
		active    bool
		createdAt time.Time
		updatedAt time.Time
		user      aggregate.User
		err       error
	)

	err = s.Scan(&id, &username, &email, &active, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	user = aggregate.NewUserBuilder(id, username, valueobject2.NewEmail(email)).
		Active(active).
		CreatedAt(createdAt).
		UpdatedAt(updatedAt).
		Build()
	return user, nil
}
